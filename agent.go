package auth

import (
	"fmt"
	"strings"

	"github.com/avct/uasurfer"
)

type Agentable interface {
	Browser() string
	OS() string
	Device() string
}

// Agent implements the Agentable interface
type Agent struct {
	browser string
	os      string
	device  string
}

// NewAgent returns a new Agent
func NewAgent(userAgent string) Agent {
	ua := uasurfer.Parse(userAgent)

	return Agent{
		browser: strings.Replace(fmt.Sprintf("%v - %v", ua.Browser.Name, ua.Browser.Version.Major), "Browser", "", -1),
		os:      strings.Replace(fmt.Sprintf("%v - %v", ua.OS.Name, ua.OS.Version.Major), "OS", "", -1),
		device:  strings.Replace(fmt.Sprintf("%v", ua.DeviceType), "Device", "", -1),
	}
}

// GetBrowser implements the Agentable interface
func (a Agent) GetBrowser() string {
	return a.browser
}

// GetOS implements the Agentable interface
func (a Agent) GetOS() string {
	return a.os
}

// GetDevice implements the Agentable interface
func (a Agent) GetDevice() string {
	return a.device
}
