package controller

import (
	"chatter/router"
	"net/http"
)

const jspage = `
   setInterval(function(){
     $( '#chat' ).load('/chat/ #chat');
   }, 5000);

`

func Load() {
	router.Get("/", index)
	router.Post("/", loginApi)
	router.Get("/chat/", chatApi)
}

func index(w http.ResponseWriter, r *http.Request) {
	d := &Page{}
	d.pageApi(w, r)
}
