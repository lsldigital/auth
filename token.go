package auth

// Tokenable stores token information
type Tokenable interface {
	ID() string
	Username() string
	HasPermission(permission string) bool
	String() string
}

// Token implements the Tokenable interface
type Token struct {
	id          string
	username    string
	permissions []string
}

// NewToken returns a new Token
func NewToken(id string, username string, permissions []string) *Token {
	return &Token{
		id:          id,
		username:    username,
		permissions: permissions,
	}
}

// ID implements the Tokenable interface
func (t Token) ID() string {
	return t.id
}

// Username implements the Tokenable interface
func (t Token) Username() string {
	return t.username
}

// HasPermission implements the Tokenable interface
func (t Token) HasPermission(permission string) bool {
	for _, p := range t.permissions {
		if p == permission {
			return true
		}
	}
	return false
}

// String implements the Tokenable and Stringer interfaces
func (t Token) String() string {
	return t.id + " " + t.username
}
