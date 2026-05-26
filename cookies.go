package surf

import (
	"github.com/enetx/http"
)

// Cookies represents a list of HTTP Cookies.
type Cookies []*http.Cookie

// Contains checks if the cookies collection contains a cookie that matches the provided pattern.
// The pattern parameter can be either a string or a pointer to a regexp.Regexp object.
// The method returns true if a matching cookie is found and false otherwise.
func (cs Cookies) Contains(pattern any) bool { _ = "STUB: not implemented"; return false }
