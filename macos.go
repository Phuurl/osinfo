//go:build darwin
// +build darwin

package osinfo

import (
	"os"

	"howett.net/plist"
)

// Used for deserialising the macOS version plist file
type macOsReleasePlist struct {
	ProductBuildVersion       string `plist:"ProductBuildVersion"`
	ProductName               string `plist:"ProductName"`
	ProductUserVisibleVersion string `plist:"ProductUserVisibleVersion"`
}

// osHandler gathers OS information and returns a populated osinfo.Details struct.
// Depending on the error encountered, if err != nil, the Details struct may still be partially populated.
// For macOS, the following information is populated:
//	- Name: OS name (eg `macOS` or `OSX`)
//	- Version: OS version (eg `10.16`)
//	- Build: OS build (eg `21C52`)
func osHandler() (Details, error) {
	versionFile, err := os.Open("/System/Library/CoreServices/SystemVersion.plist")
	if err != nil {
		return Details{}, err
	}
	defer func(versionFile *os.File) {
		_ = versionFile.Close()
	}(versionFile)

	var plistResult macOsReleasePlist
	decoder := plist.NewDecoder(versionFile)
	err = decoder.Decode(&plistResult)
	if err != nil {
		return Details{}, err
	}

	return Details{
		Name:    plistResult.ProductName,
		Version: plistResult.ProductUserVisibleVersion,
		Build:   plistResult.ProductBuildVersion,
	}, err
}
