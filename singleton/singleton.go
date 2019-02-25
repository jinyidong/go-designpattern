package singleton

import (
	"fmt"
	"sync"
)

var m *Manager

type Manager struct {
}

var once sync.Once

func GetInstance() *Manager {
	once.Do(func() {
		m = &Manager{}
	})
	return m
}

func (p Manager) Manage() {
	fmt.Println("manage...")
}
