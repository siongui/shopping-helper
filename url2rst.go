package main

import "net/url"

func url2rst(postURL string) string {
	u, err := url.Parse(postURL)
	if err != nil {
		panic(err)
	}
	u.RawQuery = ""
	if u.Host == "www.buy123.com.tw" {
		return parseBuy123(u.String())
	}
	return u.String()
}
