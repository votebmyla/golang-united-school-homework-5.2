package cache

import "time"

type Cache struct {
	KeyVal map[string]KeyValue
}

type KeyValue struct {
	Value    string
	Deadline time.Time
}

func NewCache(c Cache) Cache {
	cache := Cache{KeyVal: make(map[string]KeyValue)}
	return cache
}

func (c *Cache) Get(key string) (string, bool) {
	for k, v := range c.KeyVal {
		if k == key && v.Deadline.After(time.Now()) {
			return v.Value, true
		}
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	deadline := time.Time{}
	for k := range c.KeyVal {
		if k == key {
			c.KeyVal[key] = KeyValue{value, deadline}
			return
		}
	}
	c.KeyVal[key] = KeyValue{value, deadline}
}

func (c *Cache) Keys() []string {
	keys := []string{}
	for k, v := range c.KeyVal {
		if v.Deadline.After(time.Now()) || v.Deadline.Equal(time.Time{}) {
			keys = append(keys, k)
		}
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.KeyVal[key] = KeyValue{value, deadline}
}
