package dns_like_dig

import (
	"fmt"
	"github.com/pkg/errors"
	"net"
)

// Lookup Mapping url -> host IPs
type Lookup struct {
	cname string
	hosts []string
}

func (d *Lookup) String() string {
	result := ""
	for _, host := range d.hosts {
		result += fmt.Sprintf("%s IN A %s\n", d.cname, host)
	}
	return result
}

func LookupAddress(address string) (*Lookup, error) {
	cname, err := net.LookupCNAME(address)
	if err != nil {
		return nil, errors.Wrap(err, "error looking up CNAME")
	}

	hosts, err := net.LookupHost(address)
	if err != nil {
		return nil, errors.Wrap(err, "error looking up HOST")
	}

	return &Lookup{cname: cname, hosts: hosts}, nil
}
