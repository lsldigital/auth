package auth

type Authable interface {
	Login(username, password string) (Tokenable, error)
	Logout(token Tokenable) error
}

// Auther implements the Authable interface
type Auther struct{}

// Login implements the Authable interface
func (a Auther) Login(username, password string) (Tokenable, error) {
	return nil, nil
}

// Logout implements the Authable interface
func (a Auther) Logout(token Tokenable) error {
	return nil
}
