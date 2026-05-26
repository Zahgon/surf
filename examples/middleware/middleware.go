package main

import (
	"fmt"
	"log"

	"github.com/enetx/surf"
)

func main() {
	const url = "https://yahoo.com"

	cli := surf.NewClient().
		Builder().
		With(jar).
		With(dns).
		With(baseURL).
		With(ua2, 2). // Apply `ua2` middleware with priority 2 (executes after `ua`)
		With(ua).     // Apply `ua` middleware with priority 0 (executes before `ua2`)
		Build().
		Unwrap()

	r := cli.Get(url).Do()
	if r.IsErr() {
		log.Fatal(r.Err())
	}

	defer r.Ok().Body.Close()

	fmt.Println(r.Ok().URL)
	fmt.Println(r.Ok().UserAgent)
}

func dns(client *surf.Client) error { _ = "STUB: not implemented"; return nil }

func jar(client *surf.Client) error { _ = "STUB: not implemented"; return nil }

func baseURL(req *surf.Request) error { _ = "STUB: not implemented"; return nil }

func ua(req *surf.Request) error { _ = "STUB: not implemented"; return nil }

func ua2(req *surf.Request) error { _ = "STUB: not implemented"; return nil }
