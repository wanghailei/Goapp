package hello

import "fmt"
import "html/template"
import "net/http"

const guestbookForm = `
	<html><body><form action="/sign" method="post">
		<div><textarea name="content" rows="3" cols="100"></textarea></div>
		<div><input type="submit" value="Sign Guestbook"></div>
	</form></body></html>`

const signTemplateHTML = `
	<html><body>
		<p>You wrote:</p>
		<pre>{{.}}</pre>
	</body></html>
	`
var signTemplate = template.Must(template.New("sign").Parse(signTemplateHTML))

func init () {
	http.HandleFunc("/", guest)
	http.HandleFunc("/sign", sign)
}
	
func guest (rep http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rep, guestbookForm)
}
	
func sign (rep http.ResponseWriter, req *http.Request) {
	// The sign function gets the form data by calling r.FormValue.
	// Then passes it to signTemplate.Execute that writes the rendered template to the http.ResponseWriter.
	err := signTemplate.Execute(rep, req.FormValue("content"))
	if err != nil {
		http.Error(rep, err.Error(), http.StatusInternalServerError)
	}
}