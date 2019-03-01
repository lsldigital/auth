package auth

import (
	"errors"
	"time"
)

var (
	ErrAlreadyRunning = errors.New("scheduler already running")
	ErrNotRunning     = errors.New("scheduler not running")
	ErrStop           = errors.New("scheduler did not stop")
	ErrTimeout        = errors.New("timed out")
)

type Schedulable interface {
	Start(interval time.Duration) error
	Stop() error
	IsRunning() bool
}

// Scheduler implements the Schedulable interface
type Scheduler struct {
	store    Storable
	running  bool
	stopSig  chan struct{}
	stopResp chan bool
}

// NewScheduler returns a new Scheduler
func NewScheduler(store Storable) *Scheduler {
	scheduler := &Scheduler{store: store}
	scheduler.stopSig = make(chan struct{})
	scheduler.stopResp = make(chan bool)

	return scheduler
}

// Start implements the Schedulable interface
func (s Scheduler) Start(interval time.Duration) error {
	if s.running {
		return ErrAlreadyRunning
	}
	task := func() {
		// TODO: implement cleanup functionality
	}
	go func() {
		s.running = true
		task()

		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				task()
			case <-s.stopSig:
				s.running = false
				s.stopResp <- true
				return
			}
		}
	}()
	return nil
}

// Stop implements the Schedulable interface
func (s Scheduler) Stop() error {
	if !s.running {
		return ErrNotRunning
	}
	s.stopSig <- struct{}{}

	select {
	case res := <-s.stopResp:
		if !res {
			return ErrStop
		}
		return nil
	case <-time.After(3 * time.Second):
		return ErrTimeout
	}
}

// IsRunning implements the Schedulable interface
func (s Scheduler) IsRunning() bool {
	return s.running
}
