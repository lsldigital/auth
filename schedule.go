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
	DB      *storm.DB
	Running bool
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
	return s.Running
}
