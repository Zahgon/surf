package surf

import (
	"io"

	"github.com/enetx/g"
)

// Multipart represents multipart form data with fields and files.
type Multipart struct {
	fields g.MapOrd[g.String, g.String]
	files  g.Slice[*MultipartFile]
	retry  bool
}

// MultipartFile represents a single file for multipart upload.
type MultipartFile struct {
	fieldName   g.String
	fileName    g.String
	contentType g.String
	file        *g.File
	reader      io.Reader
}

// NewMultipart creates a new empty Multipart object.
func NewMultipart() *Multipart { _ = "STUB: not implemented"; return nil }

// Field adds a form field to the multipart.
func (m *Multipart) Field(name, value g.String) *Multipart { _ = "STUB: not implemented"; return nil }

// File adds a physical file to the multipart.
func (m *Multipart) File(fieldName g.String, file *g.File) *Multipart {
	_ = "STUB: not implemented"
	return nil
}

// FileReader adds a file from io.Reader to the multipart.
func (m *Multipart) FileReader(fieldName, fileName g.String, reader io.Reader) *Multipart {
	_ = "STUB: not implemented"
	return nil
}

// FileString adds a file from string content to the multipart.
func (m *Multipart) FileString(fieldName, fileName, content g.String) *Multipart {
	_ = "STUB: not implemented"
	return nil
}

// FileBytes adds a file from byte slice to the multipart.
func (m *Multipart) FileBytes(fieldName, fileName g.String, data g.Bytes) *Multipart {
	_ = "STUB: not implemented"
	return nil
}

// ContentType sets the content type for the last added file.
// Must be called immediately after File/FileReader/FileString/FileBytes.
func (m *Multipart) ContentType(ct g.String) *Multipart { _ = "STUB: not implemented"; return nil }

// FileName overrides the filename for the last added file.
// Useful when you want a different name than the physical file.
func (m *Multipart) FileName(name g.String) *Multipart { _ = "STUB: not implemented"; return nil }

// Retry controls whether the multipart body should be buffered in memory
// to support retries on status codes (429, 503, 5xx, etc.).
//
// When set, the body is fully read into memory before sending,
// allowing the client to replay it on retry.
//
// Recommended only for small requests (≤ 5–10 MB).
func (m *Multipart) Retry() *Multipart { _ = "STUB: not implemented"; return nil }

// prepareWriter writes the multipart data to a writer and returns the content type and write error.
func (m *Multipart) prepareWriter(boundary func() g.String) (io.ReadCloser, string, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}

func escapeQuotes(s g.String) string { _ = "STUB: not implemented"; return "" }
