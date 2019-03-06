package auth

// Keyable stores key information
type Keyable interface {
	SessionID() string
	UserID() string
	OriginID() string
}

// Key implements the Keyable interface
type Key struct {
	sessionID string
	userID    string
	originID  string
}

// NewKey returns a new Key
func NewKey(sessionID, userID, originID string) *Key {
	return &Key{sessionID: sessionID, userID: userID, originID: originID}
}

// SessionID implements the Keyable interface
func (k Key) SessionID() string {
	return k.sessionID
}

// UserID implements the Keyable interface
func (k Key) UserID() string {
	return k.userID
}

// OriginID implements the Keyable interface
func (k Key) OriginID() string {
	return k.originID
}
