package auth

type Authable interface {
	Login(username, password string) (Tokenable, error)
	Logout(token Tokenable) error
}

// Auther implements the Authable interface
type Auther struct {
	// TODO: imlementation? graphql? db? sql?
}

// NewAuther return a new Auther
func NewAuther() *Auther {
	return &Auther{}
}

// Login implements the Authable interface
func (a Auther) Login(username string, password string) (Tokenable, error) {
	return nil, nil
}

// Logout implements the Authable interface
func (a Auther) Logout(token Tokenable) error {
	return nil
}
