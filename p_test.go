package naive

import (
	"testing"
)

// a cgnat ip address is a valid ip address
// between 100.64.0.0 and 100.127.255.255
func TestCgnatIp(t *testing.T) {
	var tests = []struct {
		ip    string
		needs bool
	}{
		{"100.64.0.0", true},
		{"100.127.255.255", true},
		{"100.96.1.2", true},
		{"127.1", false},
		{"1.12.123.1234", false},
		{"100.64.0. 0", false},	// no whitespace between octets
		{" 100.64.0.0", false}, 	// no beginning whitespace
		{"100.64.0.0 ", false},	// no trailing whitespace 
		{"100.255.255.1234", false},	// string too long (and invalid octet)
		{"256.257.258.259", false},
		{"127.0.0.1", false},
		{"10.0.0.1", false},
		{"1.1.1.1", false},
		{"1.-42.0.0", false},
		{"100.64.0.256", false},
		{"100.63.0.0", false},
		{"100.96.0.355", false},
		{"100.96.invalid.0", false},
		{"100.64", false},
		{"100.64..", false},
		{"100.64..0", false},
		{"100.64.1.2.3", false},
	}
	for _, v := range tests {
		got := CgnatIp(v.ip)
		if v.needs != got {
			t.Errorf("isCgnat(%v), needs: %v, got: %v", v.ip, v.needs, got)
		}
	}
}
