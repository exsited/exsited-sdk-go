package auth

import (
	"context"
	"strconv"
	"time"

	"github.com/exsited/exsited-sdk-go/v4/config"
	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
	"github.com/exsited/exsited-sdk-go/v4/redact"
)

// TokenProvider obtains and refreshes the OAuth2 access token.
//
// Success means any 2xx: the token endpoint may answer 200 or 201. It issues a
// refresh token and honours the refresh_token grant, so a renewal tries that first
// and falls back to a full client_credentials request.
//
// A v4 client authenticates at /api/v4/oauth2/token.
type TokenProvider struct {
	config    *config.SdkConfig
	store     TokenStore
	transport *exhttp.HTTPTransport
	namespace string
}

func NewTokenProvider(
	config *config.SdkConfig, store TokenStore, transport *exhttp.HTTPTransport,
) *TokenProvider {
	return &TokenProvider{
		config:    config,
		store:     store,
		transport: transport,
		namespace: BuildCredentialsNamespace(
			config.Auth.BaseURL, config.Auth.ClientID,
			config.Auth.ClientSecret, config.Auth.GrantType),
	}
}

// AccessToken returns a usable token, fetching one only if needed.
//
// A store that cannot be read fails the call rather than being treated as empty;
// see auth.ErrTokenCache.
func (p *TokenProvider) AccessToken(ctx context.Context) (string, error) {
	record, ok, err := p.store.Read(p.namespace)
	if err != nil {
		return "", err
	}
	if ok && !record.NeedsRefresh() {
		return record.AccessToken, nil
	}
	return p.renew(ctx, "")
}

// Refresh replaces a token the API has just rejected.
//
// If the store already holds a different one, another worker refreshed first and
// that token is used rather than minting a third.
func (p *TokenProvider) Refresh(ctx context.Context, stale string) (string, error) {
	return p.renew(ctx, stale)
}

// InvalidateAccessToken evicts a token the API has just refused, if the cache
// still holds exactly that one. It reports whether an entry was removed.
//
// Without it, a refused token whose refresh failed (endpoint unreachable, a 4xx,
// a refused lock) stays in the shared file cache, and every later process sends
// it, gets 401 and fails the same way.
//
// It is a compare-and-clear rather than a bare clear: between the 401 and this
// call another worker may have written a live token, and on these credentials a
// new grant invalidates the previous token, so discarding a live entry starts a
// herd in which each grant kills the one before it. An empty dead token names no
// token, so nothing is evicted and this answers false.
//
// It holds the same locks renew uses, the store's in-process renewal mutex and its
// cross-process sidecar, so the read and the clear cannot straddle another
// process's write. A lock that cannot be taken fails closed: the entry is left alone.
func (p *TokenProvider) InvalidateAccessToken(dead string) (bool, error) {
	if dead == "" {
		return false, nil
	}
	removed := false
	err := p.store.WithLock(p.namespace, func() error {
		record, ok, err := p.store.Read(p.namespace)
		if err != nil {
			return err
		}
		if !ok || record.AccessToken != dead {
			// Already replaced by another worker, or never there: nothing to evict.
			return nil
		}
		if err := p.store.Clear(p.namespace); err != nil {
			return err
		}
		removed = true
		return nil
	})
	if err != nil {
		return false, err
	}
	return removed, nil
}

func (p *TokenProvider) renew(ctx context.Context, stale string) (string, error) {
	existing, ok, err := p.store.Read(p.namespace)
	if err != nil {
		return "", err
	}
	if ok && existing.AccessToken != stale && !existing.NeedsRefresh() {
		return existing.AccessToken, nil
	}

	var minted *TokenRecord
	if ok && existing.RefreshToken != "" {
		minted = p.requestWithRefreshGrant(ctx, existing.RefreshToken)
	}
	if minted == nil {
		fresh, err := p.requestWithCredentials(ctx)
		if err != nil {
			return "", err
		}
		minted = fresh
	}

	record := *minted
	// A store that cannot lock across processes refuses, and the refusal is the
	// call's: writing anyway could overwrite another process's entry in the shared
	// file. A caller who would rather have the token than the exclusion sets
	// InMemoryToken.
	if err := p.store.WithLock(p.namespace, func() error {
		return p.store.Write(p.namespace, record)
	}); err != nil {
		return "", err
	}
	return record.AccessToken, nil
}

