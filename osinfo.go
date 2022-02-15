package osinfo

import (
	"errors"
	"runtime"
)

// Details contains the detected OS name and version information.
// macOS:
//	- Name: OS name (eg `macOS` or `OSX`)
//	- Version: OS version (eg `10.16`)
//	- Build: OS build (eg `21C52`)
// Linux:
//	- Name: Distro name (eg `Ubuntu`)
//	- Version: Distro release version (eg `20.04`)
//	- Build: Linux kernel release (eg `5.11.0-1028-aws`)
// Windows:
//	- Name: Full OS name (eg `Windows 10 Pro`)
//	- Version: OS major version (eg `10`)
//	- Build: OS build (eg `21H1`)
type Details struct {
	Name    string
	Version string
	Build   string
}

// Info contains standard GOOS and GOARCH info, and the more detailed Details struct
type Info struct {
	GOOS    string
	GOARCH  string
	Details Details
}

// GetOsInfo returns a populated Info struct with OS information, and an error object.
// Depending on the error encountered, if err != nil, the Info struct may still be partially populated.
func GetOsInfo() (Info, error) {
	osInfo := Info{
		GOOS:   runtime.GOOS,
		GOARCH: runtime.GOARCH,
	}
	var err error = nil
	switch osInfo.GOOS {
	case "linux", "darwin", "windows":
		osInfo.Details, err = osHandler()
	default:
		return osInfo, errors.New("unsupported GOOS")
	}

	return osInfo, err
}
