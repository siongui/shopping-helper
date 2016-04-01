package main

import (
	"fmt"
	"net/http"
)

var indexHtml = `<!doctype html>
<html>
<head><title>Link to Rst Image</title></head>
<body>
<form action="/post" method="post">
  <input size="100" name="url">
  <button>Send</button>
</form>
</body>
</html>`

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, indexHtml)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	url := r.PostFormValue("url")
	fmt.Fprintf(w, url2rst(url))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/post", postHandler)
	http.ListenAndServe(":8000", nil)
}
