package controller

import (
//        "fmt"
        "chatter/router"
        "net/http"
        "text/template"
)

type data struct {
    Style string
}

func Load() {
        router.Get("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
    s := data{Style: style}
    t, _ := template.New("webpage").Parse(page)
    _ = t.Execute(w, s)
}

