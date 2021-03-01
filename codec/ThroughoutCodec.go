package codec

/**
 * @Author: zuodebiao
 * @Date: 2021/2/25 下午1:46
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

// 透传编码器
type ThroughoutCodec struct {}

func (e ThroughoutCodec) Decode(data []byte) (interface{}, error) {
	return data, nil
}

func (e ThroughoutCodec) Encode(data interface{}) ([]byte, error) {
	return data.([]byte), nil
}