package cache

import (
	"reader/model"
	"sync"
)

type Cache struct {
	data map[string]model.Message
	mx   *sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]model.Message),
	}
}

func (c *Cache) Set(message model.Message) {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.data[message.OrderUid] = message
}

func (c *Cache) Get(id string) (model.Message, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()

	msg, ok := c.data[id]
	return msg, ok
}
