package auth

import (
	"errors"
	"time"

	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
)

// Storable constants
const (
	NoLimit = -1
)

// Storable errors
var (
	ErrNotFound = errors.New("record not found")
)

// Storable stores session information in <any> store
type Storable interface {
	Save(session Sessionable) error
	Get(key Keyable, limit int) (Sessionables, error)
	Delete(key Keyable) error
	Cleanup() error
}

// Record defines the structure for storing
// information in the store
type Record struct {
	SessionID   string `storm:"id"`
	SessionType Type
	UserID      string `storm:"index"`
	UABrowser   string
	UAOS        string
	UADevice    string
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
		UABrowser:   session.UserAgent().Browser(),
		UAOS:        session.UserAgent().OS(),
		UADevice:    session.UserAgent().Device(),
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
func (s Store) Get(key Keyable, limit int) (Sessionables, error) {
	var records []Record

	query := s.db.Select(q.And(getConditions(key)...))
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
		ua := Agent{
			browser: r.UABrowser,
			os:      r.UAOS,
			device:  r.UADevice,
		}
		sessions = append(sessions, *NewSession(
			r.SessionID,
			r.SessionType,
			r.UserID,
			ua,
			r.Permissions,
			r.OriginID,
			r.Origin,
		))
	}

	return sessions, nil
}

// Delete implements the Storable interface
func (s Store) Delete(key Keyable) error {
	sessions, err := s.Get(key, NoLimit)
	if err != nil {
		return err
	}
	for i := 0; i < sessions.Length(); i++ {
		s.db.DeleteStruct(Record{
			SessionID: sessions.Get(i).SessionID(),
		})
	}

	return nil
}

// Cleanup implements the Storable interface
func (s Store) Cleanup() error {
	var records []Record

	if err := s.db.Select(NewExpiredTimeMatcher("CreatedAt", "Expiry")).Find(&records); err != nil {
		return err
	}

	for _, r := range records {
		s.Delete(NewSession(r.SessionID, r.SessionType, r.UserID, nil, []string{}, r.OriginID, ""))
	}

	return nil
}

// getConditions returns a list of conditions for indexed fields
func getConditions(key Keyable) []q.Matcher {
	var (
		conditions []q.Matcher

		sessionID = key.SessionID()
		userID    = key.UserID()
		originID  = key.OriginID()
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
