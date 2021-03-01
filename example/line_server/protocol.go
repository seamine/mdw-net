package main

import (
	"github.com/seamine/mdw-net/codec"
	"github.com/seamine/mdw-net/cutter"
	"github.com/seamine/mdw-net/transport"
)

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 下午12:37
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type LineProtocol struct {

	UStack  LineInputStack
	DStack  LineOutputStack

	Coder transport.MessageCodec

	Handler LineMessageHandler

}

func NewLineProtocol() LineProtocol {
	
	protocol := LineProtocol{
		UStack: LineInputStack{
			Cutter: cutter.DelimiterCutter{
				Delimiter:"\r\n",
			},
		},
		Coder: codec.ThroughoutCodec{},
	}

	return protocol
}

func (l LineProtocol) InputStack() transport.ListenInputStack {
	return &l.UStack
}

func (l LineProtocol) OutputStack() transport.ListenOutputStack {
	return &l.DStack
}

func (l LineProtocol) GetCodec() transport.MessageCodec {
	return l.Coder
}

func (l LineProtocol) GetMessageHandler() []transport.MessageHandler {
	return []transport.MessageHandler{l.Handler}
}