package cache

import "time"

type Cache struct {
	KeyVal map[string]KeyValue
}

type KeyValue struct {
	Value    string
	Deadline time.Time
}

func NewCache() Cache {
	return Cache{KeyVal: map[string]KeyValue{}}
}

func (c Cache) Get(key string) (string, bool) {
	// for k, v := range c.KeyVal {
	// 	if k == key && v.Deadline.After(time.Now()) {
	// 		return v.Value, true
	// 	}
	// }

	v, ok := c.KeyVal[key]

	if v.Deadline.IsZero() {
		return v.Value, ok
	}
	if time.Now().Before(v.Deadline) {
		return v.Value, ok
	}

	return "", false
}

func (c Cache) Put(key, value string) {
	c.KeyVal[key] = KeyValue{value, time.Time{}}
}

func (c Cache) Keys() []string {
	keys := []string{}
	for k, v := range c.KeyVal {
		if v.Deadline.After(time.Now()) || v.Deadline.Equal(time.Time{}) {
			keys = append(keys, k)
		}
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.KeyVal[key] = KeyValue{value, deadline}
}
