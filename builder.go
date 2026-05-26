package surf

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/http"
	"github.com/enetx/surf/profiles"
)

// Builder provides a fluent interface for configuring HTTP clients with various advanced features
// including proxy settings, TLS fingerprinting, HTTP/2 and HTTP/3 support, retry logic,
// redirect handling, and browser impersonation capabilities.
type Builder struct {
	retryCodes               g.Slice[int]                               // HTTP status codes that trigger retries
	proxy                    g.String                                   // Proxy URL for client connections
	cli                      *Client                                    // The client being configured
	checkRedirect            func(*http.Request, []*http.Request) error // Custom redirect policy function
	http2settings            *HTTP2Settings                             // HTTP/2 specific settings
	http3settings            *HTTP3Settings                             // HTTP/3 specific settings
	cliMWs                   *middleware[*Client]                       // Priority-ordered client middlewares
	headersApplier           profiles.HeadersApplier                    // Profile-specific request header pipeline (set by Impersonate)
	retryWait                time.Duration                              // Wait duration between retry attempts
	retryMax                 int                                        // Maximum number of retry attempts
	maxRedirects             int                                        // Maximum number of redirects to follow
	forceHTTP1               bool                                       // Force HTTP/1.1 protocol usage
	forceHTTP2               bool                                       // Force HTTP/2 protocol usage
	forceHTTP3               bool                                       // Force HTTP/3 protocol usage
	cacheBody                bool                                       // Enable response body caching
	followOnlyHostRedirects  bool                                       // Only follow redirects within same host
	forwardHeadersOnRedirect bool                                       // Preserve headers during redirects
	ja                       bool                                       // Enable JA3 TLS fingerprinting
	disableCompression       bool                                       // Disable automatic response body decompression
}

// Build applies all configured settings and returns the client.
// Returns g.Result with error if any middleware fails.
func (b *Builder) Build() g.Result[*Client] { _ = "STUB: not implemented"; return nil }

// With registers middleware into the client builder with optional priority.
//
// It accepts one of the following middleware function types:
//   - func(*surf.Client) error   — client middleware, modifies or initializes the client
//   - func(*surf.Request) error  — request middleware, intercepts or transforms outgoing requests
//   - func(*surf.Response) error — response middleware, intercepts or transforms incoming responses
//
// Parameters:
//   - middleware: A function matching one of the supported middleware types.
//   - priority (optional): Integer priority level. Lower values run earlier. Defaults to 0.
//
// Middleware with the same priority are executed in order of insertion (FIFO).
// If the middleware type is not recognized, With panics with an informative error.
//
// Example:
//
//	// Adding client middleware to modify client settings.
//	.With(func(client *surf.Client) error {
//	    // Custom logic to modify the client settings.
//	    return nil
//	})
//
//	// Adding request middleware to intercept outgoing requests.
//	.With(func(req *surf.Request) error {
//	    // Custom logic to modify outgoing requests.
//	    return nil
//	})
//
//	// Adding response middleware to intercept incoming responses.
//	.With(func(resp *surf.Response) error {
//	    // Custom logic to handle incoming responses.
//	    return nil
//	})
//
// Note: Ensure that middleware functions adhere to the specified function signatures to work correctly with the With method.
func (b *Builder) With(middleware any, priority ...int) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// addCliMW adds a client middleware to the ClientBuilder.
func (b *Builder) addCliMW(m func(*Client) error, priority g.Int) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// addReqMW adds a request middleware to the ClientBuilder.
func (b *Builder) addReqMW(m func(*Request) error, priority g.Int) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// addRespMW adds a response middleware to the ClientBuilder.
func (b *Builder) addRespMW(m func(*Response) error, priority g.Int) *Builder {
	_ = "STUB: not implemented"
	return nil
}

func (b *Builder) Boundary(boundary func() g.String) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// SecureTLS enables TLS certificate verification.
// By default surf skips certificate verification (InsecureSkipVerify=true).
// Call this for production use where certificate validation is required.
func (b *Builder) SecureTLS() *Builder { _ = "STUB: not implemented"; return nil }

// WebSocketGuard enables middleware that blocks WebSocket upgrade responses (HTTP 101).
// Without this, surf allows 101 Switching Protocols to pass through — compatible with
// websocket.Dial and other WebSocket libraries that use surf.Std() as HTTPClient.
// Enable this only if you want to explicitly reject unexpected WebSocket upgrades.
func (b *Builder) WebSocketGuard() *Builder { _ = "STUB: not implemented"; return nil }

// H2C configures the client to handle HTTP/2 Cleartext (h2c).
func (b *Builder) H2C() *Builder { _ = "STUB: not implemented"; return nil }

// HTTP2Settings configures settings related to HTTP/2 and returns an http2s struct.
func (b *Builder) HTTP2Settings() *HTTP2Settings { _ = "STUB: not implemented"; return nil }

// HTTP3Settings configures settings related to HTTP/3 and returns an http3s struct.
func (b *Builder) HTTP3Settings() *HTTP3Settings { _ = "STUB: not implemented"; return nil }

// ForceHTTP3 configures the client to use HTTP/3 forcefully.
func (b *Builder) ForceHTTP3() *Builder { _ = "STUB: not implemented"; return nil }

// Impersonate configures something related to impersonation and returns an impersonate struct.
func (b *Builder) Impersonate() *Impersonate { _ = "STUB: not implemented"; return nil }

// JA configures the client to use a specific TLS fingerprint.
func (b *Builder) JA() *JA { _ = "STUB: not implemented"; return nil }

// UnixSocket sets the path for a Unix domain socket.
// This allows the HTTP client to connect to the server using a Unix domain
// socket instead of a traditional TCP/IP connection.
func (b *Builder) UnixSocket(address g.String) *Builder { _ = "STUB: not implemented"; return nil }

