package surf

import (
	"github.com/enetx/http2"
	"github.com/enetx/surf/profiles"
)

// h2adapter wraps *HTTP2Settings to satisfy profiles.H2Config (which returns the interface type
// instead of *HTTP2Settings, so direct method satisfaction is impossible). Each method delegates
// to the underlying *HTTP2Settings and returns the adapter to keep the chain fluent.
type h2adapter struct{ s *HTTP2Settings }

func (a h2adapter) HeaderTableSize(v uint32) profiles.H2Config {
	_ = "STUB: not implemented"
	return *new(profiles.H2Config)
}

func (a h2adapter) EnablePush(v uint32) profiles.H2Config {
	_ = "STUB: not implemented"
	return *new(profiles.H2Config)
}

func (a h2adapter) MaxConcurrentStreams(v uint32) profiles.H2Config {
	_ = "STUB: not implemented"
	return *new(profiles.H2Config)
}

func (a h2adapter) InitialWindowSize(v uint32) profiles.H2Config {
	_ = "STUB: not implemented"
	return *new(profiles.H2Config)
}

func (a h2adapter) MaxFrameSize(v uint32) profiles.H2Config {
	_ = "STUB: not implemented"
	return *new(profiles.H2Config)
}

func (a h2adapter) MaxHeaderListSize(v uint32) profiles.H2Config {
	_ = "STUB: not implemented"
	return *new(profiles.H2Config)
}

func (a h2adapter) NoRFC7540Priorities(v uint32) profiles.H2Config {
	_ = "STUB: not implemented"
	return *new(profiles.H2Config)
}

func (a h2adapter) ConnectionFlow(v uint32) profiles.H2Config {
	_ = "STUB: not implemented"
	return *new(profiles.H2Config)
}

func (a h2adapter) InitialStreamID(v uint32) profiles.H2Config {
	_ = "STUB: not implemented"
	return *new(profiles.H2Config)
}

func (a h2adapter) PriorityParam(v http2.PriorityParam) profiles.H2Config {
	_ = "STUB: not implemented"
	return *new(profiles.H2Config)
}

func (a h2adapter) PriorityFrames(v []http2.PriorityFrame) profiles.H2Config {
	_ = "STUB: not implemented"
	return *new(profiles.H2Config)
}

// h3adapter wraps *HTTP3Settings to satisfy profiles.H3Config.
type h3adapter struct{ s *HTTP3Settings }

func (a h3adapter) QpackMaxTableCapacity(v uint64) profiles.H3Config {
	_ = "STUB: not implemented"
	return *new(profiles.H3Config)
}

func (a h3adapter) MaxFieldSectionSize(v uint64) profiles.H3Config {
	_ = "STUB: not implemented"
	return *new(profiles.H3Config)
}

func (a h3adapter) QpackBlockedStreams(v uint64) profiles.H3Config {
	_ = "STUB: not implemented"
	return *new(profiles.H3Config)
}

func (a h3adapter) EnableConnectProtocol(v uint64) profiles.H3Config {
	_ = "STUB: not implemented"
	return *new(profiles.H3Config)
}

func (a h3adapter) SettingsH3Datagram(v uint64) profiles.H3Config {
	_ = "STUB: not implemented"
	return *new(profiles.H3Config)
}

func (a h3adapter) H3Datagram(v uint64) profiles.H3Config {
	_ = "STUB: not implemented"
	return *new(profiles.H3Config)
}

func (a h3adapter) EnableWebtransport(v uint64) profiles.H3Config {
	_ = "STUB: not implemented"
	return *new(profiles.H3Config)
}

func (a h3adapter) Grease() profiles.H3Config {
	_ = "STUB: not implemented"
	return *new(profiles.H3Config)
}
