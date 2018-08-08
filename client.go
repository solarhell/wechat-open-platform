package wechatopenplatform

import "net/http"

type Client struct {
	client    *http.Client
	appId     string
	appSecret string
}

func NewClient(httpClient *http.Client, appId, appSecret string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	c := &Client{
		client:    httpClient,
		appId:     appId,
		appSecret: appSecret,
	}

	return c
}
