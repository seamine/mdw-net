package transport

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 下午2:10
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type DataHandler interface {

	Handle(conn *Connection, data []byte) error

}