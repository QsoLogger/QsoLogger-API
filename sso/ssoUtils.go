package sso

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func httpGet(getUrl string) ([]byte, error) {
	var resBodyBytes []byte
	var err error
	var res *http.Response

	req, err := http.NewRequest("GET", getUrl, nil)
	if err == nil {
		req.Header.Set("Content-Type", "application/json")
	}

	client := http.Client{}
	client.Timeout = 2 * time.Second

	if len(proxy) > 0 {
		var proxyUrl *url.URL
		proxyUrl, err = url.Parse(proxy)
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	}

	if err == nil {
		res, err = client.Do(req)
	}

	if err == nil {
		if res.StatusCode == http.StatusOK {
			resBodyBytes, err = io.ReadAll(res.Body)
			res.Body.Close()
		} else {
			msg := fmt.Sprintf("http respons code is %d", res.StatusCode)
			err = errors.New(msg)
		}
	}

	if err != nil {
		log.Println(err)
	}

	return resBodyBytes, err
}
