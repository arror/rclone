//go:build !windows
// +build !windows

package buildinfo

// GetOSVersion returns OS version, kernel and bitness
func GetOSVersion() (osVersion, osKernel string) {
	return
}
