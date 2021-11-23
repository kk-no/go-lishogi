package example

import "github.com/kk-no/go-lishogi"

func ExampleNewClient() {
	cli := lishogi.NewClient(lishogi.SetAccessToken("access_token"))
	_ = cli
}
