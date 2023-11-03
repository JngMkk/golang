package dnspkg

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
)

// DNS 정보 저장
type Lookup struct {
	cname string
	hosts []string
}

// lookup 객체 출력
func (d *Lookup) String() string {
	result := ""

	for _, host := range d.hosts {
		result += fmt.Sprintf("%s IN A %s\n", d.cname, host)
	}

	return result
}

// 입력된 주소에 대한 cname과 host로 구성된 DNS Lookup을 반환
func LookupAddr(addr string) (*Lookup, error) {
	cname, err := net.LookupCNAME(addr)
	if err != nil {
		return nil, errors.Wrap(err, "error looking up CNAME")
	}

	hosts, err := net.LookupHost(addr)
	if err != nil {
		return nil, errors.Wrap(err, "error looking up HOST")
	}

	return &Lookup{cname: cname, hosts: hosts}, nil
}
