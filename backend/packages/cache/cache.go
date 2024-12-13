package cache

import (
	"sync"
	"unsafe"
)

type Safe struct {
	mu    sync.RWMutex
	store map[string][]byte
}

func NewSafeMap() *Safe {
	return &Safe{
		store: make(map[string][]byte),
	}
}

func (sm *Safe) Get(key string) ([]byte, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, exists := sm.store[key]
	return value, exists
}

func (sm *Safe) Set(key string, value []byte) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.store[key] = value
}

func (sm *Safe) Del(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.store, key)
}

func (s *Safe) GetKeys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	keys := make([]string, 0, len(s.store))
	for key := range s.store {
		keys = append(keys, key)
	}
	return keys
}

func (sm *Safe) MemoryUsage() int64 {
	var size int64
	size += int64(unsafe.Sizeof(sm.mu))
	size += int64(unsafe.Sizeof(sm.store))
	for k, v := range sm.store {
		size += int64(len(k))
		size += int64(len(v))
	}
	return size
}

func (sm *Safe) MemoryUsageKey(key string) int64 {
	var size int64
	size += int64(len(sm.store[key]))
	return size
}
