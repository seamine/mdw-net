package transport

import (
	"errors"
	"fmt"
	"sync"
)

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 上午9:22
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type ConnectionManager struct {

	Online map[int64]*Connection
	Active map[int64]*Connection
	Idle   map[int64]*Connection

	lock sync.RWMutex
	max int

	ConnectionIdSequence int64
}

func NewConnectionManager(max int) *ConnectionManager {

	manager := &ConnectionManager{}
	manager.Init(max)

	return manager
}

func (c *ConnectionManager) Init(max int) {

	c.Online = make(map[int64]*Connection)
	c.Active = make(map[int64]*Connection)
	c.Idle = make(map[int64]*Connection)
	c.ConnectionIdSequence = 1
	c.max = max

}

func (c *ConnectionManager) Add(conn *Connection) (int64, bool) {

	c.lock.Lock(); defer c.lock.Unlock()

	if len(c.Online) >= c.max {
		return 0, false
	}

	connId := c.ConnectionIdSequence; c.ConnectionIdSequence += 1
	conn.Id = connId

	c.Online[connId] = conn
	c.Idle[connId] = conn

	return connId, true
}

func (c *ConnectionManager) Find(connId int64) (*Connection, error) {

	var conn *Connection
	var found bool

	c.lock.RLock(); defer c.lock.RUnlock()

	if connId <= 0 {
		return nil, errors.New("invalid connection object")
	}

	if conn, found = c.Online[connId]; !found {
		return nil, errors.New(fmt.Sprintf("connection id %v not found", connId))
	}

	return conn, nil
}

func (c *ConnectionManager) Delete(conn *Connection) error {

	c.lock.Lock(); defer c.lock.Unlock()

	if conn == nil || conn.Id == 0 {
		return errors.New("invalid connection object")
	}

	if _, found := c.Online[conn.Id]; !found {
		return errors.New(fmt.Sprintf("connection id %v not found", conn.Id))
	}

	delete(c.Online, conn.Id)
	delete(c.Active, conn.Id)
	delete(c.Idle, conn.Id)

	return nil
}

func (c *ConnectionManager) DeleteById(connId int64) error {

	c.lock.Lock(); defer c.lock.Unlock()

	if connId <= 0 {
		return errors.New("invalid connection object")
	}

	if _, found := c.Online[connId]; !found {
		return errors.New(fmt.Sprintf("connection id %v not found", connId))
	}

	delete(c.Online, connId)
	delete(c.Active, connId)
	delete(c.Idle, connId)

	return nil
}

func (c *ConnectionManager) DeleteAll() error {

	c.lock.Lock(); defer c.lock.Unlock()

	for connId, _ := range c.Online {
		delete(c.Online, connId)
		delete(c.Active, connId)
		delete(c.Idle, connId)
	}

	return nil
}

func (c *ConnectionManager) SetActive(conn *Connection) error {

	c.lock.Lock(); defer c.lock.Unlock()

	if conn == nil || conn.Id == 0 {
		return errors.New("invalid connection object")
	}

	if _, found := c.Online[conn.Id]; !found {
		return errors.New(fmt.Sprintf("connection id %v not found", conn.Id))
	}

	c.Active[conn.Id] = conn

	delete(c.Idle, conn.Id)

	return nil
}

func (c *ConnectionManager) SetIdle(conn *Connection) error {

	c.lock.Lock(); defer c.lock.Unlock()

	if conn == nil || conn.Id == 0 {
		return errors.New("invalid connection object")
	}

	if _, found := c.Online[conn.Id]; !found {
		return errors.New(fmt.Sprintf("connection id %v not found", conn.Id))
	}

	c.Idle[conn.Id] = conn

	delete(c.Active, conn.Id)

	return nil
}

func (c *ConnectionManager) Count() (int, int, int, int) {

	c.lock.RLock(); defer c.lock.RUnlock()

	return c.max, len(c.Online), len(c.Active), len(c.Idle)
}