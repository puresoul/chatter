package controller

import (
	"chatter/router"
//	"chatter/datastore"
	"net/http"
)

func Load() {
	router.Get("/", index)
	router.Post("/", loginApi)
	router.Get("/chat/", chatApi)
}

func index(w http.ResponseWriter, r *http.Request) {
	d := &Page{}
	d.pageApi(w, r)
}
