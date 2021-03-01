package transport

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 上午9:39
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type IoEventId int

const (
	// 连接已经Accept并加入到连接管理器中
	IoEventAccepted IoEventId = iota
	// 准备连接
	IoEventWillConnect
	// 己连接
	IoEventConnected
	// 将要关闭连接
	IoEventWillClose
	// 连接已关闭
	IoEventClosed
	// 数据到达
	IoEventDataArrived
	// Encoder处理后的数据准备发送
	IoEventDataWillSend
	// Encoder处理后的数据已发送
	IoEventDataSent
	// data decode成功后消息包到达
	IoEventMessageArrived
	// data将要encode
	IoEventMessageWillEncode
	// data已经encode
	IoEventMessageEncoded
	// 连接数据处理时发生错误，可能是网络错误，也可能是应用层处理错误
	IoEventError

	IoEventMax
)

type IoEventData struct {

	Conn *Connection

	Data []byte

	Message interface{}

	Error error

}

type IoEventHandler interface {

	Handle(eventId IoEventId, data *IoEventData) error

}
