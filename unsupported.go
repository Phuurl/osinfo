//go:build !linux && !darwin && !windows
// +build !linux,!darwin,!windows

package osinfo

import "errors"

func osHandler() (Details, error) {
	return Details{}, errors.New("unsupported GOOS")
}
