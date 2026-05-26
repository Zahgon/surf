package drainbody

import (
	"io"
)

// DrainBody reads all of b to memory and returns the bytes and a new ReadCloser.
// It returns an error if the initial slurp of all bytes fails.
// The returned bytes can be reused for retry support.
func DrainBody(b io.ReadCloser) ([]byte, io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return nil, *new(io.ReadCloser), nil
}
