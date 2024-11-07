package sDomain

type DomainBindInfoOut struct {
	BindDomain string
	BindIp     string
}

func (s *sDomain) DomainBindInfo() (out DomainBindInfoOut, err error) {
	out.BindIp = "127.0.0.1"
	out.BindDomain = "name.shopkone.top"
	return out, err
}