// DNS sets the custom DNS resolver address.
func (b *Builder) DNS(dns g.String) *Builder { _ = "STUB: not implemented"; return nil }

// DNSOverTLS configures the client to use DNS over TLS.
func (b *Builder) DNSOverTLS() *DNSOverTLS { _ = "STUB: not implemented"; return nil }

// Timeout sets the timeout duration for the client.
func (b *Builder) Timeout(timeout time.Duration) *Builder { _ = "STUB: not implemented"; return nil }

// TLSConfig sets a custom TLS configuration for the client.
func (b *Builder) TLSConfig(config *tls.Config) *Builder { _ = "STUB: not implemented"; return nil }

// InterfaceAddr sets the local network interface for outbound connections.
// Accepts either an IP address (e.g., "192.168.1.100", "::1") or an interface name (e.g., "eth0", "en0").
func (b *Builder) InterfaceAddr(address g.String) *Builder { _ = "STUB: not implemented"; return nil }

// Proxy sets the proxy URL for the client.
func (b *Builder) Proxy(proxy g.String) *Builder { _ = "STUB: not implemented"; return nil }

// BasicAuth sets the basic authentication credentials for the client.
func (b *Builder) BasicAuth(authentication g.String) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// BearerAuth sets the bearer token for the client.
func (b *Builder) BearerAuth(authentication g.String) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// UserAgent sets the user agent for the client.
func (b *Builder) UserAgent(userAgent any) *Builder { _ = "STUB: not implemented"; return nil }

// SetHeaders sets headers for the request, replacing existing ones with the same name.
func (b *Builder) SetHeaders(headers ...any) *Builder { _ = "STUB: not implemented"; return nil }

// AddHeaders adds headers to the request, appending to any existing headers with the same name.
func (b *Builder) AddHeaders(headers ...any) *Builder { _ = "STUB: not implemented"; return nil }

// AddCookies adds cookies to the request.
func (b *Builder) AddCookies(cookies ...*http.Cookie) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithContext associates the provided context with the request.
func (b *Builder) WithContext(ctx context.Context) *Builder { _ = "STUB: not implemented"; return nil }

// ContentType sets the content type for the client.
func (b *Builder) ContentType(contentType g.String) *Builder { _ = "STUB: not implemented"; return nil }

// CacheBody configures whether the client should cache the body of the response.
func (b *Builder) CacheBody() *Builder { _ = "STUB: not implemented"; return nil }

// GetRemoteAddress configures whether the client should get the remote address.
func (b *Builder) GetRemoteAddress() *Builder { _ = "STUB: not implemented"; return nil }

// DisableKeepAlive disable keep-alive connections.
func (b *Builder) DisableKeepAlive() *Builder { _ = "STUB: not implemented"; return nil }

// DisableCompression disables automatic response body decompression.
func (b *Builder) DisableCompression() *Builder { _ = "STUB: not implemented"; return nil }

// Retry configures the retry behavior of the client.
//
// Parameters:
//
//	retryMax:  Maximum number of retry attempts. If zero or negative the
//	           retry loop is disabled.
//	retryWait: Minimum wait between retries. If the server responds with a
//	           Retry-After header on a retryable status, the actual pause
//	           becomes max(retryWait, Retry-After). Both legal forms are
//	           honoured: integer delay-seconds and HTTP-date (IMF-fixdate,
//	           RFC 850, ANSI C asctime). Absent or malformed values fall
//	           back to retryWait.
//	codes:     Optional list of HTTP status codes that trigger retries.
//	           If no codes are provided, the defaults are used
//	           (500 Internal Server Error, 429 Too Many Requests,
//	           503 Service Unavailable).
//
// Upper bound for the entire retry cycle, including a long Retry-After
// sleep, is the request context deadline — set it with
// Request.WithContext(ctxWithDeadline). Builder.Timeout only bounds a
// single cli.Do invocation (connect + transmission + body read); it does
// not interrupt the sleep between retries.
func (b *Builder) Retry(retryMax int, retryWait time.Duration, codes ...int) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// ForceHTTP1 configures the client to use HTTP/1.1 forcefully.
func (b *Builder) ForceHTTP1() *Builder { _ = "STUB: not implemented"; return nil }

// ForceHTTP2 configures the client to use HTTP/2 forcefully.
func (b *Builder) ForceHTTP2() *Builder { _ = "STUB: not implemented"; return nil }

// Session configures whether the client should maintain a session.
func (b *Builder) Session() *Builder { _ = "STUB: not implemented"; return nil }

// MaxRedirects sets the maximum number of redirects the client should follow.
func (b *Builder) MaxRedirects(maxRedirects int) *Builder { _ = "STUB: not implemented"; return nil }

// NotFollowRedirects disables following redirects for the client.
func (b *Builder) NotFollowRedirects() *Builder { _ = "STUB: not implemented"; return nil }

// FollowOnlyHostRedirects configures whether the client should only follow redirects within the
// same host.
func (b *Builder) FollowOnlyHostRedirects() *Builder { _ = "STUB: not implemented"; return nil }

// ForwardHeadersOnRedirect adds a middleware to the ClientBuilder object that ensures HTTP headers are
// forwarded during a redirect.
func (b *Builder) ForwardHeadersOnRedirect() *Builder { _ = "STUB: not implemented"; return nil }

// RedirectPolicy sets a custom redirect policy for the client.
func (b *Builder) RedirectPolicy(fn func(*http.Request, []*http.Request) error) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// String generate a string representation of the ClientBuilder instance.
func (b Builder) String() string { _ = "STUB: not implemented"; return "" }
