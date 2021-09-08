package errors

// Panos is an error returned from PAN-OS.
//
// The error contains both the error message and the code returned from PAN-OS.
type Panos struct {
	Msg  string
	Code int
}

// Error returns the error message.
func (e Panos) Error() string {
	return e.Msg
}

// ObjectNotFound returns true if this is an object not found error.
func (e Panos) ObjectNotFound() bool {
	return e.Code == 7
}

// ObjectNotFound returns an object not found error.
func ObjectNotFound() Panos {
	return Panos{
		Msg:  "Object not found",
		Code: 7,
	}
}
