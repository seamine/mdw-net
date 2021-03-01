package main

import "flybees.com.cn/mdw/mdw-net/transport"

/**
 * @Author: zuodebiao
 * @Date: 2021/2/25 下午12:36
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type EchoDumpDataHandler struct {}

func (e EchoDumpDataHandler) Handle(conn *transport.Connection, data []byte) error {
	return nil
}