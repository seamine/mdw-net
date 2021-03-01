package main

import (
	"github.com/seamine/mdw-net/cutter"
	"github.com/seamine/mdw-net/handler"
	"github.com/seamine/mdw-net/transport"
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