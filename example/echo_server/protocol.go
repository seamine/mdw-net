package main

import (
	"github.com/seamine/mdw-net/codec"
	"github.com/seamine/mdw-net/transport"
)

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 下午12:37
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type EchoProtocol struct {
	UStack  EchoInputStack
	DStack  EchoOutputStack
	Coder   codec.ThroughoutCodec
	Handler EchoMessageHandler
}

func (l EchoProtocol) InputStack() transport.ListenInputStack {
	return &l.UStack
}

func (l EchoProtocol) OutputStack() transport.ListenOutputStack {
	return &l.DStack
}

func (l EchoProtocol) GetCodec() transport.MessageCodec {
	return l.Coder
}

func (l EchoProtocol) GetMessageHandler() []transport.MessageHandler {
	return []transport.MessageHandler{l.Handler}
}