package btcnet

import "net"

var dnsSeeds = []struct {
	Name     string
	Hostname string
}{
	{Name: "bitcoin.sipa.be", Hostname: "seed.bitcoin.sipa.be"},
	{Name: "bluematt.me", Hostname: "dnsseed.bluematt.me"},
	{Name: "dashjr.org", Hostname: "dnsseed.bitcoin.dashjr.org"},
	{Name: "xf2.org", Hostname: "bitseed.xf2.org"},
}

// DiscoverDNS finds bitcoin nodes using DNS discovery
func DiscoverDNS() []net.IP {
	addrs := make([]net.IP, 0)

	for _, s := range dnsSeeds {
		a, err := net.LookupIP(s.Hostname)
		if err != nil {
			continue
		}
		addrs = append(addrs, a...)
	}

	return addrs
}
