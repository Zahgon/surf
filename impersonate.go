package surf

import (
	"github.com/enetx/surf/profiles"
)

type Impersonate struct {
	builder *Builder
	os      profiles.OSKey
}

// RandomOS selects a random OS (Windows, macOS, Linux, Android, or iOS) for the impersonate.
func (im *Impersonate) RandomOS() *Impersonate { _ = "STUB: not implemented"; return nil }

// Windows sets the OS to Windows.
func (im *Impersonate) Windows() *Impersonate { _ = "STUB: not implemented"; return nil }

// MacOS sets the OS to macOS.
func (im *Impersonate) MacOS() *Impersonate { _ = "STUB: not implemented"; return nil }

// Linux sets the OS to Linux.
func (im *Impersonate) Linux() *Impersonate { _ = "STUB: not implemented"; return nil }

// Android sets the OS to Android.
func (im *Impersonate) Android() *Impersonate { _ = "STUB: not implemented"; return nil }

// IOS sets the OS to iOS.
func (im *Impersonate) IOS() *Impersonate { _ = "STUB: not implemented"; return nil }

// Chrome impersonates Chrome browser v145.
func (im *Impersonate) Chrome() *Builder { _ = "STUB: not implemented"; return nil }

// Firefox impersonates Firefox browser v148.
func (im *Impersonate) Firefox() *Builder { _ = "STUB: not implemented"; return nil }

// applyVariant materialises a browser and form-factor profile onto the Builder. Profile packages
// own all data (TLS spec, boundary, H2/H3 SETTINGS, header set), this method owns sequencing.
func (im *Impersonate) applyVariant(v profiles.Variant) *Builder {
	_ = "STUB: not implemented"
	return nil
}
