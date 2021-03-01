package main

import "github.com/seamine/mdw-net/transport"

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 下午1:40
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type EchoConfig struct {

	Echo EchoProtocol

	ConnMgr *transport.ConnectionManager
}

func NewEchoConfig(connMgr *transport.ConnectionManager) *EchoConfig {
	ec := &EchoConfig{
		ConnMgr:connMgr,
	}
	return ec
}

func (l EchoConfig) ConnectionManager() *transport.ConnectionManager {
	return ConnMgr
}

func (l EchoConfig) Protocol() transport.ListenProtocol {
	return l.Echo
}

func (l EchoConfig) Address() string {
	return "localhost:3733"
}

func (l EchoConfig) Ipv6Only() bool {
	return false
}

func (l EchoConfig) RecvTimeoutSeconds() int {
	return 20
}

func (l EchoConfig) SendTimeoutSeconds() int {
	return 20
}

func (l EchoConfig) IdleTimeoutSeconds() int {
	return 10
}