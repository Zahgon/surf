package surf

import (
	"github.com/enetx/g"
)

// Debug is a struct that holds debugging information for an HTTP response.
type Debug struct {
	print g.Builder // Debug information text.
	resp  Response  // Associated Response.
}

// Debug returns a debug instance associated with a Response.
func (resp Response) Debug() *Debug { _ = "STUB: not implemented"; return nil }

// Print prints the debug information.
func (d *Debug) Print() { _ = "STUB: not implemented"; return }

// Request appends the request details to the debug information.
func (d *Debug) Request(verbos ...bool) *Debug { _ = "STUB: not implemented"; return nil }

// Response appends the response details to the debug information.
func (d *Debug) Response(verbos ...bool) *Debug { _ = "STUB: not implemented"; return nil }
