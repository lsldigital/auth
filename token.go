package auth

type Tokenable interface {
	GetID() string
	GetUsername() string
	String() string
}

// Token implements the Tokenable interface
type Token struct {
	ID       string
	Username string
}

// GetID implements the Tokenable interface
func (t Token) GetID() string {
	return t.ID
}

// GetUsername implements the Tokenable interface
func (t Token) GetUsername() string {
	return t.Username
}

// String implements the Tokenable and Stringer interfaces
func (t Token) String() string {
	return t.ID + " " + t.Username
}
