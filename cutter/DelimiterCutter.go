package cutter

import (
	"errors"
	"strings"
)

/**
 * @Author: zuodebiao
 * @Date: 2021/2/25 下午1:28
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

// 分割符Cutter
type DelimiterCutter struct {
	Delimiter string
}

func (f DelimiterCutter) Cut(data []byte) (int, bool, error) {

	var delimiterIndex = 0

	str := string(data)

	switch len(f.Delimiter) {
	case 0:
		return 0, false, errors.New("invalid delimiter")
	case 1:
		delimiterIndex = strings.IndexByte(str, f.Delimiter[0])
	default:
		delimiterIndex = strings.Index(str, f.Delimiter)
	}

	delimiterIndex += len(f.Delimiter)

	return delimiterIndex, true, nil
}