package quickserver

import (
	"net"
	"sync"
)

type SyncInfo struct {
	Sid  string
	Conn *net.Conn
	Cmd  chan []byte
}

type MyMap struct {
	lock *sync.RWMutex
	bm   map[string]SyncInfo
}

func NewMyMap() *MyMap {
	return &MyMap{
		lock: new(sync.RWMutex),
		bm:   make(map[string]SyncInfo),
	}
}

//Get from maps return the k's value
func (m *MyMap) Get(k string) SyncInfo {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if val, ok := m.bm[k]; ok {
		return val
	}
	return SyncInfo{}
}

// Maps the given key and value. Returns false
func (m *MyMap) Set(k string, v SyncInfo) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.bm[k] = v

}

// Returns true if k is exist in the map.
func (m *MyMap) Check(k string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if _, ok := m.bm[k]; !ok {
		return false
	}
	return true
}

func (m *MyMap) Delete(k string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.bm, k)
}
