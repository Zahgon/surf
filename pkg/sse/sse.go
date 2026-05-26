package sse

import (
	"io"

	"github.com/enetx/g"
)

// Event represents a Server-Sent Event.
type Event struct {
	ID    g.String // ID uniquely identifies the event.
	Event g.String // Event specifies the type of the event.
	Data  g.String // Data holds the raw JSON data associated with the event as a string.
	Retry g.Int    // Retry indicates the number of retry attempts for the event.
}

// reset resets the event fields to their zero values or initial states.
func (e *Event) reset() { _ = "STUB: not implemented"; return }

// parse parses the event data based on the event type.
func (e *Event) parse(t, data g.String) { _ = "STUB: not implemented"; return }

// Skip checks if the event should be skipped.
func (e *Event) Skip() bool { _ = "STUB: not implemented"; return false }

// Done checks if the event processing is done.
func (e *Event) Done() bool { _ = "STUB: not implemented"; return false }

// Read reads Server-Sent Events (SSE) from the provided reader and calls the provided function for each event.
func Read(reader io.Reader, fn func(event *Event) bool) error {
	_ = "STUB: not implemented"
	return nil
}
