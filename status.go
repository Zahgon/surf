package surf

// StatusCode represents an HTTP status code with convenient classification methods.
// Extends the basic integer status code with methods to easily identify the response category.
type StatusCode int

// IsInformational returns true if the status code is in the informational range [100, 200].
func (s StatusCode) IsInformational() bool { _ = "STUB: not implemented"; return false }

// IsSuccess returns true if the status code indicates a successful response [200, 300].
func (s StatusCode) IsSuccess() bool { _ = "STUB: not implemented"; return false }

// IsRedirection returns true if the status code indicates a redirection [300, 400].
func (s StatusCode) IsRedirection() bool { _ = "STUB: not implemented"; return false }

// IsClientError returns true if the status code indicates a client error [400, 500].
func (s StatusCode) IsClientError() bool { _ = "STUB: not implemented"; return false }

// IsServerError returns true if the status code indicates a server error [500, ∞].
func (s StatusCode) IsServerError() bool {
	_ = "STUB: not implemented"

	// Text returns the textual representation of the status code.
	return false
}

func (s StatusCode) Text() string { _ = "STUB: not implemented"; return "" }
