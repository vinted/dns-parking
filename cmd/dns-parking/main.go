package main

import (
	"flag"
	"github.com/vinted/dns-parking/pkg/config"
	"github.com/vinted/dns-parking/pkg/dns"
	"log"
)

func main() {
	var (
		configFile    string
		listenAddress string
	)
	flag.StringVar(&configFile, "configFile", "/etc/dns-parking/config.json", "Path to config file.")
	flag.StringVar(&listenAddress, "listenAddress", "0.0.0.0:53", "Address to bind to.")
	flag.Parse()
	err := config.Init(configFile)
	if err != nil {
		log.Fatalf("Failed to read configuration. %v", err)
	}
	dns.Start(listenAddress)
}
