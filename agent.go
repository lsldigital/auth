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

// Browser implements the Agentable interface
func (a Agent) Browser() string {
	return a.browser
}

// OS implements the Agentable interface
func (a Agent) OS() string {
	return a.os
}

// Device implements the Agentable interface
func (a Agent) Device() string {
	return a.device
}
