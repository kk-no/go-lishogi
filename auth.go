package lishogi

import "net/http"

type Authenticator interface {
	SetAuthentication(r *http.Request) error
}

type AuthMethod func(c *Client)

func SetAccessToken(token string) AuthMethod {
	return func(c *Client) {
		c.authenticator = &AccessToken{
			token: token,
		}
	}
}

type AccessToken struct {
	token string
}

func (a *AccessToken) SetAuthentication(r *http.Request) error {
	r.Header.Set("Authorization", "Bearer "+a.token)
	return nil
}
