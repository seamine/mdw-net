package transport

import (
	"github.com/revel/revel"
	"github.com/shirou/gopsutil/cpu"
)

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 上午11:03
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type WorkerItem struct {
	Conn    *Connection
	Message interface{}
}

type WorkerPool struct {
	MessageQueue chan *WorkerItem
}

func NewWorkerPool(maxLen, workerNum int, handler MessageHandler) *WorkerPool {

	wp := &WorkerPool{}
	wp.Init(maxLen, workerNum, handler)

	return wp
}

func (w *WorkerPool) Init(maxLen, workerNum int, handler MessageHandler) {

	var err error

	w.MessageQueue = make(chan *WorkerItem, maxLen)

	if workerNum <= 0 {
		if workerNum, err = cpu.Counts(false); err != nil {
			workerNum = 8
		}
	}

	for i := 0; i < workerNum; i += 1 {
		go w.ProcessMessage(handler)
	}
}

func (w *WorkerPool) Handle(conn *Connection, message interface{}) error {

	w.MessageQueue <- &WorkerItem{
		Conn:    conn,
		Message: message,
	}

	return nil
}

func (w *WorkerPool) ProcessMessage(handler MessageHandler) {

	for {
		select {
		case item := <- w.MessageQueue:
			if err := handler.Handle(item.Conn, item.Message); err != nil {
				revel.AppLog.Errorf("%v", err)
			}
		}
	}
}
