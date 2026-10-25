package auth

import (
	"fmt"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The two token credentials are validated against two different grammars on purpose.
//
// An access token becomes the credentials of an `Authorization: Bearer` header,
// so it has to satisfy RFC 6750 section 2.1 `b64token`:
//
//	b64token = 1*( ALPHA / DIGIT / "-" / "." / "_" / "~" / "+" / "/" ) *"="
//
// A refresh token is never a header. It is a field of a token request body,
// RFC 6749 section 5.1, whose grammar is `*VSCHAR` (visible ASCII, %x20-7E),
// which keeps punctuation Bearer has no room for, such as a comma, a colon or a
// quote. Holding a refresh token to the header grammar would refuse tokens this
// API is entitled to issue.
//
// They are validated rather than coerced. Go's type assertion already refuses a
// JSON number; these refuse every other string that is not a credential: one with
// a newline, which splits a header; one with a space, which ends the credential
// and starts a parameter; an empty one; and one with a control character.

// validateAccessToken checks an access token against RFC 6750 2.1 b64token.
func validateAccessToken(token string) error {
	if token == "" {
		return refuseCredential("access_token", "is empty", token)
	}
	padding := false
	for index := 0; index < len(token); index++ {
		character := token[index]
		if character == '=' {
			// Padding, and only ever at the end.
			padding = true
			continue
		}
		if padding {
			return refuseCredential("access_token",
				"carries a character after its '=' padding", token)
		}
		if !isB64TokenCharacter(character) {
			return refuseCredential("access_token", fmt.Sprintf(
				"carries the byte %#02x, which the RFC 6750 2.1 b64token grammar an "+
					"Authorization: Bearer header is built from does not allow", character), token)
		}
	}
	if padding && len(token) == countOf(token, '=') {
		return refuseCredential("access_token", "is nothing but '=' padding", token)
	}
	return nil
}

// validateRefreshToken checks a refresh token against RFC 6749 5.1 *VSCHAR.
//
// An empty refresh token is allowed and means "there is none": the grant need not
// issue one, and TokenRecord then carries "" rather than a value.
func validateRefreshToken(token string) error {
	for index := 0; index < len(token); index++ {
		character := token[index]
		if character < 0x20 || character > 0x7e {
			return refuseCredential("refresh_token", fmt.Sprintf(
				"carries the byte %#02x, which is outside the RFC 6749 5.1 visible-ASCII "+
					"range a token-request field is written in", character), token)
		}
	}
	return nil
}

// validateTokenCredentials checks both, and is what every entry point calls.
func validateTokenCredentials(accessToken, refreshToken string) error {
	if err := validateAccessToken(accessToken); err != nil {
		return err
	}
	return validateRefreshToken(refreshToken)
}

func isB64TokenCharacter(character byte) bool {
	switch {
	case character >= 'A' && character <= 'Z':
		return true
	case character >= 'a' && character <= 'z':
		return true
	case character >= '0' && character <= '9':
		return true
	}
	switch character {
	case '-', '.', '_', '~', '+', '/':
		return true
	}
	return false
}

func countOf(text string, wanted byte) int {
	total := 0
	for index := 0; index < len(text); index++ {
		if text[index] == wanted {
			total++
		}
	}
	return total
}

// refuseCredential names the problem without printing the credential; the length
// it gives tells a truncated value from a mistyped one and discloses nothing.
func refuseCredential(field, problem, token string) error {
	return &exerrors.SdkError{
		Kind: exerrors.ErrAuth,
		Message: fmt.Sprintf(
			"the %s %s, so it is not a credential this SDK will send. Its value is not "+
				"shown; it is %d bytes long", field, problem, len(token)),
		ErrorType: string(exhttp.ErrorTypeAuthentication),
	}
}
