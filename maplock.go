package maplock

import "sync"

func New() *MapLock {
	return &MapLock{m: map[string]*sync.Mutex{}}
}

type MapLock struct {
	l sync.Mutex
	m map[string]*sync.Mutex
}

func (m *MapLock) Lock(key string) {
	m.l.Lock()
	l, ok := m.m[key]
	if !ok {
		l = &sync.Mutex{}
		m.m[key] = l
	}
	m.l.Unlock()
	l.Lock()
}

func (m *MapLock) Unlock(key string) {
	m.l.Lock()
	l, ok := m.m[key]
	m.l.Unlock()
	if !ok {
		return
	}
	l.Unlock()
}
