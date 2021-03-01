package main

import (
	"flybees.com.cn/mdw/mdw-net/cutter"
	"flybees.com.cn/mdw/mdw-net/handler"
	"flybees.com.cn/mdw/mdw-net/transport"
)

/**
 * @Author: zuodebiao
 * @Date: 2021/2/25 下午3:47
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type EchoInputStack struct {

	Cutter cutter.ThroughoutCutter
	Hex handler.HexHandler

	Dump EchoDumpDataHandler

}

func (p EchoInputStack) GetCutter() transport.DataCutter {
	return p.Cutter
}

func (p EchoInputStack) GetDataHandler() []transport.DataHandler {
	return []transport.DataHandler{p.Hex, p.Dump}
}

type EchoOutputStack struct {
	Hex handler.HexHandler
}

func (p EchoOutputStack) GetDataHandler() []transport.DataHandler {
	return []transport.DataHandler{p.Hex}
}