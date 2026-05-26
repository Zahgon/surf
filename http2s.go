package surf

import (
	"github.com/enetx/http2"
)

// HTTP2Settings represents HTTP/2 settings.
// https://lwthiker.com/networks/2022/06/17/http2-fingerprinting.html
type HTTP2Settings struct {
	priorityFrames       []http2.PriorityFrame
	priorityParam        http2.PriorityParam
	builder              *Builder
	headerTableSize      uint32
	maxConcurrentStreams uint32
	initialWindowSize    uint32
	maxFrameSize         uint32
	maxHeaderListSize    uint32
	connectionFlow       uint32
	initialStreamID      uint32
	noRFC7540Priorities  uint32
	enablePush           uint32
	usePush              bool
}

// InitialStreamID sets the initial stream id for HTTP/2 streams.
func (h *HTTP2Settings) InitialStreamID(id uint32) *HTTP2Settings {
	_ = "STUB: not implemented"
	return nil
}

// HeaderTableSize sets the header table size for HTTP/2 settings.
func (h *HTTP2Settings) HeaderTableSize(size uint32) *HTTP2Settings {
	_ = "STUB: not implemented"
	return nil
}

// EnablePush enables HTTP/2 server push functionality.
func (h *HTTP2Settings) EnablePush(size uint32) *HTTP2Settings {
	_ = "STUB: not implemented"
	return nil
}

// MaxConcurrentStreams sets the maximum number of concurrent streams in HTTP/2.
func (h *HTTP2Settings) MaxConcurrentStreams(size uint32) *HTTP2Settings {
	_ = "STUB: not implemented"
	return nil
}

// InitialWindowSize sets the initial window size for HTTP/2 streams.
func (h *HTTP2Settings) InitialWindowSize(size uint32) *HTTP2Settings {
	_ = "STUB: not implemented"
	return nil
}

// MaxFrameSize sets the maximum frame size for HTTP/2 frames.
func (h *HTTP2Settings) MaxFrameSize(size uint32) *HTTP2Settings {
	_ = "STUB: not implemented"
	return nil
}

// MaxHeaderListSize sets the maximum size of the header list in HTTP/2.
func (h *HTTP2Settings) MaxHeaderListSize(size uint32) *HTTP2Settings {
	_ = "STUB: not implemented"
	return nil
}

// NoRFC7540Priorities disables RFC 7540 priority signaling in HTTP/2.
func (h *HTTP2Settings) NoRFC7540Priorities(size uint32) *HTTP2Settings {
	_ = "STUB: not implemented"
	return nil
}

// ConnectionFlow sets the flow control for the HTTP/2 connection.
func (h *HTTP2Settings) ConnectionFlow(size uint32) *HTTP2Settings {
	_ = "STUB: not implemented"
	return nil
}

// PriorityParam sets the priority parameter for HTTP/2.
func (h *HTTP2Settings) PriorityParam(priorityParam http2.PriorityParam) *HTTP2Settings {
	_ = "STUB: not implemented"
	return nil
}

// PriorityFrames sets the priority frames for HTTP/2.
func (h *HTTP2Settings) PriorityFrames(priorityFrames []http2.PriorityFrame) *HTTP2Settings {
	_ = "STUB: not implemented"
	return nil
}

// Set applies the accumulated HTTP/2 settings.
// It configures the HTTP/2 settings for the surf client.
func (h *HTTP2Settings) Set() *Builder { _ = "STUB: not implemented"; return nil }
