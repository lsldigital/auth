package auth

import "github.com/asdine/storm"

type Storable interface {
	Save(session Sessionable) error
	Get(session Sessionable) (Sessionable, error)
	Delete(session Sessionable) error
}

// Store implements Storable interface
type Store struct {
	DB *storm.DB
}

// Save implements Storable interface
func (s Store) Save(session Sessionable) error {
	return nil
}

// Get implements Storable interface
func (s Store) Get(session Sessionable) (Sessionable, error) {
	return nil, nil
}

// Delete implements Storable interface
func (s Store) Delete(session Sessionable) error {
	return nil
}
