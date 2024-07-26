package controller

import (
	"chatter/datastore"
//	"chatter/auth"
	"log"
	"net/http"
)

const loginpage = `
<div class="login" id="chat">
                  <label for="first">
                        Username:
                  </label>
                  <form action="/" method="post">
                  <input type="text" 
                         id="user" 
                         name="user" 
                         placeholder="Enter your Username" required>

                  <label for="password">
                        Password:
                  </label>
                  <input type="password"
                         id="password" 
                         name="password"
                         placeholder="Enter your Password" required>
                  
                  <div class="wrap">
                        <button type="submit">
                              Submit
                        </button>
                  </div>
                  </form>
</div>
`

func (p *Page) loginApi(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		return
	}
	name := r.FormValue("user")
	password := r.FormValue("password")


	ds := datastore.Init()
	t := ds.New(name, password)
	p.Js = returnJS(t)

	http.Redirect(w, r, "/?n="+name, http.StatusFound)
}
