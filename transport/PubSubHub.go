package transport

import (
	"container/list"
	"sync"
)

type SubFunc func(data *IoEventData)

type PubSubHub struct {
	Subscribed map[IoEventId]*list.List
	lock sync.RWMutex
}

func (p *PubSubHub) Init() {
	p.Subscribed = make(map[IoEventId]*list.List)
}

func (p *PubSubHub) Subscribe(eventId IoEventId, subFunc SubFunc) {

	if subFunc == nil {
		return
	}

	p.lock.Lock(); defer p.lock.Unlock()

	sf, ok := p.Subscribed[eventId]
	if !ok {
		sf = list.New()
		p.Subscribed[eventId] = sf
	}
	sf.PushBack(subFunc)
}

func (p *PubSubHub) Publish(eventId IoEventId, data *IoEventData) {

	p.lock.RLock(); defer p.lock.RUnlock()

	sf, ok := p.Subscribed[eventId]
	if !ok {
		return
	}
	for elem := sf.Front(); elem != nil ; elem = elem.Next() {
		fc := elem.Value.(SubFunc)
		fc(data)
	}
}