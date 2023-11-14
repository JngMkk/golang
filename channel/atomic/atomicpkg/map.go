package atomicpkg

import (
	"errors"
	"sync"
)

// 스레드 안정성을 유지하면서 값을 읽고 설정할 수 있도록 뮤텍스 사용
type SafeMap struct {
	m  map[string]string
	mu *sync.RWMutex
}

func NewSafeMap() SafeMap {
	return SafeMap{m: make(map[string]string), mu: &sync.RWMutex{}}
}

func (s *SafeMap) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.m[key] = value
}

func (s *SafeMap) Get(key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if v, ok := s.m[key]; ok {
		return v, nil
	}

	return "", errors.New("key not found")
}
