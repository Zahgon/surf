package surf

// webSocketUpgradeErrorMW detects and handles WebSocket upgrade responses.
// Returns an error when a response indicates a successful WebSocket protocol upgrade
// (HTTP 101 Switching Protocols with Upgrade: websocket header).
// This allows special handling of WebSocket connections which require different processing.
func webSocketUpgradeErrorMW(r *Response) error { _ = "STUB: not implemented"; return nil }

// decodeBodyMW automatically decompresses response bodies based on Content-Encoding header.
// Supports multiple compression algorithms:
// - deflate: DEFLATE compression (zlib format)
// - gzip: GZIP compression
// - br: Brotli compression
// - zstd: Zstandard compression
// Updates the response body reader to provide decompressed content transparently.
// Returns an error if decompression fails, otherwise the body can be read normally.
func decodeBodyMW(r *Response) error { _ = "STUB: not implemented"; return nil }
