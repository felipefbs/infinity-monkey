package pubsub

import "sync"

type Token struct {
	Index int
	Word  string
	Found bool
}

type Publisher struct {
	mutex  sync.RWMutex
	subs   map[string][]chan Token
	closed bool
}

func NewInc() *Publisher {
	return &Publisher{
		subs: make(map[string][]chan Token),
	}
}

func (m *Publisher) Subscribe(topic string, buffer int) <-chan Token {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	ch := make(chan Token, buffer)
	m.subs[topic] = append(m.subs[topic], ch)
	return ch
}

func (m *Publisher) Publish(topic string, msg Token) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if m.closed {
		return
	}

	for _, ch := range m.subs[topic] {
		go func(ch chan Token) {
			ch <- msg
		}(ch)
	}
}

func (m *Publisher) Close() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if !m.closed {
		m.closed = true
		for _, subs := range m.subs {
			for _, ch := range subs {
				close(ch)
			}
		}
	}
}
