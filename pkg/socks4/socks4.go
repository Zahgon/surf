package socks4

import (
	"bytes"
	"context"
	"net"
	"net/url"

	"golang.org/x/net/proxy"
)

const (
	socksVersion = 0x04
	socksConnect = 0x01
	socksBind    = 0x02

	accessGranted       = 0x5a
	accessRejected      = 0x5b
	accessIdentRequired = 0x5c
	accessIdentFailed   = 0x5d

	minRequestLen = 8
)

var Ident = "nobody@0.0.0.0"

func init() {
	proxy.RegisterDialerType("socks4", func(u *url.URL, d proxy.Dialer) (proxy.Dialer, error) {
		return socks4{url: u, dialer: d}, nil
	})

	proxy.RegisterDialerType("socks4a", func(u *url.URL, d proxy.Dialer) (proxy.Dialer, error) {
		return socks4{url: u, dialer: d}, nil
	})
}

type socks4 struct {
	url    *url.URL
	dialer proxy.Dialer
}

// DialContext implements proxy.ContextDialer interface
func (s socks4) DialContext(ctx context.Context, network, addr string) (c net.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// Use context-aware dialer if available

// close connection later if we got an error

// Set deadline from context

// Check context before handshake

// Dial implements proxy.Dialer interface
func (s socks4) Dial(network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (s socks4) lookupAddr(ctx context.Context, host string) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

func (s socks4) isSocks4a() bool { _ = "STUB: not implemented"; return false }

func (s socks4) parseAddr(addr string) (host string, iport int, err error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

type request struct {
	Host string
	Port int
	IP   net.IP
	Is4a bool

	err error
	buf bytes.Buffer
}

func (r *request) write(b []byte) { _ = "STUB: not implemented"; return }

func (r *request) writeString(s string) { _ = "STUB: not implemented"; return }

func (r *request) writeBigEndian(data any) { _ = "STUB: not implemented"; return }

func (r request) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
