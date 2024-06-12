package datastore

import (
	"fmt"
	"sync"
	"time"
)

type StoreFuncs struct {
	Create func(string)
	Add    func(string, string, string)
	Get    func(string, string) string
	Update func(string, int64) bool
	Del    func(string)
}

type dataStore struct {
	json  string
	key   string
}

type dataMap map[string]dataStore

var (
	dm    dataMap
	mutex sync.RWMutex
	df    dataFuncs
)

func init() {
	mutex.Lock()
	tm = make(map[string]dataStore)
	mutex.Unlock()
}

func Init() *StoreFuncs {
	df.Add = add
	df.Get = get
	df.Del = del
	return &df
}

func add(tk string) {
	mutex.Lock()
	tp := dataStore{json: j, id: i}
	dm[key] = tp
	mutex.Unlock()
}

func get(k string) dataStore {
	if v, ok := dm[k]; ok {
		return v
	}
	return dataStore{}
}

func del(k string) {
	mutex.Lock()
	delete(dm[k])
	mutex.Unlock()
}
