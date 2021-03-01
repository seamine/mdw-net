package main

import "flybees.com.cn/mdw/mdw-net/transport"

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 下午1:40
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type PushConfig struct {

	LP PushProtocol

	ConnMgr *transport.ConnectionManager
}

func NewPushConfig(connMgr *transport.ConnectionManager) *PushConfig {
	ec := &PushConfig{
		ConnMgr:connMgr,
		LP:NewPushProtocol(),
	}
	return ec
}

func (l PushConfig) ConnectionManager() *transport.ConnectionManager {
	return ConnMgr
}

func (l PushConfig) Protocol() transport.ListenProtocol {
	return l.LP
}

func (l PushConfig) Address() string {
	return "localhost:3733"
}

func (l PushConfig) Ipv6Only() bool {
	return false
}

func (l PushConfig) RecvTimeoutSeconds() int {
	return 20
}

func (l PushConfig) SendTimeoutSeconds() int {
	return 20
}

func (l PushConfig) IdleTimeoutSeconds() int {
	return 10
}