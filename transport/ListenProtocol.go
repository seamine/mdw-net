package transport

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 上午10:28
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */


type ListenProtocol interface {

	InputStack() ListenInputStack

	OutputStack() ListenOutputStack

	GetCodec() MessageCodec

	GetMessageHandler() []MessageHandler

}