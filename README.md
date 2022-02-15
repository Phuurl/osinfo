# osinfo
Go library to detect OS information of the platform it is running on.

## Install
`go get github.com/Phuurl/osinfo`

## Supported operating systems
`osinfo` will return information about the following supported OSs:
 - Linux (kernel info for all, distro information required an `/etc/os-release` file)
 - macOS
 - Windows

## Example
`osinfo.GetOsInfo()` returns an `osinfo.Info` struct, which looks like this when run on a Linux system (converted to JSON here for ease):

```json
{
        "GOOS": "linux",
        "GOARCH": "arm64",
        "Details": {
                "Name": "Ubuntu",
                "Version": "20.04",
                "Build": "5.11.0-1028-aws"
        }
}
```

For more information on usage and the different returns for the different supported OSs, please see the [GoDocs](https://pkg.go.dev/github.com/Phuurl/osinfo).