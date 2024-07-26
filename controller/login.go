package controller

import (
	"chatter/datastore"
	"chatter/auth"
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

func loginApi(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		return
	}
	name := r.FormValue("user")
	password := r.FormValue("password")

	// todo token gen, hash gen and login

	_ = auth.Hash("")

	ds := datastore.Init()
	ds.New(name, password)

	http.Redirect(w, r, "/?n="+name, http.StatusFound)
}
