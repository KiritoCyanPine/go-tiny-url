package inmemory

import (
	"github.com/kiritocyanpine/go-tiny-url/persistant"
)

type InMemory struct {
	db map[string]any
}

func CreateDB() *InMemory {
	return &InMemory{
		db: make(map[string]any),
	}
}

func (m *InMemory) Get(key string) (any, error) {
	value, exist := m.db[key]
	if !exist {
		return nil, persistant.ErrKeyNotFound
	}

	return value, nil
}

func (m *InMemory) Set(key string, value any) error {
	_, exist := m.db[key]
	if exist {
		return persistant.ErrKeyCollision
	}

	m.db[key] = value

	return nil
}
