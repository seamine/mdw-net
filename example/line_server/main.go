package main

import (
	"github.com/seamine/mdw-net/transport"
	"log"
	"time"
)

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 上午11:51
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

var(
	ConnMgr *transport.ConnectionManager
	Config transport.ListenConfig
)

func PrintStat() {

	ticker := time.NewTicker(3 * time.Second)

	for {
		select {
		case <- ticker.C:
			total, online, active, idle := ConnMgr.Count()
			log.Printf("Total:%v Online:%v Active:%v Idle:%v\n", total, online, active, idle)
		}
	}
}

func main() {

	ConnMgr = transport.NewConnectionManager(200000)
	Config = NewLineConfig(ConnMgr)

	go PrintStat()

	listener := transport.NewListener(Config)
	_ = listener.StartTCP(true)
}