package controller

import (
	"fmt"
	"net/http"
	"text/template"
	"chatter/datastore"
)

const chatpage = `
<html>
<body>
<div class="container" id="chat">
        <div class="left">
            <ul class="people">
                {{with .Usr}}
                {{range .}}
                <li onClick="setName(this)" class="person" id="{{.Name}}">
                    <span class="name">{{.Name}}</span>
                    <span class="time">{{.Time}}</span>
                </li>
                {{end}}
                {{end}}
            </ul>
        </div>
        <div class="right">
            <div class="top"><span><span class="name">NAME</span></span></div>
            <div class="chat">
                <div class="conversation-start">
                    <span>TIME</span>
                </div>
                {{with .Msg}}
                {{range .}}
                {{ if .Me }}
                <div class="bubble me">
                    {{.Me}}
                </div>
                {{end}}
                {{ if .You }}
                <div class="bubble you">
                    {{.You}}
                </div>
               {{end}}
               {{end}}
               {{end}}
            </div>
            <div class="write">
                <input type="text" />
                <a href="javascript:;" class="write-link send"></a>
            </div>
        </div>
</div>
</body>
</html>
`

type data struct {
	Name string
	Usr  []users
	Msg  []msg
}

type users struct {
	Name string
	Time string
	Msg  bool
}

type msg struct {
	Me   string
	You  string
	Done bool
	Time string
}

func chatApi(w http.ResponseWriter, r *http.Request) {
	ds := datastore.Init()

	u := users{Name: "Lol", Time: "Now"}
	m1 := msg{Me: "Hi"}
	m2 := msg{You: "Hou"}
	s := data{Usr: []users{u}, Msg: []msg{m1, m1, m2, m2, m1, m2}}

	//    x := r.FormValue("username")
	fmt.Println(ds.Get("u1", "u2"))

	t, _ := template.New("chatpage").Parse(chatpage)
	_ = t.Execute(w, s)
}
