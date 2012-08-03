package hello

import "appengine"
import "appengine/datastore"
import "appengine/user"
import "html/template"
import "net/http"
import "time"
// import "fmt"

type Greeting struct {
	Author string
	Content string
	Date time.Time
}

const guestbookTemplateHTML = `
	<html><body>
		{{range .}}
			{{with .Author}}
				<p><b>{{.}}</b> wrote:</p>
			{{else}}
				<p>An anonymous person wrote:</p>
			{{end}}
			<pre>{{.Content}}</pre>
		{{end}}
		<form action="/sign" method="post">
			<div><textarea name="content" rows="3" cols="60"></textarea></div>
			<div><input type="submit" value="Sign Guestbook"></div>
		</form>
	</body></html>`
var guestbookTemplate = template.Must(template.New("book").Parse(guestbookTemplateHTML))
								
const signTemplateHTML = `
	<html><body>
		<pre>{{.}}</pre>
	</body></html>`
var signTemplate = template.Must(template.New("sign").Parse(signTemplateHTML))

func init () {
	http.HandleFunc("/", index)
	http.HandleFunc("/sign", sign)
}

func index (rep http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	qry := datastore.NewQuery("Greeting").Order("-Date").Limit(10)
	greetings := make([]Greeting, 0, 10)
	if _, err := qry.GetAll(ctx, &greetings); err != nil {
		http.Error(rep, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := guestbookTemplate.Execute(rep, greetings); err != nil {
		http.Error(rep, err.Error(), http.StatusInternalServerError)
	}
}
	
func sign (rep http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	grt := Greeting {
		Content:req.FormValue("content"),
		Date:time.Now(),
	}
	if usr := user.Current(ctx); usr != nil {
		grt.Author = usr.String()
	}
	_, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Greeting", nil), &grt)
	if err != nil {
		http.Error(rep, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(rep, req, "/", http.StatusFound)
}