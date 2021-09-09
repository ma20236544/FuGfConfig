package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

//代理字符串
const (
	HttpProxy  = "http://127.0.0.1:7890"
	SocksProxy = "http://127.0.0.1:7891"
)

//url
const (
	CustomAdRulesListUrl = "https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/ConfigList/CustomAdRulesList.list"
)

func main() {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(HttpProxy)
	}

	httpTransport := &http.Transport{
		Proxy: proxy,
	}

	httpClient := &http.Client{
		Transport: httpTransport,
	}

	req, err := http.NewRequest("GET", CustomAdRulesListUrl, nil)
	if err != nil {
		panic(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	//defer resp.Body.Close()

	out, err := os.Create("CustomAdRulesList.txt")
	if err != nil {
		panic(err)
	}
	//defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}
