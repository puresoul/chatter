package controller

import (
        "fmt"
        "chatter/router"
        "chatter/datastore"
        "net/http"
        "text/template"
)

type data struct {
    Name string
    Usr []users
    Msg []msg
}

type users struct {
    Name string
    Time string
    Msg  bool
}

type msg struct {
    Me string
    You string
    Done bool
    Time string
}

func Load() {
        router.Get("/", index)
        router.Get("/chat/", chatApi)
	router.Get("/dbg", dbg)
	router.Get("/login/", loginApi)
}

func dbg(w http.ResponseWriter, r *http.Request) {
    ds := datastore.Init()

    ds.Add(datastore.Msg{Author: "u2", User:"u1",Text:"hi"})
    ds.Add(datastore.Msg{Author: "u1", User:"u2",Text:"hi"})
}

func loginApi(w http.ResponseWriter, r *http.Request) {
    t, _ := template.New("login").Parse(loginpage)
    _ = t.Execute(w, "")
}

func chatApi(w http.ResponseWriter, r *http.Request) {
    ds := datastore.Init()

    u := users{Name: "Lol", Time: "Now"}
    m1 := msg{Me: "Hi"}
    m2 := msg{You: "Hou"}
    s := data{Usr: []users{ u }, Msg: []msg{m1,m1,m2,m2,m1,m2}}
    
//    x := r.FormValue("username")
    fmt.Println(ds.Get("u1","u2"))

    t, _ := template.New("chatpage").Parse(chatpage)
    _ = t.Execute(w, s)
}

func index(w http.ResponseWriter, r *http.Request) {
    t, _ := template.New("webpage").Parse(page)
    _ = t.Execute(w, style)
}

