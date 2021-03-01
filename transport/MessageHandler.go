package transport

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 上午11:06
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type MessageHandler interface {

	Handle(conn *Connection, message interface{}) error

}