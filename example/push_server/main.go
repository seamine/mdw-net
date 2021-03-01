package main

import (
	"flybees.com.cn/mdw/mdw-net/transport"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

func PushHandler(w http.ResponseWriter, r *http.Request) {

	var data []byte
	var err error
	var connId int64

	clientId := r.Header.Get("X-Client-Id")

	if len(clientId) == 0 {
		fmt.Fprintln(w, "X-Client-Id not found")
		return
	}

	if data, err = ioutil.ReadAll(r.Body); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	if connId, err = strconv.ParseInt(clientId, 10, 64); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	var conn *transport.Connection

	if conn, err = ConnMgr.Find(connId); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	text := string(data)

	_ = conn.SendMessage(PushItem{
		Id:       1,
		Reserved: 2,
		Text:     text,
	})

	fmt.Fprintln(w, text)
}

func StartPushHttpServer() {
	http.HandleFunc("/", PushHandler)
	if err := http.ListenAndServe("127.0.0.1:8001", nil); err != nil {
		log.Printf("%v", err)
	}
}

func main() {

	ConnMgr = transport.NewConnectionManager(200000)
	Config = NewPushConfig(ConnMgr)

	go PrintStat()
	go StartPushHttpServer()

	listener := transport.NewListener(Config)
	_ = listener.StartTCP(true)
}