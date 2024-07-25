package controller

import (
	"chatter/datastore"
	"golang.org/x/crypto/bcrypt"
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
	name := r.FormValue("name")
	password := r.FormValue("password")

	hash, err := bcrypt.GenerateFromPassword([]byte(name+password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	ds := datastore.Init()

	ds.New(name, string(hash))

	http.Redirect(w, r, "/?t="+string(hash), http.StatusFound)
}
