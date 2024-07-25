package datastore

import (
	"fmt"
	"sync"
	"time"
)

type StoreFuncs struct {
	user string
	Add  func(NewMsg)
	Get  func(string, string) []NewMsg
	Del  func(string)
	New  func(string, string)
}

type NewMsg struct {
	Author  string
	Reciver string
	Text    string
	Time    string
	Done    bool
}

type msgStore struct {
	reciver string
	text    string
	time    string
	done    bool
}

type userMap map[string][]msgStore

type hashMap map[string]string

var (
	um    userMap
	hm    hashMap
	mutex sync.RWMutex
	df    StoreFuncs
)

func init() {
	mutex.Lock()
	um = make(map[string][]msgStore)
	hm = make(map[string]string)
	mutex.Unlock()
}

func Init() *StoreFuncs {
	df.New = new
	df.Add = add
	df.Get = get
	df.Del = del
	return &df
}

func new(nm, hs string) {
	hm[hs] = nm
	um[nm] = []msgStore{}
}

func add(m NewMsg) {
	if _, ok := um[m.Author]; !ok {
		fmt.Printf("Wrong user\n %#V  \n", m)
		return
	}
	mutex.Lock()
	tp := msgStore{reciver: m.Reciver, text: m.Text, time: fmt.Sprint(time.Now()), done: false}
	um[m.Author] = append(um[m.Author], tp)
	fmt.Printf("%#V \n", um)
	mutex.Unlock()
}

func get(a, u string) []NewMsg {
	var o []NewMsg
	// my msg
	if v, ok := um[a]; ok {
		for _, m := range v {
			if m.reciver == u {
				o = append(o, NewMsg{Reciver: a, Text: m.text, Time: m.time})
			}
		}
	}
	// other
	if v, ok := um[u]; ok {
		for _, m := range v {
			if m.reciver == a {
				o = append(o, NewMsg{Reciver: u, Text: m.text, Time: m.time})
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
