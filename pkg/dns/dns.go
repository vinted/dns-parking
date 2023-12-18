package dns

import (
	"fmt"
	"github.com/miekg/dns"
	"github.com/vinted/dns-parking/pkg/config"
	"log"
)

func parseQuery(m *dns.Msg) {
	cfg := *config.Config
	for _, q := range m.Question {
		switch q.Qtype {
		case dns.TypeSOA:
			log.Printf("SOA for %s\n", q.Name)
			rr, err := dns.NewRR(fmt.Sprintf("%s %s IN SOA %s. %s. %s %s %s %s %s", q.Name, cfg.SOARefresh, cfg.NS[0], cfg.SOARname, cfg.SOASerial, cfg.SOARefresh, cfg.SOARetry, cfg.SOAExpire, cfg.SOATTL))
			if err == nil {
				m.Answer = append(m.Answer, rr)
			} else {
				log.Printf("Error sending SOA record: %s\n", err)
			}
		case dns.TypeNS:
			log.Printf("NS for %s\n", q.Name)
			for _, ns := range config.Config.NS {
				rr, err := dns.NewRR(fmt.Sprintf("%s %s IN NS %s.", q.Name, config.Config.SOATTL, ns))
				if err == nil {
					m.Answer = append(m.Answer, rr)
				} else {
					log.Printf("Error sending NS record: %s\n", err)
				}
			}
		}
	}
}

func handleDnsRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		parseQuery(m)
	}
	w.WriteMsg(m)
}

func Start(listenAddress string) {
	dns.HandleFunc(".", handleDnsRequest)
	server := &dns.Server{Addr: listenAddress, Net: "udp"}
	log.Printf("Server started at %s\n", listenAddress)
	err := server.ListenAndServe()
	defer server.Shutdown()
	if err != nil {
		log.Fatalf("Failed to start server: %s\n ", err.Error())
	}
}
