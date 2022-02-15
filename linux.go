//go:build linux
// +build linux

package osinfo

import (
	"errors"
	"os"
	"regexp"
	"strings"

	"golang.org/x/sys/unix"
)

// osHandler gathers OS information and returns a populated osinfo.Details struct.
// Depending on the error encountered, if err != nil, the Details struct may still be partially populated.
// For Linux, the following information is populated:
//	- Name: Distro name (eg `Ubuntu`)
//	- Version: Distro release version (eg `20.04`)
//	- Build: Linux kernel release (eg `5.11.0-1028-aws`)
func osHandler() (Details, error) {
	// Get kernel version from `uname`
	uname := unix.Utsname{}
	err := unix.Uname(&uname)
	if err != nil {
		return Details{}, err
	}
	kernelRelease := string(uname.Release[:])
	kernelRelease = strings.Trim(kernelRelease, "\u0000") // Trim trailing null bytes

	// Load os-release file for distro info parsing
	osRelease, err := os.ReadFile("/etc/os-release")
	if errors.Is(err, os.ErrNotExist) {
		return Details{
			Build: kernelRelease,
		}, errors.New("no /etc/os-release file, unable to determine distro")
	} else if err != nil {
		return Details{
			Build: kernelRelease,
		}, err
	}

	// Gather distro name
	distroNameRegex, err := regexp.Compile("(?m)^NAME=\"(.+)\"$")
	if err != nil {
		return Details{
			Build: kernelRelease,
		}, err
	}
	distroName := distroNameRegex.FindStringSubmatch(string(osRelease))
	if len(distroName) != 2 {
		return Details{
			Build: kernelRelease,
		}, errors.New("failed to match NAME /etc/os-release regex")
	}

	// Gather distro version
	distroVersionRegex, err := regexp.Compile("(?m)^VERSION_ID=\"?([^\"]+)\"?$")
	if err != nil {
		return Details{
			Build: kernelRelease,
		}, err
	}
	distroVersion := distroVersionRegex.FindStringSubmatch(string(osRelease))
	if len(distroName) != 2 {
		return Details{
			Name:  distroName[1],
			Build: kernelRelease,
		}, errors.New("failed to match VERSION_ID /etc/os-release regex")
	}

	return Details{
		Name:    distroName[1],
		Version: distroVersion[1],
		Build:   kernelRelease,
	}, nil
}
