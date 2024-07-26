package datastore

import (
	"chatter/auth"
	"fmt"
	"sync"
	"time"
)

type StoreFuncs struct {
	user string
	Add  func(NewMsg)
	Get  func(string, string) []NewMsg
	Del  func(string)
	New  func(string, string) string
	Test func(string) bool
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

type hashStore struct {
	name  string
	token string
	time  int64
}

type userMap map[string][]msgStore

type hashMap map[string]hashStore

var (
	um    userMap
	hm    hashMap
	mutex sync.RWMutex
	df    StoreFuncs
)

func init() {
	mutex.Lock()
	um = make(map[string][]msgStore)
	hm = make(map[string]hashStore)
	mutex.Unlock()
}

func Init() *StoreFuncs {
	df.Test = test
	df.New = new
	df.Add = add
	df.Get = get
	df.Del = del
	return &df
}

func new(nm, ps string) string {
	now := time.Now()
	sec := now.Unix()
	hs := auth.Hash(nm+ps)
	t := auth.Encode(fmt.Sprint(sec) + " " + hs)
	mutex.Lock()
	hm[hs] = hashStore{name: nm, token: t, time: sec}
	um[nm] = []msgStore{}
	mutex.Unlock()
	fmt.Printf("\n hashStore: %#v \n msgStore: %#v \n", hm, um)
	return t
}

func test(t string) bool {
	i, h := auth.Decode(t)
	if v, ok := hm[h]; ok {
		if i == fmt.Sprint(v.time) {
			now := time.Now()
			sec := now.Unix()
			if v.time+300 < sec {
				v.token = auth.Encode(fmt.Sprint(sec)+ " " + h)
				v.time = sec			}
			return false
		}
	}
	return true
}

func add(m NewMsg) {
	if _, ok := um[m.Author]; !ok {
		fmt.Printf("Wrong user\n %#v  \n", m)
		return
	}
	mutex.Lock()
	tp := msgStore{reciver: m.Reciver, text: m.Text, time: fmt.Sprint(time.Now()), done: false}
	um[m.Author] = append(um[m.Author], tp)
	fmt.Printf("%#v \n", um)
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
	delete(hm, k)
	mutex.Unlock()
}
