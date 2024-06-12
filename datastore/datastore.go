package datastore

import (
	"fmt"
	"sync"
	"time"
)

type StoreFuncs struct {
	user   string
	Add    func(Msg)
	Get    func(string, string) []Msg
	Del    func(string)
}

type Msg struct {
	Author string
        User  string
	Text   string
	Time  string
	Done  bool
}

type msgStore struct {
        user  string
	text   string
	time  string
	done   bool
}

type userMap map[string][]msgStore

var (
	um    userMap
	mutex sync.RWMutex
	df    StoreFuncs
)

func init() {
	mutex.Lock()
	um = make(map[string][]msgStore)
	mutex.Unlock()
}

func Init() *StoreFuncs {
	df.Add = add
	df.Get = get
	df.Del = del
	return &df
}

func add(m Msg) {
	mutex.Lock()
	tp := msgStore{user: m.User, text: m.Text, time: fmt.Sprint(time.Now()), done: false}
	um[m.Author] = append(um[m.Author], tp)
	fmt.Printf("%#V \n", um)
	mutex.Unlock()
}

func get(a, u string) []Msg {
	var o []Msg
	// my msg
	if v, ok := um[a]; ok {
		for _, m := range v {
			if m.user == u {
				o = append(o, Msg{User: a, Text: m.text, Time: m.time})
			}
		}
	}
	// other
	if v, ok := um[u]; ok {
		for _, m := range v {
			if m.user == a {
				o = append(o, Msg{User: u, Text: m.text, Time: m.time})
			}
		}
	}
	return o
}

func del(k string) {
	mutex.Lock()
	delete(um, k)
	mutex.Unlock()
}
