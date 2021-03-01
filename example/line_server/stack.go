package main

import (
	"flybees.com.cn/mdw/mdw-net/handler"
	"flybees.com.cn/mdw/mdw-net/transport"
)

/**
 * @Author: zuodebiao
 * @Date: 2021/2/25 下午4:05
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type LineInputStack struct {

	Cutter transport.DataCutter
	Hex handler.HexHandler

}

func (p LineInputStack) GetCutter() transport.DataCutter {
	return p.Cutter
}

func (p LineInputStack) GetDataHandler() []transport.DataHandler {
	return []transport.DataHandler{p.Hex}
}

type LineOutputStack struct {
	Hex handler.HexHandler
}

func (p LineOutputStack) GetDataHandler() []transport.DataHandler {
	return []transport.DataHandler{p.Hex}
}