package btcnet

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

var ipre = regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)

// ExternalAddress attempts to discover the external (non-NAT) IP address of the host
// Derived from https://github.com/bitcoin/bitcoin/blob/master/src/net.cpp
func ExternalAddress() string {
	type IpHost struct {
		Address  string
		Hostname string
		Port     int
		Path     string
	}

	hosts := []IpHost{
		IpHost{Address: "91.198.22.70", Hostname: "checkip.dyndns.org", Port: 80, Path: ""},
		IpHost{Address: "74.208.43.192", Hostname: "www.showmyip.com", Port: 80, Path: "/simple/"},
	}

	for _, h := range hosts {

		u := url.URL{
			Scheme: "http",
			Host:   fmt.Sprintf("%s:%d", h.Address, h.Port),
			Path:   h.Path,
		}

		resp, err := http.Get(u.String())
		if err != nil {
			continue
		}
		if resp.StatusCode != 200 {
			resp.Body.Close()
			continue
		}

		buf := make([]byte, 1024)
		bufReader := bufio.NewReader(resp.Body)
		n, _ := bufReader.Read(buf)
		resp.Body.Close()
		if err != nil || n == 0 {
			continue
		}

		ip := ipre.FindString(string(buf))

		if ip != "" {
			return ip
		}

	}

	return ""
}
