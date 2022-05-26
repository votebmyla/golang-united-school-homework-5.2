package cache

import "time"

type Cache struct {
	KeyVal []KeyValue
}

type KeyValue struct {
	Key      string
	Value    string
	Deadline time.Time
}

func NewCache() Cache {
	return Cache{}
}

func (c *Cache) Get(key string) (string, bool) {
	now := time.Now()

	for i, v := range c.KeyVal {
		if key == c.KeyVal[i].Key && now.Before(v.Deadline) {
			return v.Value, true
		}
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	deadline := time.Time{}

	for i, v := range c.KeyVal {
		if key == v.Key {
			c.KeyVal[i].Value = value
			c.KeyVal[i].Deadline = deadline
			return
		}
	}

	data := KeyValue{
		Key:      key,
		Value:    value,
		Deadline: deadline,
	}
	c.KeyVal = append(c.KeyVal, data)
}

func (c *Cache) Keys() []string {
	keys := []string{}
	now := time.Now()
	for _, v := range c.KeyVal {
		if now.Before(v.Deadline) || v.Deadline.Equal(time.Time{}) {
			keys = append(keys, v.Key)
		}
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	for _, v := range c.KeyVal {
		if v.Key == key {
			v.Value = value
			v.Deadline = deadline
			return
		}
	}

	data := KeyValue{
		Key:      key,
		Value:    value,
		Deadline: deadline,
	}
	c.KeyVal = append(c.KeyVal, data)
}
