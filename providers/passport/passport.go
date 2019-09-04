package passport

import (
	"context"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/fatih/structs"
	authsdk "go.lsl.digital/lardwaz/sdk/auth"
	passportsdk "go.lsl.digital/passport/sdk/auth"
)

// Provider implements the Provider interface
type Provider struct {
	name       string
	logger     *log.Logger
	cookieName string
	endpoints  *authsdk.Endpoints
	options    *authsdk.Options
	Client     *passportsdk.Client
}

// New returns a new passport provider
func New(creds authsdk.Credentials, cookieName string, logger *log.Logger) *Provider {
	// Init logger
	if logger == nil {
		logger = log.New()
	}

	// Init insecure client
	client, err := passportsdk.NewInsecureClient(creds)
	if err != nil {
		logger.Errorf("Error creating new auth client: %v", err)
		return nil
	}

	return &Provider{
		name:       "passport",
		logger:     logger,
		cookieName: cookieName,
		Client:     client,
	}
}

// Name returns the name of provider
func (p Provider) Name() string {
	return p.name
}

// Options returns the provider options
func (p Provider) Options() *authsdk.Options {
	return p.options
}

// SetOptions sets the provider options
func (p *Provider) SetOptions(options *authsdk.Options) {
	p.options = options
}

// Endpoints returns the provider endpoints
func (p Provider) Endpoints() *authsdk.Endpoints {
	return p.endpoints
}

// SetEndpoints sets the provider endpoints
func (p *Provider) SetEndpoints(endpoints *authsdk.Endpoints) {
	p.endpoints = endpoints
}

// Session returns a Session
func (p Provider) Session(req *http.Request) (authsdk.Session, error) {
	cookie, err := req.Cookie(p.cookieName)
	if err != nil {
		p.logger.Errorf("%v: %v", authsdk.ErrTokenNotFound, err)
		return nil, authsdk.ErrTokenNotFound
	}

	token := cookie.Value

	gReq := &passportsdk.GetInfoRequest{
		Key: token,
	}

	gResp, err := p.Client.Wire.GetInfo(context.Background(), gReq)
	if err != nil {
		p.logger.Errorf("%v: %v", authsdk.ErrFetchUserData, err)
		return nil, authsdk.ErrFetchUserData
	}

	raw := structs.Map(gResp)

	user := authsdk.User{
		RawData:   raw,
		Token:     token,
		Provider:  p.name,
		UserID:    gResp.GetUserID(),
		Username:  gResp.GetUsercode(),
		Email:     gResp.GetEmail(),
		FirstName: gResp.GetFirstname(),
		LastName:  gResp.GetLastname(),
		Types:     gResp.GetRoles(),
		Actions:   gResp.GetPermissions(),
	}

	return NewSession(user), nil
}
