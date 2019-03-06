package auth

// Authable regroups primary functions for auth
// Main interface of auth package
type Authable interface {
	Login(username, password string) (Tokenable, error)
	Logout(sessionID string) error
}

// Auther implements the Authable interface
type Auther struct {
	store Storable
	// TODO: imlementation? graphql? db? sql?
}

// NewAuther return a new Auther
func NewAuther(store Storable) *Auther {
	return &Auther{store: store}
}

// Login implements the Authable interface
func (a Auther) Login(username string, password string) (Tokenable, error) {
	return nil, nil
}

// Logout implements the Authable interface
func (a Auther) Logout(sessionID string) error {
	a.store.Delete(NewKey(sessionID, "", ""))
	return nil
}
