package main

import "encoding/binary"

/**
 * @Author: zuodebiao
 * @Date: 2021/2/25 下午2:09
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type PushCodec struct {}

func (p PushCodec) Decode(data []byte) (interface{}, error) {

	item := PushItem {
		Id:       data[0],
		Reserved: data[1],
		Text:     string(data[4:len(data)-4]),
	}

	return item, nil
}

func (p PushCodec) Encode(data interface{}) ([]byte, error) {

	item := data.(PushItem)

	jsonBytes := []byte(item.Text)

	dataLen := len(jsonBytes) + 4

	ret := make([]byte, dataLen)
	ret[0] = item.Id
	ret[1] = item.Reserved

	binary.BigEndian.PutUint16(ret[2:4], uint16(len(jsonBytes)))

	copy(ret[4:], jsonBytes)

	return ret, nil
}