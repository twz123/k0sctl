package version

import (
	"errors"
	"strings"
)

var (
	// Version of the product, is set during the build
	Version = "0.0.0"
	// GitCommit is set during the build
	GitCommit = "HEAD"
	// Environment of the product, is set during the build
	Environment = "development"
)

// IsPre is true when the current version is a prerelease
func IsPre() bool {
	return strings.Contains(Version, "-")
}

// Malformed indicates that the k0s version is invalid.
func Malformed(err error) error {
	if err == nil {
		return nil
	}

	return &malformed{err}
}

// IsMalformed indicates if the given error indicates that the k0s version is
// invalid.
func IsMalformed(err error) bool {
	var checked *malformed
	return errors.As(err, &checked)
}

type malformed struct{ error }

func (m *malformed) Unwrap() error {
	return m.error
}
