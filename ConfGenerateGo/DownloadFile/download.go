package DownloadFile

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
) //url
const (
	// ios_rule_script
	url1 = "https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/master/rule/QuantumultX/Advertising/Advertising.list"
	// FuGfConfig
	url2 = "https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/ConfigList/CustomAdRulesList.list"
)

func Out1() {
	print("hello")
}

func DownloadFile() {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(HttpProxy)
	}

	httpTransport := &http.Transport{
		Proxy: proxy,
	}

	httpClient := &http.Client{
		Transport: httpTransport,
	}

	req, err := http.NewRequest("GET", url2, nil)
	if err != nil {
		panic(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	out, err := os.Create("CustomAdRulesList.txt")
	if err != nil {
		panic(err)
	}

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}