package transport

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 上午10:19
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type DataCutter interface {

	Cut(data []byte) (int, bool, error)

}