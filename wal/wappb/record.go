package walpb

import "errors"

var (
	// ErrCRCMismatch is error for crc mismatch
	ErrCRCMismatch = errors.New("walpb: crc mismatch")
)
