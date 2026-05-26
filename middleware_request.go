package surf

import (
	"github.com/enetx/g"
)

// defaultUserAgentMW sets the default User-Agent header for surf requests.
// Only sets the header if no User-Agent is already present in the request.
// Uses the predefined _userAgent constant as the default value.
func defaultUserAgentMW(req *Request) error { _ = "STUB: not implemented"; return nil }

// Set the default user-agent header.

// userAgentMW configures a custom User-Agent header for HTTP requests.
// Supports various input types for flexibility:
// - string or g.String: Uses the value directly
// - []string or g.Slice[string]: Randomly selects from the slice (useful for rotation)
// - g.Slice[g.String]: Randomly selects from g.String slice
// Returns an error for unsupported types or empty slices.
func userAgentMW(req *Request, userAgent any) error { _ = "STUB: not implemented"; return nil }

// got101ResponseMW configures request tracing to handle HTTP 101 Switching Protocols responses.
// Sets up client trace callbacks to detect and handle protocol switching responses.
// Returns an error specifically for HTTP 101 responses to allow special handling of protocol upgrades.
// Other 1xx responses are ignored and allowed to proceed normally.
func got101ResponseMW(req *Request) error { _ = "STUB: not implemented"; return nil }

// remoteAddrMW configures request tracing to capture the remote server address.
// Sets up client trace callbacks to extract and store the remote address
// of the server connection for later access. This information can be useful
// for logging, debugging, or connection analysis purposes.
func remoteAddrMW(req *Request) error { _ = "STUB: not implemented"; return nil }

// bearerAuthMW configures Bearer token authentication for HTTP requests.
// Adds an Authorization header with the Bearer token format if a token is provided.
// Only sets the header if the token is not empty, allowing conditional authentication.
func bearerAuthMW(req *Request, token g.String) error { _ = "STUB: not implemented"; return nil }

// basicAuthMW configures HTTP Basic Authentication for requests.
// Expects authentication string in "username:password" format.
// Skips setting auth if Authorization header already exists.
// Returns an error if username or password fields are empty.
func basicAuthMW(req *Request, authentication g.String) error {
	_ = "STUB: not implemented"
	return nil
}

// contentTypeMW configures the Content-Type header for HTTP requests.
// Sets the MIME type of the request body content to inform the server
// how to interpret the request data. Returns an error if contentType is empty.
func contentTypeMW(req *Request, contentType g.String) error { _ = "STUB: not implemented"; return nil }
