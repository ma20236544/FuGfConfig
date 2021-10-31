package FileOperations

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

func DownloadFile(Url string, filePath string) {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(HttpProxy)
	}

	httpTransport := &http.Transport{
		Proxy: proxy,
	}

	httpClient := &http.Client{
		Transport: httpTransport,
	}

	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	out, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}
