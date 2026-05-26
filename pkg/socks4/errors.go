package socks4

type (
	ErrWrongNetwork    struct{}
	ErrConnRejected    struct{}
	ErrIdentRequired   struct{}
	ErrDialFailed      struct{ err error }
	ErrBuffer          struct{ err error }
	ErrIO              struct{ err error }
	ErrInvalidResponse struct{ resp byte }

	ErrWrongAddr struct {
		msg string
		err error
	}

	ErrHostUnknown struct {
		msg string
		err error
	}
)

func (e *ErrDialFailed) Error() string { _ = "STUB: not implemented"; return "" }
func (e *ErrDialFailed) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *ErrHostUnknown) Error() string { _ = "STUB: not implemented"; return "" }

func (e *ErrHostUnknown) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *ErrBuffer) Error() string { _ = "STUB: not implemented"; return "" }
func (e *ErrBuffer) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *ErrIO) Error() string { _ = "STUB: not implemented"; return "" }
func (e *ErrIO) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *ErrWrongAddr) Error() string { _ = "STUB: not implemented"; return "" }
func (e *ErrWrongAddr) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *ErrWrongNetwork) Error() string { _ = "STUB: not implemented"; return "" }

func (e *ErrConnRejected) Error() string    { _ = "STUB: not implemented"; return "" }
func (e *ErrIdentRequired) Error() string   { _ = "STUB: not implemented"; return "" }
func (e *ErrInvalidResponse) Error() string { _ = "STUB: not implemented"; return "" }
