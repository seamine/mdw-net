package cutter

/**
 * @Author: zuodebiao
 * @Date: 2021/2/25 下午3:59
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

// 透传Cutter
type ThroughoutCutter struct {}

func (f ThroughoutCutter) Cut(data []byte) (int, bool, error) {
	return len(data), true, nil
}
