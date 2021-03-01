package handler

import (
	"encoding/hex"
	"flybees.com.cn/mdw/mdw-net/transport"
	"log"
)

/**
 * @Author: zuodebiao
 * @Date: 2021/2/25 下午2:58
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type HexHandler struct {}

func (h HexHandler) Handle(conn *transport.Connection, data []byte) error {

	log.Printf("%v", hex.EncodeToString(data))

	return nil
}