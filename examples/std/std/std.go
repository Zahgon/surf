package main

import (
	"net/http"

	"github.com/enetx/surf"
)

func main() {
	cli := surf.NewClient().
		Builder().
		// Proxy("socks5://127.0.0.1:1080").
		Impersonate().Firefox().
		Build().
		Unwrap()

	test(cli.Std())
}

func test(client *http.Client) { _ = "STUB: not implemented"; return }
