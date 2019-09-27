package passport

import "go.lsl.digital/lardwaz/sdk/auth"

// Session implements the Session interface
type Session struct {
	user        *auth.User
	permissions map[string]struct{}
}

// NewSession returns a new Session
func NewSession(user auth.User) *Session {
	permsMap := make(map[string]struct{})
	for _, p := range user.Actions {
		permsMap[p] = struct{}{}
	}

	return &Session{
		user:        &user,
		permissions: permsMap,
	}
}

// User returns the user info
func (s Session) User() *auth.User {
	return s.user
}

// IsAllowed returns whether user can perform an action
func (s Session) IsAllowed(action string) bool {
	_, found := s.permissions[action]
	return found
}
