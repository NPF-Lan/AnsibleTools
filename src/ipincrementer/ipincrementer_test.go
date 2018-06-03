package main

import (
	"net"
	"testing"
)

func TestIp2int(t *testing.T) {
	tables := []struct {
		x net.IP
		y uint32
	}{
		{net.ParseIP("127.0.0.1"), 2130706433},
		{net.ParseIP("192.168.0.1"), 3232235521},
		{net.ParseIP("1.1.1.1"), 16843009},
		{net.ParseIP("255.255.255.255"), 4294967295},
	}

	for _, table := range tables {
		ipInt := ip2int(table.x)
		if ipInt != table.y {
			t.Errorf("ip2int of (%s) was incorrect, got: %d, want: %d.", table.x, ipInt, table.y)
		}
	}
}

func TestUint2Ip(t *testing.T) {
	tables := []struct {
		x net.IP
		y uint32
	}{
		{net.ParseIP("127.0.0.1"), 2130706433},
		{net.ParseIP("192.168.0.1"), 3232235521},
		{net.ParseIP("1.1.1.1"), 16843009},
		{net.ParseIP("255.255.255.255"), 4294967295},
	}

	for _, table := range tables {
		ip := int2ip(table.y)
		if ip.String() != table.x.String() {
			t.Errorf("int2ip of (%d) was incorrect, got: %s, want: %s.", table.y, ip, table.x)
		}
	}
}
