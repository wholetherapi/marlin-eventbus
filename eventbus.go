package eventbus

import (
	"sync"
)

type Handler func(payload any)

type Bus struct {
	mu       sync.RWMutex
	handlers map[string][]Handler
}

func New() *Bus {
	return &Bus{handlers: make(map[string][]Handler)}
}

func (b *Bus) Subscribe(topic string, h Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers[topic] = append(b.handlers[topic], h)
}

func (b *Bus) Publish(topic string, payload any) {
	b.mu.RLock()
	handlers := append([]Handler(nil), b.handlers[topic]...)
	b.mu.RUnlock()
	for _, h := range handlers {
		h(payload)
	}
}
