package surf

import (
	"context"
	"net"

	"github.com/enetx/g"
)

// https://adguard-dns.io/kb/general/dns-providers/

// DNSOverTLS is a configuration struct for DNS over TLS settings.
type DNSOverTLS struct{ builder *Builder }

// AdGuard sets up DNS over TLS with AdGuard DNS.
func (dot *DNSOverTLS) AdGuard() *Builder { _ = "STUB: not implemented"; return nil }

// Google sets up DNS over TLS with Google Public DNS.
func (dot *DNSOverTLS) Google() *Builder { _ = "STUB: not implemented"; return nil }

// Cloudflare sets up DNS over TLS with Cloudflare DNS.
func (dot *DNSOverTLS) Cloudflare() *Builder { _ = "STUB: not implemented"; return nil }

// Quad9 sets up DNS over TLS with Quad9 DNS.
func (dot *DNSOverTLS) Quad9() *Builder { _ = "STUB: not implemented"; return nil }

// Switch sets up DNS over TLS with SWITCH DNS.
func (dot *DNSOverTLS) Switch() *Builder { _ = "STUB: not implemented"; return nil }

// CIRAShield sets up DNS over TLS with CIRA Canadian Shield DNS.
func (dot *DNSOverTLS) CIRAShield() *Builder { _ = "STUB: not implemented"; return nil }

// Ali sets up DNS over TLS with AliDNS.
func (dot *DNSOverTLS) Ali() *Builder { _ = "STUB: not implemented"; return nil }

// Quad101 sets up DNS over TLS with Quad101 DNS.
func (dot *DNSOverTLS) Quad101() *Builder { _ = "STUB: not implemented"; return nil }

// SB sets up DNS over TLS with Secure DNS (dot.sb).
func (dot *DNSOverTLS) SB() *Builder { _ = "STUB: not implemented"; return nil }

// Forge sets up DNS over TLS with DNS Forge.
func (dot *DNSOverTLS) Forge() *Builder { _ = "STUB: not implemented"; return nil }

// LibreDNS sets up DNS over TLS with LibreDNS.
func (dot *DNSOverTLS) LibreDNS() *Builder { _ = "STUB: not implemented"; return nil }

// resolver returns a custom net.Resolver that uses a dial function to create a secure connection
// to the DNS server using DNS over TLS.
func (DNSOverTLS) resolver(serverName g.String, addresses ...g.String) *net.Resolver {
	_ = "STUB: not implemented"
	return nil
}

// AddProvider sets up DNS over TLS with a custom DNS provider.
// It configures a custom net.Resolver using the resolver method.
func (dot *DNSOverTLS) AddProvider(serverName g.String, addresses ...g.String) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// dial returns a dial function that establishes a secure connection to a random DNS server address
// from the given list using DNS over TLS.
func dial(serverName string, addresses ...string) func(context.Context, string, string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return nil
}
