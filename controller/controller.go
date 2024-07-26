package controller

import (
	"chatter/router"
//	"chatter/datastore"
	"net/http"
)

func Load() {
	router.Get("/", index)
	router.Post("/", login)
	router.Get("/chat/", chatApi)
}

var (
    p *Page
)

func init() {
	p = &Page{}
}

func index(w http.ResponseWriter, r *http.Request) {
	p.pageApi(w, r)
}

func login(w http.ResponseWriter, r *http.Request) {
	p.loginApi(w, r)
}
