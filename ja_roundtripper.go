package surf

import (
	"context"
	"net"

	"github.com/enetx/http"
	"github.com/enetx/http2"

	utls "github.com/refraction-networking/utls"
)

// roundtripper is a higher-level wrapper around HTTP transports, providing
// TLS session resumption and protocol selection.
type roundtripper struct {
	http1tr            *http.Transport
	http1trFallback    *http.Transport
	http2tr            *http2.Transport
	clientSessionCache utls.ClientSessionCache
	ja                 *JA
}

// newRoundTripper creates a new roundtripper wrapping the given base transport
// and using JA configuration.
func newRoundTripper(ja *JA, base http.RoundTripper) http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}

// RoundTrip executes a single HTTP request.
// Optimized for parsing different sites (no per-request allocations).
func (rt *roundtripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Plain HTTP always uses HTTP/1.1 (HTTP/2 requires TLS)

// handleHTTPSRequest handles HTTPS requests with optional HTTP/2 support.
// Reuses pre-built transports to avoid allocations.
func (rt *roundtripper) handleHTTPSRequest(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	// If HTTP/1 is forced, use it directly
	return nil, nil
}

// Try HTTP/2 first

// HTTP/2 failed - fallback to HTTP/1.1

// Restore request body if needed for retry

// Retry with HTTP/1.1

// CloseIdleConnections closes all idle connections.
func (rt *roundtripper) CloseIdleConnections() { _ = "STUB: not implemented"; return }

// buildHTTP2Transport builds a new HTTP/2 transport using settings from builder.
func (rt *roundtripper) buildHTTP2Transport() *http2.Transport {
	_ = "STUB: not implemented"
	return nil
}

// Pre-allocate settings slice to avoid multiple allocations

// dialTLS performs TLS handshake using uTLS with default ALPN (h2, http/1.1).
func (rt *roundtripper) dialTLS(ctx context.Context, network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// dialTLSHTTP2 performs TLS handshake and ensures ALPN selected HTTP/2.
// If ALPN negotiated HTTP/1.1 (or no protocol), this fails before any HTTP/2 bytes are written.
func (rt *roundtripper) dialTLSHTTP2(ctx context.Context, network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// dialTLSHTTP1 performs TLS handshake using uTLS with HTTP/1.1 only ALPN.
// Used for fallback when HTTP/2 connection fails.
func (rt *roundtripper) dialTLSHTTP1(ctx context.Context, network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// tlsHandshake performs a full TLS handshake using uTLS, applying JA fingerprint
// presets and optionally enabling session resumption.
func (rt *roundtripper) tlsHandshake(ctx context.Context, network, addr string, forceHTTP1 bool) (*utls.UConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Apply HTTP/1 ALPN if forced

// supportsResumption checks if a ClientHelloSpec supports TLS session resumption.
func supportsResumption(spec utls.ClientHelloSpec) bool { _ = "STUB: not implemented"; return false }

// Early exit if all TLS 1.3 components are found

// If any TLS 1.3 PSK-related extensions are present,
// session resumption is valid only when all required
// TLS 1.3 resumption indicators are present simultaneously.

// Otherwise, fall back to TLS 1.2 semantics where the presence of
// SessionTicketExtension alone indicates support for session resumption.

// setAlpnProtocolToHTTP1 modifies the given ClientHelloSpec to prefer HTTP/1.1
// by updating or adding the ALPN extension.
func setAlpnProtocolToHTTP1(utlsSpec *utls.ClientHelloSpec) { _ = "STUB: not implemented"; return }
