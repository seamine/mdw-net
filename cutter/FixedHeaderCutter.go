package cutter

import (
	"encoding/binary"
	"errors"
)

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 下午3:58
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

// 固定包头包体Cutter
type FixedHeaderCutter struct {

	ByteOrder binary.ByteOrder

	// 包头长度，单位：字节
	HeaderBytesNum  int

	// 包头中的包体长度字段偏移, 单位：字节
	LengthOffset 	int

	// 包头中的包体长度字段占用字节数
	LengthBytesNum 	int

	// 包体之后的固定字符数
	Padding int

	MinLengthValue int

	MaxLengthValue int

}

func (f FixedHeaderCutter) Cut(data []byte) (int, bool, error) {

	var bodyLen uint32

	if len(data) < f.HeaderBytesNum {
		return 0, false, nil
	}

	//if f.LengthOffset + f.LengthBytesNum > f.HeaderBytesNum {
	//	return 0, false, errors.New("invalid argument")
	//}

	switch f.LengthBytesNum {
	case 1:
		bodyLen = uint32(data[f.LengthOffset])
	case 2:
		bodyLen = uint32(f.ByteOrder.Uint16(data[f.LengthOffset:f.LengthOffset+2]))
	case 4:
		bodyLen = f.ByteOrder.Uint32(data[f.LengthOffset:f.LengthOffset+4])
	default:
		return 0, false, errors.New("invalid length bytes")
	}

	messageLen := f.HeaderBytesNum + int(bodyLen) + f.Padding

	if messageLen < len(data){
		return 0, false, nil
	}

	return messageLen, true, nil
}