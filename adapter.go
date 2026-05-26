package surf

import (
	_http "net/http"
	"net/url"

	"github.com/enetx/http"
)

// Std returns a standard net/http.Client that wraps the configured surf client.
// This is useful for integrating with third-party libraries that expect a standard net/http.Client
// while preserving most surf features.
//
// Supported features:
//   - JA3/TLS fingerprinting
//   - HTTP/2 settings
//   - Cookies and sessions
//   - Request/Response middleware
//   - Headers (User-Agent, custom headers)
//   - Proxy configuration
//   - Timeout settings
//   - Redirect policies
//   - Impersonate browser headers
//
// Known limitations:
//   - Retry logic is NOT supported (implemented in Request.Do(), not in transport)
//   - Response body caching is NOT supported
//   - Remote address tracking is NOT supported
//   - Request timing information is NOT available
//
// For applications requiring retry logic, consider implementing it at the application level
// or use surf.Client directly for those specific requests.
//
// Example usage:
//
//	surfClient := surf.NewClient().
//		Builder().
//		JA3().Chrome().
//		Session().
//		Build()
//
//	// For libraries expecting net/http.Client
//	stdClient := surfClient.Std()
//
//	botClient := &BaseBotClient{
//		Client: *stdClient,
//	}
func (c *Client) Std() *_http.Client { _ = "STUB: not implemented"; return nil }

// TransportAdapter adapts surf.Client to net/http.RoundTripper
// It uses the full surf pipeline including middleware
type TransportAdapter struct {
	transport http.RoundTripper
	client    *Client
}

func (s *TransportAdapter) CloseIdleConnections() { _ = "STUB: not implemented"; return }

// RoundTrip implements net/http.RoundTripper interface using surf's full pipeline
func (s *TransportAdapter) RoundTrip(req *_http.Request) (*_http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// request converts net/http.Request to github.com/enetx/http.Request.
// It preserves all fields including headers, body, and context while
// adapting to the enetx/http package types.
func request(_req *_http.Request) *http.Request { _ = "STUB: not implemented"; return nil }

// deprecated but kept for compatibility

// response converts github.com/enetx/http.Response to net/http.Response.
// It preserves all response fields including status, headers, and body
// while adapting back to standard net/http types.
func response(resp *http.Response, _req *_http.Request) *_http.Response {
	_ = "STUB: not implemented"
	return nil
}

// CookieJarAdapter adapts github.com/enetx/http.CookieJar to net/http.CookieJar.
// It provides bidirectional cookie conversion between the two HTTP packages,
// ensuring cookies set through either interface work correctly.
type CookieJarAdapter struct{ jar http.CookieJar }

// SetCookies implements http.CookieJar interface.
// It converts standard net/http cookies to enetx/http format and
// delegates to the underlying surf cookie jar.
func (c *CookieJarAdapter) SetCookies(u *url.URL, _cookies []*_http.Cookie) {
	_ = "STUB: not implemented"
	return
}

// Cookies implements http.CookieJar interface.
// It retrieves cookies from the underlying surf cookie jar and
// converts them to standard net/http cookie format.
func (c *CookieJarAdapter) Cookies(u *url.URL) []*_http.Cookie {
	_ = "STUB: not implemented"
	return nil
}

// redirect adapts surf's redirect policy function to work with standard net/http.
// It converts net/http requests to enetx/http format, calls the surf redirect policy,
// and returns the result. This ensures custom redirect policies work correctly
// through the standard http.Client interface.
func redirect(fn func(*http.Request, []*http.Request) error) func(*_http.Request, []*_http.Request) error {
	_ = "STUB: not implemented"
	return nil
}
