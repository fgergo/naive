// Package naive is supposed to promote naive servers
// without real authentication or authorization.
// Background: https://crawshaw.io/blog/remembering-the-lan
package naive

import (
	"strconv"
	"strings"
)

// CgnatIp(ip) returns true if ip is a valid ip address
// in the CGNAT address space (100.64.0.0/10)
// otherwise returns false.
// Details in rfc6598.
func CgnatIp(ip string) bool {
	octets := strings.Split(ip, ".")
	if len(octets) < 4 {
		return false
	}
	for i, s := range octets {
		o, err := strconv.Atoi(s)
		if err != nil || o < 0 || o > 255 {
			return false
		}
		switch i {
		case 0:
			if o != 100 {
				return false
			}
		case 1:
			if o < 64 || o > 127 {
				return false
			}
		case 4:
			return false
		}
	}

	return true
}

// tba possibly:
// - user and possibly session management without authentication
// - service discovery on cgnat subnet
// - mechanism to check if binary is "running on" cgnat subnet
// - ...
