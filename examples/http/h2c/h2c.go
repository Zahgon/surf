package main

import (
	"fmt"
	"log"

	"github.com/enetx/http2"
	"github.com/enetx/surf"
)

func main() {
	http2.VerboseLogs = true

	go H2CServerUpgrade()

	r := surf.NewClient().Builder().H2C().Build().Unwrap().Get("http://localhost:1010").Do()
	if r.IsErr() {
		log.Fatal(r.Err())
	}

	fmt.Println()
	r.Ok().Debug().Request(true).Response(true).Print()
}

func H2CServerUpgrade() { _ = "STUB: not implemented"; return }
