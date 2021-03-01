package transport

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 上午9:41
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type MessageCodec interface {

	Decode(data []byte) (interface{}, error)

	Encode(data interface{}) ([]byte, error)

}