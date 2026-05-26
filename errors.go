package surf

// Custom error types for surf HTTP client operations.
// These errors provide specific information about different failure scenarios
// that can occur during HTTP requests and responses.

type (
	// ErrWebSocketUpgrade indicates that a request received a WebSocket upgrade response.
	// This error is returned when the server responds with HTTP 101 Switching Protocols
	// for WebSocket connections, which require special handling.
	ErrWebSocketUpgrade struct{ Msg string }

	// ErrUserAgentType indicates an invalid user agent type was provided.
	// This error is returned when the user agent parameter is not of a supported type
	// (string, g.String, slices, etc.).
	ErrUserAgentType struct{ Msg string }

	// Err101ResponseCode indicates a 101 Switching Protocols response was received.
	// This error is used to handle HTTP 101 responses that require protocol upgrades.
	Err101ResponseCode struct{ Msg string }

	// ErrHTTP2Fallback indicates that an HTTPS request attempted HTTP/2 first,
	// then tried to fall back to HTTP/1.1, but both attempts failed.
	//
	// Both underlying errors are accessible via Unwrap, enabling
	// errors.Is and errors.As to match against either the HTTP/2 or HTTP/1.1 failure.
	ErrHTTP2Fallback struct {
		HTTP2 error
		HTTP1 error
	}
)

func (e *ErrWebSocketUpgrade) Error() string { _ = "STUB: not implemented"; return "" }

func (e *ErrUserAgentType) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Err101ResponseCode) Error() string { _ = "STUB: not implemented"; return "" }

func (e *ErrHTTP2Fallback) Error() string { _ = "STUB: not implemented"; return "" }

func (e *ErrHTTP2Fallback) Unwrap() []error { _ = "STUB: not implemented"; return nil }
