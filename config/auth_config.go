package config

import (
	"net/url"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	"github.com/exsited/exsited-sdk-go/v4/redact"
)

// AuthConfig is who you are and which instance you are talking to.
//
// Four credentials, exactly the four the token endpoint takes.
type AuthConfig struct {
	// ClientName labels this instance. Shown in logs and used to select it.
	ClientName string
	// BaseURL is the instance URL, e.g. https://dev-api.exsited.com.
	BaseURL  string
	ClientID string
	// ClientSecret is masked by String().
	ClientSecret string
	// RedirectURI is sent in the client_credentials token body, and omitted from
	// the refresh_token one, which does not take it.
	RedirectURI string
	// GrantType defaults to client_credentials.
	GrantType string
	// TokenCachePath caches the token to that file. Empty means the platform
	// default; set InMemoryToken to keep it in memory instead.
	TokenCachePath string
	// InMemoryToken keeps the token for the life of the process.
	InMemoryToken bool
}

// Validate reports every problem at once rather than the first, so one round of
// fixing clears them all.
func (a *AuthConfig) Validate() error {
	var problems []string
	require := func(name, value string) {
		if strings.TrimSpace(value) == "" {
			problems = append(problems, name+": must not be blank")
		}
	}
	require("ClientName", a.ClientName)
	require("BaseURL", a.BaseURL)
	require("ClientID", a.ClientID)
	require("ClientSecret", a.ClientSecret)
	require("RedirectURI", a.RedirectURI)

	if strings.TrimSpace(a.BaseURL) != "" {
		trimmed := strings.TrimRight(strings.TrimSpace(a.BaseURL), "/")
		parsed, err := url.Parse(trimmed)
		switch {
		case err != nil:
			problems = append(problems, "BaseURL: not a URL: "+trimmed)
		case parsed.Scheme != "http" && parsed.Scheme != "https":
			problems = append(problems, "BaseURL: needs an http:// or https:// scheme")
		case parsed.User != nil:
			// An HTTP client turns userinfo into a Basic header that replaces the
			// bearer token.
			problems = append(problems, "BaseURL: must not carry user:password@ credentials")
		case parsed.RawQuery != "" || parsed.Fragment != "":
			problems = append(problems, "BaseURL: must not carry a query string or fragment")
		case portProblem(parsed) != "":
			problems = append(problems, "BaseURL: "+portProblem(parsed))
		default:
			a.BaseURL = trimmed
		}
	}
	if a.GrantType == "" {
		a.GrantType = "client_credentials"
	}
	if len(problems) > 0 {
		name := a.ClientName
		if strings.TrimSpace(name) == "" {
			name = "<unnamed>"
		}
		return exerrors.ConfigErrorf("%s", joinProblems(name, problems))
	}
	return nil
}

// String masks the secret, so an instance is safe to log.
func (a AuthConfig) String() string {
	return "AuthConfig(ClientName=" + a.ClientName +
		", BaseURL=" + a.BaseURL +
		", ClientID=" + a.ClientID +
		", ClientSecret=" + redact.Redacted +
		", RedirectURI=" + a.RedirectURI + ")"
}
