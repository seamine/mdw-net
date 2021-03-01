package transport

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 上午10:39
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type ListenConfig interface {

	ConnectionManager() *ConnectionManager

	Protocol() ListenProtocol

	Address() string

	Ipv6Only() bool

	RecvTimeoutSeconds() int

	SendTimeoutSeconds() int

	IdleTimeoutSeconds() int
}
