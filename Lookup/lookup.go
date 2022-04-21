package Lookup

import "net"

func LookupDnsRecords(host string) ([]string, error) {
	records := make([]string, 0)
	ips, err := net.LookupHost(host)
	if err != nil {
		return nil, err
	}
	nss, err := net.LookupNS(host)
	if err == nil {
		for _, ns := range nss {
			records = append(records, ns.Host)
		}
	}
	mxs, err := net.LookupMX(host)
	if err == nil {
		for _, mx := range mxs {
			records = append(records, mx.Host)
		}
	}
	records = append(records, ips...)

	return records, nil
}
