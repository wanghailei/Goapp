package hello

import "fmt"
import "net/http"

func init () {
	http.HandleFunc("/", handler)
	http.HandleFunc("/login", handleLogin)
}

func handler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "A new start with GAE!")
}

func handleLogin (res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "<h1>Goapp Login</h1>")
}