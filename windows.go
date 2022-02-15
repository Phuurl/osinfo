//go:build windows
// +build windows

package osinfo

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

// osHandler gathers OS information and returns a populated osinfo.Details struct.
// Depending on the error encountered, if err != nil, the Details struct may still be partially populated.
// For Windows, the following information is populated:
//	- Name: Full OS name (eg `Windows 10 Pro`)
//	- Version: OS major version (eg `10`)
//	- Build: OS build (eg `21H1`)
func osHandler() (Details, error) {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return Details{}, err
	}
	defer key.Close()

	name, _, err := key.GetStringValue("ProductName")
	if err != nil {
		return Details{}, err
	}
	ver, _, err := key.GetIntegerValue("CurrentMajorVersionNumber")
	if err != nil {
		return Details{}, err
	}
	build, _, err := key.GetStringValue("DisplayVersion")
	if err != nil {
		return Details{}, err
	}

	return Details{
		Name:    name,
		Version: fmt.Sprintf("%d", ver),
		Build:   build,
	}, err
}
