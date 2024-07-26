// Package static serves static files like CSS, JavaScript, and images.
package controller

import (
	"net/http"
	"text/template"
	"chatter/datastore"
)

const mainPage = `
<!DOCTYPE html>
<html lang="en" >
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1"><link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/meyer-reset/2.0/reset.min.css">
{{.Style}}
</head>

<body>
{{.Content}}
</body>

<script src="https://code.jquery.com/jquery-3.7.0.js"></script>
<script>
{{.Js}}
</script>

</html>
`

type Page struct {
	Style   string
	Js      string
	Content string
}

func returnJs(t string) string {
    return `
   setInterval(function(){
     $( '#chat' ).load('/chat/?t=`+t+` #chat');
   }, 5000);
`
}

func (p *Page) pageApi(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("webpage").Parse(mainPage)

	p.Style = style
	n := r.URL.Query().Get("n")

	if n == "" {
		p.Content = loginpage
		_ = t.Execute(w, p)
		return
	}

	// todo full auth with hash rotation

	ds := datastore.Init()
	if ds.Test(n) {
		p.Js = returnJs("")
	}
	
	p.Content = `<div class="wrapper" id="chat">`

	_ = t.Execute(w, p)
}
