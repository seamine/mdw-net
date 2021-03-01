package transport

import (
	"errors"
	"log"
	"net"
	"time"
)

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 上午9:52
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type Listener struct {

	EventMgr EventManager

	listener net.Listener

	Config ListenConfig
}

func NewListener(config ListenConfig) *Listener {

	listener := Listener{}
	listener.Config = config

	return &listener
}

func (l *Listener) StartTCP(join bool) error {

	var err error

	if l.Config.Ipv6Only() {
		if l.listener, err = net.Listen("tcp6", l.Config.Address()); err != nil {
			return err
		}
	} else {
		if l.listener, err = net.Listen("tcp", l.Config.Address()); err != nil {
			return err
		}
	}

	log.Printf("listen %v ok", l.Config.Address())

	if join {
		l.acceptRoutine()
	} else {
		go l.acceptRoutine()
	}

	return nil
}

func (l *Listener) Stop() {

	if l.listener != nil {
		_ = l.listener.Close()
		l.listener = nil
	}
}

func (l *Listener) acceptRoutine() {

	connMgr := l.Config.ConnectionManager()

	for {

		netConn, err := l.listener.Accept()

		if err != nil {

			l.EventMgr.EmitError(nil, err)

			time.Sleep(10 * time.Millisecond)

			continue
		}

		conn := Connection{}
		conn.Listener = l
		conn.Init(netConn)

		if _, ok := connMgr.Add(&conn); !ok {
			l.EventMgr.EmitAcceptFailed(&conn, errors.New("max connections exceeds"))
			continue
		}

		if err = l.EventMgr.EmitAccepted(&conn); err != nil {
			_ = connMgr.Delete(&conn);
			continue
		}

		conn.Start()
	}
}