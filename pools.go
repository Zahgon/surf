package surf

import (
	"compress/gzip"
	"io"
	"sync"

	"github.com/andybalholm/brotli"
	"github.com/enetx/g"
	"github.com/klauspost/compress/zstd"
)

var (
	// zstdDecoderPool pools zstd.Decoder instances.
	zstdDecoderPool = sync.Pool{
		New: func() any {
			dec, _ := zstd.NewReader(nil)
			return dec
		},
	}

	// gzipReaderPool pools gzip.Reader instances.
	gzipReaderPool = sync.Pool{
		New: func() any {
			return new(gzip.Reader)
		},
	}

	// brotliReaderPool pools brotli.Reader instances.
	brotliReaderPool = sync.Pool{
		New: func() any {
			return brotli.NewReader(nil)
		},
	}
)

// zstdReadCloser wraps a zstd decoder and returns it to the pool on Close.
type zstdReadCloser struct {
	dec *zstd.Decoder
}

// Read reads decompressed data from the decoder.
func (zr *zstdReadCloser) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Close resets the decoder and returns it to the pool.
		nil
}

func (zr *zstdReadCloser) Close() error { _ = "STUB: not implemented"; return nil }

// gzipReadCloser wraps a gzip reader and returns it to the pool on Close.
type gzipReadCloser struct {
	*gzip.Reader
}

// Close closes the reader and returns it to the pool.
func (gr *gzipReadCloser) Close() error { _ = "STUB: not implemented"; return nil }

// brotliReadCloser wraps a brotli reader and returns it to the pool on Close.
type brotliReadCloser struct {
	*brotli.Reader
}

// Close returns the reader to the pool.
func (br *brotliReadCloser) Close() error { _ = "STUB: not implemented"; return nil }

// acquireGzipReader gets a gzip.Reader from the pool and resets it with the provided reader.
// Returns the reader wrapped in gzipReadCloser for automatic pool management.
func acquireGzipReader(r io.Reader) g.Result[io.ReadCloser] { _ = "STUB: not implemented"; return nil }

// acquireBrotliReader gets a brotli.Reader from the pool and resets it with the provided reader.
// Returns the reader wrapped in brotliReadCloser for automatic pool management.
func acquireBrotliReader(r io.Reader) g.Result[io.ReadCloser] {
	_ = "STUB: not implemented"
	return nil
}

// acquireZstdReader gets a zstd.Decoder from the pool and resets it with the provided reader.
// Returns the decoder wrapped in zstdReadCloser for automatic pool management.
func acquireZstdReader(r io.Reader) g.Result[io.ReadCloser] { _ = "STUB: not implemented"; return nil }
