package main

import "github.com/seamine/mdw-net/transport"

/**
 * @Author: zuodebiao
 * @Date: 2021/2/25 下午12:36
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type EchoMessageHandler struct {}

func (e EchoMessageHandler) Handle(conn *transport.Connection, message interface{}) error {
	return conn.SendMessage(message)
}
