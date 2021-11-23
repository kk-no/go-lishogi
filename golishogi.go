package lishogi

import (
	"net/http"
	"net/url"
)

var defaultBaseURL = &url.URL{Scheme: "https", Host: "lishogi.org"}

type Client struct {
	HTTPClient *http.Client

	baseURL       *url.URL
	authenticator Authenticator
}

func NewClient(setAuth AuthMethod) *Client {
	cli := &Client{
		HTTPClient: http.DefaultClient,
		baseURL:    defaultBaseURL,
	}

	setAuth(cli)

	return cli
}
