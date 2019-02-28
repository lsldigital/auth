package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/avct/uasurfer"
)

type Sessionable interface {
	GetID() string
	GetUserID() string
	GetOriginID() string
	GetCreatedAt() time.Time
	GetExpiry() time.Duration // Used by scheduler
}

type Agent struct {
	Browser string
	OS      string
	Device  string
}

// NewAgent returns a new Agent
func NewAgent(userAgent string) Agent {
	ua := uasurfer.Parse(userAgent)

	return Agent{
		Browser: strings.Replace(fmt.Sprintf("%v - %v", ua.Browser.Name, ua.Browser.Version.Major), "Browser", "", -1),
		OS:      strings.Replace(fmt.Sprintf("%v - %v", ua.OS.Name, ua.OS.Version.Major), "OS", "", -1),
		Device:  strings.Replace(fmt.Sprintf("%v", ua.DeviceType), "Device", "", -1),
	}
}

// Session implements the Sessionable interface
type Session struct {
	ID          string
	Type        int
	UserID      string
	UserAgent   Agent
	OriginID    string
	Origin      string
	Permissions []string
	Expiry      time.Duration
	CreatedAt   time.Time
}

// GetID implements the Sessionable interface
func (s Session) GetID() string {
	return s.ID
}

// GetUserID implements the Sessionable interface
func (s Session) GetUserID() string {
	return s.UserID
}

// GetOriginID implements the Sessionable interface
func (s Session) GetOriginID() string {
	return s.OriginID
}

// GetCreatedAt implements the Sessionable interface
func (s Session) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// GetExpiry implements the Sessionable interface
func (s Session) GetExpiry() time.Duration {
	return s.Expiry
}
