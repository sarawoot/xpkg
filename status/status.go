package status

import (
	"fmt"

	"github.com/sarawoot/xpkg/code"
)

type Status struct {
	Code    code.Code `json:"code,omitempty"`
	Message string    `json:"message,omitempty"`
}

// New returns a Status representing c and msg.
func New(c code.Code, msg string) *Status {
	return &Status{Code: c, Message: msg}
}

// Newf returns New(c, fmt.Sprintf(format, a...)).
func Newf(c code.Code, format string, a ...interface{}) *Status {
	return New(c, fmt.Sprintf(format, a...))
}

// Err returns an error representing c and msg.  If c is OK, returns nil.
func Err(c code.Code, msg string) error {
	if c == code.OK {
		return nil
	}

	return New(c, msg).Err()
}

// Errorf returns Error(c, fmt.Sprintf(format, a...)).
func Errorf(c code.Code, format string, a ...interface{}) error {
	return Err(c, fmt.Sprintf(format, a...))
}

func Code(err error) code.Code {
	// Don't use FromError to avoid allocation of OK status.
	if err == nil {
		return code.OK
	}
	if se, ok := err.(interface {
		GetStatus() *Status
	}); ok {
		return se.GetStatus().GetCode()
	}
	return code.Unknown
}

func (s *Status) Err() error {
	if s.GetCode() == code.OK {
		return nil
	}
	return &Error{s: s}
}

func (s *Status) String() string {
	// return fmt.Sprintf("error: code = %d desc = %s", s.GetCode(), s.GetMessage())
	return s.GetMessage()
}

// Code returns the status code contained in s.
func (s *Status) GetCode() code.Code {
	if s == nil {
		return code.OK
	}
	return s.Code
}

// Message returns the message contained in s.
func (s *Status) GetMessage() string {
	if s == nil {
		return ""
	}
	return s.Message
}
