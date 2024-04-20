package status

type Error struct {
	s *Status
}

func (e *Error) Error() string {
	return e.s.String()
}

// GRPCStatus returns the Status represented by se.
func (e *Error) GetStatus() *Status {
	return e.s
}

// Is implements future error.Is functionality.
// A Error is equivalent if the code and message are identical.
func (e *Error) Is(target error) bool {
	tse, ok := target.(*Error)
	if !ok {
		return false
	}
	return e.s.Code == tse.s.Code
}
