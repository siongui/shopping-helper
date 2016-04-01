package main

import (
	"bytes"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"text/template"
)

const rstTmpl = `
.. image:: {{.Image}}
   :alt: {{.Name}}
   :target: {{.Url}}
   :align: center
`

type buy123ProductInfo struct {
	Name        string
	Description string
	Image       string
	Url         string
}

func parseBuy123(url string) string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	jsonBlob := doc.Find("script[type=\"application/ld+json\"]").Text()
	i := buy123ProductInfo{}
	err = json.Unmarshal([]byte(jsonBlob), &i)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("buy123").Parse(rstTmpl)
	if err != nil {
		panic(err)
	}
	var rst bytes.Buffer
	err = tmpl.Execute(&rst, i)
	if err != nil {
		panic(err)
	}

	print(rst.String())
	return rst.String()
}
