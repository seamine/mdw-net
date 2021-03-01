package transport

import "log"

/**
 * @Author: zuodebiao
 * @Date: 2021/2/24 上午11:44
 * Copyright(C) 2019 Xingfeng Technology (Shenzheng) Co,Ltd.
 */

type EventManager struct {

}

func NewEventManager() *EventManager {
	return nil
}

func (e *EventManager) Handle(eventId IoEventId, data *IoEventData) error  {
	switch eventId {
	case IoEventAccepted:
		log.Printf("[%v] %v connected\n", data.Conn.Id, data.Conn.endpoint)
	case IoEventWillClose:
		log.Printf("[%v] %v will close\n", data.Conn.Id, data.Conn.endpoint)
	case IoEventClosed:
		log.Printf("[%v] %v closed\n", data.Conn.Id, data.Conn.endpoint)
	}
	return nil
}

func (e *EventManager) EmitWillClose(conn *Connection, err error) {
	eventData := &IoEventData{
		Conn:    conn,
		Error:   err,
	}
	_ = e.Handle(IoEventWillClose, eventData)
}

func (e *EventManager) EmitClosed(conn *Connection, err error) {
	eventData := &IoEventData{
		Conn:    conn,
		Error:   err,
	}
	_ = e.Handle(IoEventClosed, eventData)
}

func (e *EventManager) EmitDataArrived(conn *Connection, data []byte) {
	eventData := &IoEventData{
		Conn:    conn,
		Data:data,
	}
	_ = e.Handle(IoEventDataArrived, eventData)
}

func (e *EventManager) EmitMessageArrived(conn *Connection, data interface{}) {
	eventData := &IoEventData{
		Conn:    conn,
		Message: data,
	}
	_ = e.Handle(IoEventMessageArrived, eventData)
}

func (e *EventManager) EmitMessageWillEncode(conn *Connection, data []byte) {
	eventData := &IoEventData{
		Conn:    conn,
		Data: data,
	}
	_ = e.Handle(IoEventMessageWillEncode, eventData)
}

func (e *EventManager) EmitMessageEncoded(conn *Connection, data []byte) {
	eventData := &IoEventData{
		Conn:    conn,
		Data: data,
	}
	_ = e.Handle(IoEventMessageEncoded, eventData)
}

func (e *EventManager) EmitDataWillSend(conn *Connection, data []byte) {
	eventData := &IoEventData{
		Conn:    conn,
		Data: data,
	}
	_ = e.Handle(IoEventDataWillSend, eventData)
}

func (e *EventManager) EmitDataSent(conn *Connection, data []byte) {
	eventData := &IoEventData{
		Conn:    conn,
		Data: data,
	}
	_ = e.Handle(IoEventDataSent, eventData)
}

func (e *EventManager) EmitError(conn *Connection, err error) {
	eventData := &IoEventData{
		Conn:    conn,
		Error:   err,
	}
	_ = e.Handle(IoEventError, eventData)
}

func (e *EventManager) EmitAccepted(conn *Connection) error {

	if err := e.Handle(IoEventAccepted, &IoEventData{Conn: conn}); err != nil {
		e.EmitError(conn, err)
		conn.conn.Close()
		e.EmitClosed(conn, err)
		return err
	}

	return nil
}

func (e *EventManager) EmitAcceptFailed(conn *Connection, err error) {
	e.EmitError(conn, err)
	conn.conn.Close()
	e.EmitClosed(conn, err)
}