// Package static serves static files like CSS, JavaScript, and images.
package controller

import (
	"net/http"
	"text/template"
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

func (p *Page) pageApi(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("webpage").Parse(mainPage)

	p.Style = style

	if r.URL.Query().Get("t") == "" {
		p.Content = loginpage
		_ = t.Execute(w, p)
		return
	}

	p.Content = `<div class="wrapper" id="chat">`
	p.Js = jspage

	_ = t.Execute(w, p)
}
