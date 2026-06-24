package validation

import (
	"errors"
	"net/url"
	"strings"
)

const maxURLLength = 2048

var (
	ErrURLEmpty       = errors.New("URL is required")
	ErrURLTooLong     = errors.New("URL exceeds maximum length")
	ErrURLInvalid     = errors.New("URL is not valid")
	ErrURLScheme      = errors.New("URL must use http or https")
	ErrURLNoHost      = errors.New("URL must have a valid host")
	ErrURLHasFragment = errors.New("URL must not contain a fragment (#)")
	ErrURLHasUserInfo = errors.New("URL must not contain credentials")
)

// ValidateURL checks that the input is a safe, well-formed HTTP/HTTPS URL.
// Rejects: empty, too long, non-http(s) schemes, embedded credentials,
// fragments, and anything that cannot be parsed as a URL.
func ValidateURL(raw string) error {
	if strings.TrimSpace(raw) == "" {
		return ErrURLEmpty
	}

	if len(raw) > maxURLLength {
		return ErrURLTooLong
	}

	parsed, err := url.ParseRequestURI(raw)
	if err != nil {
		return ErrURLInvalid
	}

	scheme := strings.ToLower(parsed.Scheme)
	if scheme != "http" && scheme != "https" {
		return ErrURLScheme
	}

	if parsed.Host == "" {
		return ErrURLNoHost
	}

	// Reject embedded credentials (http://user:pass@host) — SSRF / info-leak vector
	if parsed.User != nil {
		return ErrURLHasUserInfo
	}

	// Fragments are client-side only; storing them is misleading and can mask payloads
	if parsed.Fragment != "" {
		return ErrURLHasFragment
	}

	return nil
}
