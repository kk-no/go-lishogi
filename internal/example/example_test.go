package example

import (
	"os"

	"github.com/kk-no/go-lishogi"
)

func ExampleNewClient() {
	token := os.Getenv("LISHOGI_ACCESS_TOKEN")
	cli := lishogi.NewClient(lishogi.SetAccessToken(token))
	_ = cli
	// // Output:
}
