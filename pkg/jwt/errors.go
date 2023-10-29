package auth

import "errors"

var (
	ErrAPISecretRequired      = errors.New("api secret required")
	ErrTokenLifecycleRequired = errors.New("token lifecycle required")
)
