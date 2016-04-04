package main

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"html/template"
	"net/http"
)

type HNews struct {
	Title       string
	URL         string
	CommentsURL string
}

type Index struct {
	Textarea template.HTML
}

var htmlTmpl = `<!doctype html><html>
<head><title>Link to Rst Image</title></head>
<body>
<form action="/" method="post">
  <input size="100" name="url" placeholder="HN comments URL">
  <button>Send</button><br><br>
  <textarea id="ta" rows="5" cols="80">{{.Textarea}}</textarea><br>
  <button type="button" id="copy">Copy textarea to clipboard</button>
</form>
<script>
  var textareaElm = document.getElementById("ta");
  var copyElm = document.getElementById("copy");
  copyElm.onclick = function(event) {
    textareaElm.select();
    var isSuccessful = document.execCommand('copy');
    if (isSuccessful) {
      textareaElm.value = "Copy OK";
    } else {
      textareaElm.value = "Copy Fail";
    }
  }
</script>
</body></html>`

var rstTmpl = ".. [1] `{{.Title}} <{{.URL}}>`_\n" +
	"       (`HN comments <{{.CommentsURL}}>`__)\n"

func processHNURL(url string) string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	news := HNews{CommentsURL: url}
	newsElm := doc.Find("td.title > a")
	news.Title = newsElm.Text()
	news.URL, _ = newsElm.Attr("href")

	tmpl, err := template.New("hn").Parse(rstTmpl)
	if err != nil {
		panic(err)
	}

	var rst bytes.Buffer
	err = tmpl.Execute(&rst, news)
	if err != nil {
		panic(err)
	}

	return rst.String()
}

func handler(w http.ResponseWriter, r *http.Request) {
	idx := Index{}
	if r.Method == "POST" {
		hnurl := r.PostFormValue("url")
		idx.Textarea = template.HTML(processHNURL(hnurl))
	}
	tmpl, err := template.New("index").Parse(htmlTmpl)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, idx)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
