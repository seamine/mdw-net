package transport

import (
	"errors"
	"net"
	"sync"
	"time"
)

/**
 * @Author: zuodebiao
 * @Date: 2021/2/3 上午11:40
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type Connection struct {
	Id       int64
	Context  interface{}
	Listener *Listener
	Debug    bool

	// private
	conn                   net.Conn
	endpoint               string
	lock                   sync.Mutex
	createTime             time.Time
	updateTime             time.Time
	received               int64
	sent                   int64
	idle                   bool
	connMgr                *ConnectionManager
	eventMgr               *EventManager
	upstreamDataHandlers   []DataHandler
	messageHandlers        []MessageHandler
	downstreamDataHandlers []DataHandler
	coder                  MessageCodec

	recvBuf []byte
	dataBuf []byte

	sendQueue    chan interface{}
	sendMsgQueue chan []byte
	idleQueue    chan bool
}

func (c *Connection) Init(conn net.Conn) {
	c.conn = conn
	c.endpoint = c.conn.RemoteAddr().String()
	c.recvBuf = make([]byte, 1024)
	c.sendQueue = make(chan interface{}, 10)
	c.sendMsgQueue = make(chan []byte, 10)
	c.idleQueue = make(chan bool)
	c.createTime = time.Now()
	c.updateTime = c.createTime
	c.idle = true
	c.received = 0
	c.sent = 0
	c.connMgr = c.Listener.Config.ConnectionManager()
	c.eventMgr = &c.Listener.EventMgr
	c.upstreamDataHandlers = c.Listener.Config.Protocol().InputStack().GetDataHandler()
	c.downstreamDataHandlers = c.Listener.Config.Protocol().OutputStack().GetDataHandler()
	c.messageHandlers = c.Listener.Config.Protocol().GetMessageHandler()
	c.coder = c.Listener.Config.Protocol().GetCodec()
}

func (c *Connection) Start() {
	go c.idleRoutine()
	go c.recvRoutine()
	go c.sendRoutine()
}

func (c *Connection) Stop() {

	c.lock.Lock(); defer c.lock.Unlock()

	if c.conn == nil {
		return
	}

	c.eventMgr.EmitWillClose(c, nil)

	c.conn.Close()

	if c.connMgr != nil {
		_ = c.connMgr.Delete(c)
	}

	c.eventMgr.EmitClosed(c, nil)

	c.conn = nil
	c.recvBuf = nil
	c.dataBuf = nil
	c.Id = 0
	c.Context = nil
	c.Listener = nil
	c.endpoint = ""
	c.received = 0
	c.sent = 0
	c.connMgr = nil

	close(c.sendQueue)
	close(c.sendMsgQueue)
	close(c.idleQueue)
}

func (c *Connection) Handle(conn *Connection, data []byte) error {

	var err error
	var message interface{}

	// data handler chain called
	{
		for _, dataHandler := range c.upstreamDataHandlers {
			if dataHandler != nil {
				if err = dataHandler.Handle(c, data); err != nil {
					return err
				}
			}
		}
	}

	// decode
	{
		c.eventMgr.EmitDataArrived(c, data)

		if message, err = c.coder.Decode(data); err != nil {
			return err
		}

		c.eventMgr.EmitMessageArrived(c, message)
	}

	// message handles chain called
	{
		for _, messageHandler := range c.messageHandlers {
			if messageHandler != nil {
				if err = messageHandler.Handle(c, message); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (c *Connection) recvRoutine() {

	var err error

	defer func() {
		c.Stop()
	}()

	config := c.Listener.Config

	if err = c.RecvDataWithCutter(config.Protocol().InputStack().GetCutter(), c, config.RecvTimeoutSeconds()); err != nil {
		c.Listener.EventMgr.EmitError(c, err)
	}
}

func (c *Connection) idleRoutine() {

	var run = true

	ticker := time.NewTicker(60 * time.Minute)
	ticker.Stop()

	defer func() {
		ticker.Stop()
	}()

	idleTimeoutSeconds := c.Listener.Config.IdleTimeoutSeconds()

	for run {
		select {
		case _, run = <- c.idleQueue:
			if run {
				ticker.Reset(time.Duration(idleTimeoutSeconds) * time.Second)
			}
		case <- ticker.C:
			if !c.idle {
				_ = c.connMgr.SetIdle(c)
				c.idle = true
			}
		}
	}
}

func (c *Connection) sendRoutine() {

	var data []byte
	var err error
	var quit = false

	defer func() {
		c.Stop()
	}()

	sendTimeoutSeconds := c.Listener.Config.SendTimeoutSeconds()

	for !quit {

		select {
		case msg := <-c.sendQueue:
			if msg == nil {
				quit = true; continue
			}

			c.eventMgr.EmitMessageWillEncode(c, data)

			if data, err = c.coder.Encode(msg); err != nil {
				quit = true; continue
			}

			c.eventMgr.EmitMessageEncoded(c, data)

		case data := <- c.sendMsgQueue:

			if data == nil {
				quit = true; continue
			}
		}

		for _, dataHandler := range c.downstreamDataHandlers {
			if dataHandler != nil {
				if err = dataHandler.Handle(c, data); err != nil {
					quit = true
					break
				}
			}
		}

		if quit == true {
			break
		}

		c.eventMgr.EmitDataWillSend(c, data)

		if err = c.SendData(data, sendTimeoutSeconds); err != nil {
			quit = true; continue
		}

		c.eventMgr.EmitDataSent(c, data)

		c.SetActive()
		c.updateTime = time.Now()
	}
}

func (c *Connection) SetActive() {
	if c.idle {
		c.idle = false
		_ = c.connMgr.SetActive(c)
		c.idleQueue <- true
	}
}

func (c *Connection) SendMessage(data interface{}) error {

	c.sendQueue <- data

	return nil
}

func (c *Connection) SendData(data []byte, sendTimeoutSeconds int) error {

	var err error
	var sent = 0
	var n int

	offset := 0
	startTime := time.Now()

	for sent < len(data) {

		if err = c.conn.SetWriteDeadline(startTime.Add(time.Duration(sendTimeoutSeconds) * time.Second)); err != nil {
			return err
		}

		if n, err = c.conn.Write(data[offset:]); err != nil {
			return err
		}

		sent += n

		c.sent += int64(n)
	}

	return nil
}

func (c *Connection) RecvDataWithLength(length int, timeout int) ([]byte, error) {

	var data []byte
	var tmp []byte
	var err error

	now := time.Now()

	for len(data) < length {

		tmp = make([]byte, length-len(data))

		if err = c.conn.SetReadDeadline(now.Add(time.Duration(timeout) * time.Second)); err != nil {
			return nil, err
		}

		n, err := c.conn.Read(tmp)

		if err != nil {
			return nil, err
		}

		c.received += int64(n)

		data = append(data, tmp[:n]...)
	}

	return data, nil
}

func (c *Connection) RecvData(timeoutSeconds int) ([]byte, error) {

	var err error
	var data []byte

	now := time.Now()

	if err = c.conn.SetReadDeadline(now.Add(time.Duration(timeoutSeconds) * time.Second)); err != nil {
		return nil, err
	}

	n, err := c.conn.Read(c.recvBuf)

	if err != nil {
		return nil, err
	}

	data = c.recvBuf[:n]
	c.received += int64(n)

	return data, nil
}

func (c *Connection) RecvDataWithCutter(cutter DataCutter, handler DataHandler, timeoutSeconds int) error {

	var err error
	var data []byte
	var truncate = false
	var offset = 0

	if handler == nil {
		return errors.New("handler is null")
	}

	for {

		now := time.Now()

		if err = c.conn.SetReadDeadline(now.Add(time.Duration(timeoutSeconds) * time.Second)); err != nil {
			return err
		}

		n, err := c.conn.Read(c.recvBuf)

		if err != nil {
			return err
		}

		chunkData := c.recvBuf[:n]

		c.updateTime = time.Now()
		c.dataBuf = append(c.dataBuf, chunkData...)
		c.received += int64(n)
		c.SetActive()

		if offset, truncate, err = cutter.Cut(c.dataBuf); err != nil {
			return err
		}

		if truncate && offset > 0 && offset <= len(c.dataBuf) {

			data = c.dataBuf[:offset];

			if err = handler.Handle(c, data); err != nil {
				return err
			}

			c.dataBuf= c.dataBuf[offset:]
		}
	}
}
