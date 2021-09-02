package provider

import (
	"net"
	"testing"
)

func TestFakeInit(t *testing.T) {
	var f Fake
	m := make(map[string]string)
	err := f.Init(m)

	if err != nil {
		t.Errorf("got %q, wanted nil", err)
	}
}

func TestFakeGetARecord(t *testing.T) {
	var f Fake
	m := make(map[string]string)
	f.Init(m)

	// Test unset record
	ip, err := f.GetARecord("NXDOMAIN")
	if ip != nil && err != nil {
		t.Errorf("got %q and %q, wanted nil nil", ip, err)
	}

	// Set record and test it
	ip = net.IPv4(10, 0, 0, 1)
	fqdn := "example.tld"
	f.SetARecord(fqdn, ip)

	got, err := f.GetARecord(fqdn)
	if got.String() != ip.String() || err != nil {
		t.Errorf("got %q and err = %q, wanted %q and nil", got, err, ip)
	}
}
