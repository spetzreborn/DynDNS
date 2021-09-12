package client

import (
	"testing"
)

func TestIPInit(t *testing.T) {

	type TestIP struct {
		key         string
		value       string
		expectedErr bool
	}

	tests := []TestIP{
		{
			key:         "ip",
			value:       "127.0.0.1",
			expectedErr: false,
		},
		{
			key:         "Wrong key",
			value:       "127.0.0.1",
			expectedErr: true,
		},
		{
			key:         "ip",
			value:       "127", // Not an IP,
			expectedErr: true,
		},
		{
			key:         "ip",
			value:       "2001:db8::1", // Not an IPv4
			expectedErr: true,
		},
	}

	for _, test := range tests {
		param := make(map[string]string)
		param[test.key] = test.value

		var i IP
		err := i.Init(param)

		if !test.expectedErr && err != nil {
			t.Errorf("got %q, wanted error %t", err, test.expectedErr)
		}

		if test.expectedErr && err == nil {
			t.Errorf("got %q, wanted error %t", err, test.expectedErr)
		}
	}
}

func TestIPGetIPv4(t *testing.T) {
	var i IP
	m := make(map[string]string)
	m["ip"] = "127.0.0.1"
	i.Init(m)

	ip := i.GetIPv4()
	if ip.String() != m["ip"] {
		t.Errorf("got %q, wanted %q", ip.String(), m["ip"])
	}
}
