package main

import "github.com/seamine/mdw-net/transport"

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 下午1:40
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type LineConfig struct {

	LP LineProtocol

	ConnMgr *transport.ConnectionManager
}

func NewLineConfig(connMgr *transport.ConnectionManager) *LineConfig {
	ec := &LineConfig{
		ConnMgr:connMgr,
		LP:NewLineProtocol(),
	}
	return ec
}

func (l LineConfig) ConnectionManager() *transport.ConnectionManager {
	return ConnMgr
}

func (l LineConfig) Protocol() transport.ListenProtocol {
	return l.LP
}

func (l LineConfig) Address() string {
	return "localhost:3733"
}

func (l LineConfig) Ipv6Only() bool {
	return false
}

func (l LineConfig) RecvTimeoutSeconds() int {
	return 20
}

func (l LineConfig) SendTimeoutSeconds() int {
	return 20
}

func (l LineConfig) IdleTimeoutSeconds() int {
	return 10
}