package controller

import (
        "fmt"
        "chatter/router"
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
}

type msg struct {
    Me string
    You string
    Time string
}

func Load() {
        router.Get("/", index)
        router.Get("/chat/", chatapi)
}

func chatapi(w http.ResponseWriter, r *http.Request) {
    u := users{Name: "Lol", Time: "Now"}
    m1 := msg{Me: "Hi"}
    m2 := msg{You: "Hou"}
    s := data{Usr: []users{ u }, Msg: []msg{m1,m1,m2,m2,m1,m2}}
    
    x := r.FormValue("username")
    fmt.Println(x)

    t, _ := template.New("chatpage").Parse(chatpage)
    _ = t.Execute(w, s)
}

func index(w http.ResponseWriter, r *http.Request) {
    t, _ := template.New("webpage").Parse(page)
    _ = t.Execute(w, style)
}

