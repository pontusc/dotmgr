package config

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// Automatically validates any string given for Repository URLs
func (u *RepoURL) UnmarshalText(text []byte) error {
	s := string(text)
	var errs []error

	if !strings.HasPrefix(s, "https://") && !strings.HasPrefix(s, "git@") {
		return fmt.Errorf("invalid prefix, no 'git@' or 'https://' found")
	}

	// Validate SSH git URLs
	if strings.HasPrefix(s, "git@") {
		remainder := s[4:] // keep everything after 'git@'
		beforeSeparator, afterSeparator, separatorFound := strings.Cut(remainder, ":")

		if !separatorFound {
			errs = append(errs, fmt.Errorf("missing ':' separator"))
		} else {
			if beforeSeparator == "" {
				errs = append(errs, fmt.Errorf("missing host"))
			}
			if afterSeparator == "" {
				errs = append(errs, fmt.Errorf("missing path"))
			}
		}
		if !strings.HasSuffix(s, ".git") {
			errs = append(errs, fmt.Errorf("url must end in '.git'"))
		}
	} else {
		// Validate HTTPS URLs
		parsed, err := url.Parse(s)
		// Parser error
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		if parsed.Scheme != "https" {
			errs = append(errs, fmt.Errorf("scheme must be https, got %q", parsed.Scheme))
		}
		if parsed.Host == "" {
			errs = append(errs, fmt.Errorf("missing host"))
		}
		if !strings.HasSuffix(parsed.Path, ".git") {
			errs = append(errs, fmt.Errorf("url must end in '.git'"))
		}
	}

	if err := errors.Join(errs...); err != nil {
		return fmt.Errorf("invalid repository url %q: %w", s, err)
	}

	*u = RepoURL(s)
	return nil
}
