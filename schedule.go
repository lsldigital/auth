package auth

import (
	"time"

	"github.com/asdine/storm"
)

type Schedulable interface {
	Start(interval time.Duration) error
	Stop() error
	IsRunning() bool
}

// Scheduler implements the Schedulable interface
type Scheduler struct {
	db      *storm.DB
	running bool
}

// NewScheduler returns a new Scheduler
func NewScheduler(db *storm.DB) *Scheduler {
	return &Scheduler{db: db}
}

// Start implements the Schedulable interface
func (s Scheduler) Start(interval time.Duration) error {
	return nil
}

// Stop implements the Schedulable interface
func (s Scheduler) Stop() error {
	return nil
}

// IsRunning implements the Schedulable interface
func (s Scheduler) IsRunning() bool {
	return s.running
}
