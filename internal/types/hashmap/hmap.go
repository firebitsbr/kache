package hashmap

import (
	"sync"
)

type HashMap struct {
	m   map[string]string
	mux sync.Mutex
}

func New() *HashMap {
	return &HashMap{m: make(map[string]string)}
}

func (m *HashMap) Set(key, value string) {
	m.mux.Lock()
	defer m.mux.Unlock()

	m.m[key] = value
	return
}

func (m *HashMap) Get(key string) string {
	m.mux.Lock()
	defer m.mux.Unlock()

	if v, ok := m.m[key]; ok {
		return v
	}

	return ""
}

func (m *HashMap) Keys() []string {
	m.mux.Lock()
	defer m.mux.Unlock()

	keys := make([]string, len(m.m))
	i := 0
	for k := range m.m {
		keys[i] = k
		i++
	}

	return keys
}

func (m *HashMap) Fields() []string {
	m.mux.Lock()
	defer m.mux.Unlock()

	paris := make([]string, len(m.m)*2)
	i := 0

	for key, val := range m.m {
		paris[i] = key
		i++
		paris[i] = val
		i++
	}
	return paris
}

func (m *HashMap) Delete(keys ...string) int {
	m.mux.Lock()
	defer m.mux.Unlock()

	deleted := 0
	for _, key := range keys {
		if _, found := m.m[key]; found {
			delete(m.m, key)
			deleted++
		}
	}

	return deleted
}

func (m *HashMap) Exists(key string) int {
	m.mux.Lock()
	defer m.mux.Unlock()
	_, found := m.m[key]

	if found {
		return 1
	}

	return 0
}
