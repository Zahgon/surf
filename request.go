package surf

import (
	"context"
	"net"

	"github.com/enetx/g"
	"github.com/enetx/http"
)

// Request represents an HTTP request with additional surf-specific functionality.
// It wraps the standard http.Request and provides enhanced features like middleware support,
// retry capabilities, remote address tracking, and structured error handling.
type Request struct {
	err        error         // General error associated with the request (validation, setup, etc.)
	remoteAddr net.Addr      // Remote server address captured during connection
	bodyBytes  []byte        // Cached body bytes for retry support
	request    *http.Request // The underlying standard HTTP request
	cli        *Client       // The associated surf client for this request
	multipart  *Multipart    // Multipart form data for file uploads and form submissions
}

// GetRequest returns the underlying standard http.Request.
// Provides access to the wrapped HTTP request for advanced use cases.
func (req *Request) GetRequest() *http.Request {
	_ = "STUB: not implemented"

	// Multipart sets multipart form data for the request.
	// The provided Multipart object contains form fields and files to be sent.
	// Returns the request for method chaining. If m is nil, an error is set on the request.
	return nil
}

func (req *Request) Multipart(m *Multipart) *Request { _ = "STUB: not implemented"; return nil }

// prepareMultipart prepares the multipart body for the request.
// It sets up the request body with a pipe reader and configures the Content-Type header.
// Returns an error if both Body() and Multipart() were called, as they are mutually exclusive.
func (req *Request) prepareMultipart() { _ = "STUB: not implemented"; return }

// Do executes the HTTP request and returns a Response wrapped in a Result type.
// This is the main method that performs the actual HTTP request with full surf functionality:
// - Applies request middleware (authentication, headers, tracing, etc.)
// - Preserves request body for potential retries
// - Implements retry logic with configurable status codes and delays
// - Measures request timing for performance analysis
// - Handles request preparation errors and write errors
func (req *Request) Do() g.Result[*Response] {
	_ = "STUB: not implemented"
	// Return early if request has preparation errors
	return nil
}

// Restore body from saved bytes for retry attempts

// Check if retry is needed based on status code and retry configuration

// WithContext associates a context with the request for cancellation and deadlines.
// The context can be used to cancel the request, set timeouts, or pass request-scoped values.
// Returns the request for method chaining. If ctx is nil, the request is unchanged.
func (req *Request) WithContext(ctx context.Context) *Request {
	_ = "STUB: not implemented"
	return nil
}

// AddCookies adds one or more HTTP cookies to the request.
// Cookies are added to the request headers and will be sent with the HTTP request.
// Returns the request for method chaining.
func (req *Request) AddCookies(cookies ...*http.Cookie) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetHeaders sets HTTP headers for the request, replacing any existing headers with the same name.
// Supports multiple input formats:
// - Two arguments: key, value (string or g.String)
// - Single argument: http.Header, Headers, map types, or g.Map types
// Maintains header order for fingerprinting purposes when using g.MapOrd.
// Returns the request for method chaining.
func (req *Request) SetHeaders(headers ...any) *Request { _ = "STUB: not implemented"; return nil }

// AddHeaders adds HTTP headers to the request, appending to any existing headers with the same name.
// Unlike SetHeaders, this method preserves existing headers and adds new values.
// Supports the same input formats as SetHeaders.
// Returns the request for method chaining.
func (req *Request) AddHeaders(headers ...any) *Request { _ = "STUB: not implemented"; return nil }

// applyHeaders is a helper function that processes various header input formats and applies them to an HTTP request.
// It handles type checking, conversion, and delegation to the provided setOrAdd function for actual header manipulation.
// Supports ordered header maps for fingerprinting and maintains compatibility with multiple map and header types.
func (req *Request) applyHeaders(
	rawHeaders []any,
	setOrAdd func(h http.Header, key, value string),
) {
	_ = "STUB: not implemented"
	return
}

// updateRequestHeaderOrder processes ordered headers for HTTP/2 and HTTP/3 fingerprinting.
// It maintains the specific order of headers which is crucial for browser fingerprinting.
// Separates regular headers from pseudo-headers (starting with ':') and sets the appropriate
// header order keys for the transport layer to use. Returns a filtered map containing only
// non-pseudo headers with non-empty values.
func updateRequestHeaderOrder[T ~string](r *Request, h g.MapOrd[T, T]) g.MapOrd[T, T] {
	_ = "STUB: not implemented"
	return nil
}
