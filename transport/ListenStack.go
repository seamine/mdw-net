package transport

/**
 * @Author: zuodebiao
 * @Date: 2021/2/25 下午3:25
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */


type ListenInputStack interface {

	GetCutter() DataCutter

	GetDataHandler() []DataHandler

}

type ListenOutputStack interface {

	GetDataHandler() []DataHandler

}