func (p *TokenProvider) requestWithCredentials(ctx context.Context) (*TokenRecord, error) {
	// The four fields the token endpoint takes; it defines no auth_string.
	payload := map[string]any{
		"grant_type":    p.config.Auth.GrantType,
		"client_id":     p.config.Auth.ClientID,
		"client_secret": p.config.Auth.ClientSecret,
		"redirect_uri":  p.config.Auth.RedirectURI,
	}
	response, err := p.post(ctx, payload)
	if err != nil {
		return nil, err
	}
	// Refused credentials are answered with a redirect to the sign-in page rather
	// than a 401. The transport never follows it, so the 3xx itself arrives here.
	// The message names neither a credential nor the redirect's target.
	if response.StatusCode >= 300 && response.StatusCode < 400 {
		return nil, &exerrors.SdkError{
			Kind: exerrors.ErrAuth,
			Message: "the token request was refused: check client_id, client_secret and redirect_uri " +
				"(the API answered " + strconv.Itoa(response.StatusCode) + ", redirecting to its sign-in page)",
			StatusCode: response.StatusCode,
			ErrorType:  string(exhttp.ErrorTypeAuthentication),
		}
	}
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		detail := "no detail"
		if messages := response.ErrorMessages(); len(messages) > 0 {
			detail = messages[0]
		}
		return nil, &exerrors.SdkError{
			Kind:        exerrors.ErrAuth,
			Message:     "token request failed with " + strconv.Itoa(response.StatusCode) + ": " + detail,
			StatusCode:  response.StatusCode,
			ErrorType:   string(exhttp.ErrorTypeAuthentication),
			RawResponse: redact.RedactValue(response.Map),
		}
	}
	return buildTokenRecord(response.Map, "")
}

func (p *TokenProvider) requestWithRefreshGrant(ctx context.Context, refreshToken string) *TokenRecord {
	// No redirect_uri: the refresh grant does not take one.
	payload := map[string]any{
		"grant_type":    "refresh_token",
		"client_id":     p.config.Auth.ClientID,
		"client_secret": p.config.Auth.ClientSecret,
		"refresh_token": refreshToken,
	}
	response, err := p.post(ctx, payload)
	if err != nil || response.StatusCode < 200 || response.StatusCode >= 300 {
		// A failed refresh must fall back to a full grant, not propagate.
		return nil
	}
	record, err := buildTokenRecord(response.Map, refreshToken)
	if err != nil {
		return nil
	}
	return record
}

// post drives the transport directly: the client imports this package, so calling
// it here would be an import cycle. The grant is unauthenticated, takes no path
// parameters and names no typed view, so nothing the client adds is needed.
func (p *TokenProvider) post(ctx context.Context, payload map[string]any) (*exhttp.ApiResponse, error) {
	absolute, err := exhttp.BuildURL(p.config.Auth.BaseURL, p.config.APIVersion.TokenPath())
	if err != nil {
		return nil, err
	}
	exchange, err := p.transport.SendOnce(
		ctx, "POST", absolute, nil, payload, nil,
		map[string]string{"Content-Type": "application/json"}, "")
	if err != nil {
		return nil, err
	}
	return exhttp.NewApiResponseFromExchange(exchange, nil, ""), nil
}

func buildTokenRecord(body any, previousRefreshToken string) (*TokenRecord, error) {
	record, ok := body.(map[string]any)
	if !ok {
		return nil, &exerrors.SdkError{
			Kind: exerrors.ErrAuth, Message: "token response was not a JSON object",
			ErrorType: string(exhttp.ErrorTypeAuthentication),
		}
	}
	accessToken, ok := record["access_token"].(string)
	if !ok || accessToken == "" {
		return nil, &exerrors.SdkError{
			Kind: exerrors.ErrAuth, Message: "token response carried no access_token",
			ErrorType: string(exhttp.ErrorTypeAuthentication), RawResponse: redact.RedactValue(body),
		}
	}
	// Validated, never coerced: the assertion above refuses a JSON number, and this
	// refuses every other string that is not a credential. See token_grammar.go.
	if err := validateAccessToken(accessToken); err != nil {
		return nil, err
	}
	// Read through NumberAsFloat64: the transport decodes with UseNumber, so
	// expires_in arrives as a json.Number, which a float64 assertion would miss.
	ttl := float64(defaultTokenTTLSeconds)
	if given, ok := exhttp.NumberAsFloat64(record["expires_in"]); ok {
		ttl = given
	}
	if ttl < 0 {
		ttl = 0
	}
	if ttl > maxTokenTTLSeconds {
		ttl = maxTokenTTLSeconds
	}
	refreshToken := previousRefreshToken
	if given, ok := record["refresh_token"].(string); ok && given != "" {
		refreshToken = given
	}
	// A different grammar from the access token's: a refresh token is a
	// token-request field rather than a header, so it keeps the punctuation Bearer
	// has no room for.
	if err := validateRefreshToken(refreshToken); err != nil {
		return nil, err
	}
	return &TokenRecord{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(time.Duration(ttl) * time.Second).UnixMilli(),
	}, nil
}
