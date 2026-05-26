package surf

import (
	"github.com/enetx/g"
	"github.com/enetx/http"
)

// Headers represents a collection of HTTP Headers.
type Headers http.Header

// Contains checks if the header contains any of the specified patterns.
// It accepts a header name and a pattern (or list of patterns) and returns a boolean value
// indicating whether any of the patterns are found in the header values.
// The patterns can be a string, a slice of strings, or a slice of *regexp.Regexp.
func (h Headers) Contains(header g.String, patterns any) bool {
	_ = "STUB: not implemented"
	return false
}

// Values returns the values associated with a specified header key.
// It wraps the Values method from the textproto.MIMEHeader type.
func (h Headers) Values(key g.String) g.Slice[g.String] { _ = "STUB: not implemented"; return nil }

// Get returns the first value associated with a specified header key.
// It wraps the Get method from the textproto.MIMEHeader type.
func (h Headers) Get(key g.String) g.String { _ = "STUB: not implemented"; return *new(g.String) }

// Del deletes the values associated with a specified header key.
// It wraps the Del method from the textproto.MIMEHeader type.
func (h Headers) Del(key g.String) { _ = "STUB: not implemented"; return }

// Clone returns a copy of Headers or nil if Headers is nil.
func (h Headers) Clone() Headers { _ = "STUB: not implemented"; return *new(Headers) }
