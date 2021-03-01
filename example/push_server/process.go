package main

import "flybees.com.cn/mdw/mdw-net/transport"

/**
 * @Author: zuodebiao
 * @Date: 2021/2/25 下午12:36
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type PushMessageHandler struct {}

func (e PushMessageHandler) Handle(conn *transport.Connection, message interface{}) error {
	return conn.SendMessage(message)
}
