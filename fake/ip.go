package fake

import (
	"net/http"
	"net/url"
)

const PROXY_SERVER = ""

type ProxyAuth struct {
	License   string
	SecretKey string
}

func (p ProxyAuth) GetProxyClient() http.Client {
	proxyURL, _ := url.Parse("http://" + p.License + ":" + p.SecretKey + "@" + PROXY_SERVER)
	return http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
}
