package auth

import (
	"errors"
	"time"

	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
)

const (
	NoLimit = -1
)

var (
	ErrNotFound = errors.New("record not found")
)

type Storable interface {
	Save(session Sessionable) error
	Get(session Sessionable, limit int) (Sessionables, error)
	Delete(session Sessionable) error
}

// Record is used as the structure for storing
// information in the store
type Record struct {
	SessionID   string `storm:"id,index"`
	SessionType int
	UserID      string `storm:"index"`
	UserAgent   Agentable
	Permissions []string
	OriginID    string `storm:"index"`
	Origin      string
	Expiry      time.Duration
	CreatedAt   time.Time
}

// Store implements the Storable interface
type Store struct {
	db *storm.DB
}

// NewStore returns a new Store
func NewStore(db *storm.DB) *Store {
	return &Store{db: db}
}

// Save implements the Storable interface
func (s Store) Save(session Sessionable) error {
	record := &Record{
		SessionID:   session.SessionID(),
		SessionType: session.SessionType(),
		UserID:      session.UserID(),
		UserAgent:   session.UserAgent(),
		Permissions: session.Permissions(),
		OriginID:    session.OriginID(),
		Origin:      session.Origin(),
		Expiry:      session.Expiry(),
		CreatedAt:   session.CreatedAt(),
	}
	if err := s.db.One("SessionID", session.SessionID(), &record); err == storm.ErrNotFound {
		return s.db.Save(&record)
	}

	return s.db.Update(&record)
}

// Get implements the Storable interface
func (s Store) Get(session Sessionable, limit int) (Sessionables, error) {
	var records []Record

	query := s.db.Select(q.And(getConditions(session)...))
	if limit > 0 {
		query = query.Limit(limit)
	}

	if err := query.Find(&records); err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, ErrNotFound
	}

	var sessions Sessions
	for _, r := range records {
		sessions = append(sessions, Session{
			sessionID:   r.SessionID,
			sessionType: r.SessionType,
			userID:      r.UserID,
			userAgent:   r.UserAgent,
			permissions: r.Permissions,
			originID:    r.OriginID,
			origin:      r.Origin,
			expiry:      r.Expiry,
			createdAt:   r.CreatedAt,
		})
	}

	return sessions, nil
}

// Delete implements the Storable interface
func (s Store) Delete(session Sessionable) error {
	sessions, err := s.Get(session, NoLimit)
	if err != nil {
		return err
	}
	for i := 0; i < sessions.Length(); i++ {
		s.db.DeleteStruct(sessions.Get(i))
	}
	return nil
}

// getConditions returns a list of conditions for indexed fields
func getConditions(session Sessionable) []q.Matcher {
	var (
		conditions []q.Matcher

		sessionID = session.SessionID()
		userID    = session.UserID()
		originID  = session.OriginID()
	)
	if sessionID != "" {
		conditions = append(conditions, q.Eq("SessionID", sessionID))
	}
	if userID != "" {
		conditions = append(conditions, q.Eq("UserID", userID))
	}
	if originID != "" {
		conditions = append(conditions, q.Eq("OriginID", originID))
	}

	return conditions
}
