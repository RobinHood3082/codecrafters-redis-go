package main

import "sync"

type DB struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewDB() *DB {
	return &DB{
		data: make(map[string]string),
	}
}

func (db *DB) Get(key string) (string, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	value, ok := db.data[key]
	return value, ok
}

func (db *DB) Set(key, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[key] = value
}

func (db *DB) Del(key string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.data, key)
}
