package main

import "sync"

type ItemController struct {
	items []string
	mutex *sync.RWMutex
}

func (c *ItemController) addItem(item string) {
	c.mutex.Lock()
	c.items = append(c.items, item)
	c.mutex.Unlock()
}

func (c *ItemController) readItems() []string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.items
}

func NewItemController() *ItemController {
	return &ItemController{
		items: make([]string, 0),
		mutex: &sync.RWMutex{},
	}
}
