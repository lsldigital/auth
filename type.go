package auth

import "time"

// Session types
const (
	TypeSessionUnknown Type = iota
	TypeSessionRedirection
	TypeSessionAuthorization
)

// Type is a session type
type Type int

// Timeout returns timeout for a session type
func (t Type) Timeout() time.Duration {
	switch t {
	case TypeSessionRedirection:
		return 15 * time.Minute
	case TypeSessionAuthorization:
		return (24 * time.Hour) * 30 // 30 days
	}

	return 0
}

// String implements the Stringer interface
func (t Type) String() string {
	switch t {
	case TypeSessionRedirection:
		return "redirection"
	case TypeSessionAuthorization:
		return "authorization"
	}

	return ""
}
