package handler

import (
	"github.com/seamine/mdw-net/transport"
	"log"
)

/**
 * @Author: zuodebiao
 * @Date: 2021/3/2 下午1:49
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type StringHandler struct {}

func (h StringHandler) Handle(conn *transport.Connection, data []byte) error {

	log.Printf("%v", string(data))

	return nil
}