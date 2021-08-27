package models

import "sync"

type data struct {
	mutex sync.RWMutex // multiple readers
	msg   string
}

var (
	CurrentData = data{}
)

func (current *data) Read() string {
	current.mutex.RLock()
	defer current.mutex.RUnlock()

	return current.msg
}

func (current *data) Write(msg string) {
	current.mutex.Lock()
	defer current.mutex.Unlock()

	current.msg = msg
}
