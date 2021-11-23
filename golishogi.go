package lishogi

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

var (
	defaultBaseURL = &url.URL{Scheme: "https", Host: "lishogi.org"}
	apiBasePath    = "/api"
)

type Client struct {
	httpClient    *http.Client
	baseURL       *url.URL
	authenticator Authenticator

	Team TeamService
}

func NewClient(setAuth AuthMethod) *Client {
	cli := &Client{
		httpClient: http.DefaultClient,
		baseURL:    defaultBaseURL,
	}

	setAuth(cli)

	cli.Team = NewTeamService(apiBasePath, cli)

	return cli
}

func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	var js []byte = nil

	if body != nil {
		js, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	if err := c.authenticator.SetAuthentication(req); err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) CreateAndDo(method, relPath string, data, resource interface{}) error {
	if strings.HasPrefix(relPath, "/") {
		relPath = strings.TrimLeft(relPath, "/")
	}

	req, err := c.NewRequest(method, relPath, data)
	if err != nil {
		return err
	}

	if _, err := c.doGetHeaders(req, resource); err != nil {
		return err
	}

	return nil
}

func (c *Client) doGetHeaders(req *http.Request, v interface{}) (http.Header, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resErr := CheckResponseError(resp); resErr != nil {
		return nil, resErr
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return nil, err
		}
	}
	return resp.Header, nil
}

func CheckResponseError(r *http.Response) error {
	if !isErrorStatusCode(r.StatusCode) {
		return nil
	}
	return errors.New("error cause") // TODO: implements error case
}

func isErrorStatusCode(code int) bool {
	return http.StatusBadRequest <= code
}

func (c *Client) Get(path string, resource interface{}) error {
	return c.CreateAndDo(http.MethodGet, path, nil, resource)
}
