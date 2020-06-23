package domain

import "errors"

var (
	ErrInvalidVCS       = errors.New("invalid vcs type")
	ErrProjectNotFound  = errors.New("project not found")
	ErrPermissionDenied = errors.New("permission denied")
	ErrTooManyRequests  = errors.New("too many requests")
)
