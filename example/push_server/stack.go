package main

import (
	"github.com/seamine/mdw-net/handler"
	"github.com/seamine/mdw-net/transport"
)

/**
 * @Author: zuodebiao
 * @Date: 2021/2/25 下午3:14
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type PushInputStack struct {

	Cutter transport.DataCutter
	Hex handler.HexHandler

}

func (p PushInputStack) GetCutter() transport.DataCutter {
	return p.Cutter
}

func (p PushInputStack) GetDataHandler() []transport.DataHandler {
	return []transport.DataHandler{p.Hex}
}

type PushOutputStack struct {
	Hex handler.HexHandler
}

func (p PushOutputStack) GetDataHandler() []transport.DataHandler {
	return []transport.DataHandler{p.Hex}
}