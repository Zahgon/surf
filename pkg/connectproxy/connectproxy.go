package connectproxy

import (
	"context"
	"io"
	"net"
	"net/url"
	"sync"

	"github.com/enetx/http"
	"github.com/enetx/http2"
	_ "github.com/enetx/surf/pkg/socks4"
)

type (
	ErrProxyURL      struct{ Msg string }
	ErrProxyStatus   struct{ Msg string }
	ErrPasswordEmpty struct{ Msg string }
	ErrProxyEmpty    struct{}
)

func (e *ErrProxyURL) Error() string      { _ = "STUB: not implemented"; return "" }
func (e *ErrProxyStatus) Error() string   { _ = "STUB: not implemented"; return "" }
func (e *ErrPasswordEmpty) Error() string { _ = "STUB: not implemented"; return "" }
func (e *ErrProxyEmpty) Error() string    { _ = "STUB: not implemented"; return "" }

type proxyDialer struct {
	ProxyURL      *url.URL
	DefaultHeader http.Header

	// overridden dialer allow to control establishment of TCP connection
	Dialer net.Dialer

	// DialTLSContext allows user to control establishment of TLS connection.
	// MUST return connection with completed Handshake, and NegotiatedProtocol.
	DialTLSContext func(ctx context.Context, network, address string) (net.Conn, string, error)

	h2Mu   sync.Mutex
	h2Conn *http2.ClientConn
	conn   net.Conn

	tr2Once sync.Once
	tr2     *http2.Transport
}

// SetResolver sets a custom DNS resolver for the proxy dialer.
// This resolver will be used for all DNS lookups including proxy server address
// and target host resolution. When set, target hostnames are pre-resolved locally
// before being sent to the proxy, ensuring DNS queries bypass the proxy.
func (c *proxyDialer) SetResolver(r *net.Resolver) { _ = "STUB: not implemented"; return }

// dialerProxy is an adapter that implements proxy.Dialer interface
// using net.Dialer to support custom DNS resolver with proxies.
type dialerProxy struct {
	dialer *net.Dialer
}

func (d *dialerProxy) Dial(network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (d *dialerProxy) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

const (
	schemeHTTP  = "http"
	schemeHTTPS = "https"
	socks4      = "socks4"
	socks4A     = "socks4a"
	socks5      = "socks5"
	socks5H     = "socks5h"
)

func NewDialer(proxy string) (*proxyDialer, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *proxyDialer) Dial(network, address string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

type ContextKeyHeader struct{}

func (c *proxyDialer) connectHTTP1(req *http.Request, conn net.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *proxyDialer) connectHTTP2(
	req *http.Request,
	conn net.Conn,
	h2clientConn *http2.ClientConn,
	closeOnError bool,
) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (c *proxyDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// Pre-resolve DNS locally if custom resolver is configured.

func (c *proxyDialer) initProxyConn(ctx context.Context, network string) (net.Conn, string, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), "", nil
}

func (c *proxyDialer) connect(req *http.Request, conn net.Conn, negotiatedProtocol string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (c *proxyDialer) Close() error { _ = "STUB: not implemented"; return nil }

func newHTTP2Conn(c net.Conn, pipedReqBody *io.PipeWriter, respBody io.ReadCloser) net.Conn {
	_ = "STUB: not implemented"
	return *new(net.Conn)
}

type http2Conn struct {
	net.Conn
	in       *io.PipeWriter
	out      io.ReadCloser
	ownsConn bool
}

func (h *http2Conn) Close() error { _ = "STUB: not implemented"; return nil }

func (h *http2Conn) Read(p []byte) (n int, err error)  { _ = "STUB: not implemented"; return 0, nil }
func (h *http2Conn) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
func (h *http2Conn) CloseConn() error                  { _ = "STUB: not implemented"; return nil }
func (h *http2Conn) CloseWrite() error                 { _ = "STUB: not implemented"; return nil }
func (h *http2Conn) CloseRead() error                  { _ = "STUB: not implemented"; return nil }
