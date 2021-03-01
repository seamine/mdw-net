package main

import (
	"encoding/binary"
	"github.com/seamine/mdw-net/cutter"
	"github.com/seamine/mdw-net/transport"
)

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 下午12:37
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type PushProtocol struct {

	UStack PushInputStack
	DStack PushOutputStack

	Coder transport.MessageCodec

	PushHandler PushMessageHandler

	WorkerHandler *transport.WorkerPool

}

func NewPushProtocol() PushProtocol {

	protocol := PushProtocol{
		Coder: PushCodec{},
		UStack: PushInputStack{
			Cutter: cutter.FixedHeaderCutter{
				ByteOrder:      binary.BigEndian,
				HeaderBytesNum: 4,
				LengthOffset:   2,
				LengthBytesNum: 2,
				Padding:        0,
				MinLengthValue: 0,
				MaxLengthValue: 0,
			},},
		WorkerHandler: transport.NewWorkerPool(512, 10, PushMessageHandler{}),
	}

	return protocol
}

func (l PushProtocol) InputStack() transport.ListenInputStack {
	return &l.UStack
}

func (l PushProtocol) OutputStack() transport.ListenOutputStack {
	return &l.DStack
}

func (l PushProtocol) GetCodec() transport.MessageCodec {
	return l.Coder
}

func (l PushProtocol) GetMessageHandler() []transport.MessageHandler {
	return []transport.MessageHandler{l.WorkerHandler}
}