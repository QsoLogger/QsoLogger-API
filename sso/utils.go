package sso

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/QsoLogger/QsoLogger-API/configure"
)

func httpGetByJson(getUrl string) ([]byte, error) {
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

	if configure.CFG.LogLevel >= configure.All {
		if err != nil {
			log.Println("sso.httpGetByJson:", err)
		}
	}

	return resBodyBytes, err
}

func appUrlPrefix(req *http.Request) string {
	scheme := "http"
	if req.TLS == nil {
		// the scheme was HTTP
		scheme = "http"
	} else {
		// the scheme was HTTPS
		scheme = "https"
	}
	urlPrefix := scheme + "://" + req.Host
	return urlPrefix
}

func getSsoLoginUrl(_ http.ResponseWriter, _ *http.Request, app_url string) string {
	app_url_hex := hex.EncodeToString([]byte(app_url))
	login_url := fmt.Sprintf("%s/ssoLogin/%s", authUrlPrefix, app_url_hex)
	return login_url
}
