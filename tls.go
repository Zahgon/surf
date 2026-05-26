package surf

import (
	"crypto/tls"
)

// TLSData represents information about a TLS certificate.
type TLSData struct {
	CommonName               []string // List of common names associated with the certificate.
	DNSNames                 []string // List of DNS names associated with the certificate.
	Emails                   []string // List of email addresses associated with the certificate.
	IssuerCommonName         []string // List of common names of the certificate issuer.
	IssuerOrg                []string // List of organizations of the certificate issuer.
	Organization             []string // List of organizations associated with the certificate.
	ExtensionServerName      string   // Server name extension of the TLS certificate.
	FingerprintSHA256        string   // SHA-256 fingerprint of the certificate.
	FingerprintSHA256OpenSSL string   // SHA-256 fingerprint compatible with OpenSSL.
	TLSVersion               string   // TLS version used.
}

// tlsGrabber takes a TLS connection state and returns a tlsData struct containing information
// about the TLS connection.
func tlsGrabber(cs *tls.ConnectionState) *TLSData { _ = "STUB: not implemented"; return nil }

// calculateFingerprints takes a TLS connection state and returns the SHA256 fingerprint
// of the first certificate in the chain or an error if no certificates are found.
func calculateFingerprints(c *tls.ConnectionState) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// openSSL takes a byte slice of a fingerprint and returns a string representation in the OpenSSL
// format.
func openSSL(fpBytes []byte) string { _ = "STUB: not implemented"; return "" }
