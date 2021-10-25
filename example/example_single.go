package main

import (
	"flag"
	"fmt"

	"github.com/evalphobia/ipregistry-go/config"
	"github.com/evalphobia/ipregistry-go/ipregistry"
)

// nolint
func main() {
	var ipaddr string
	flag.StringVar(&ipaddr, "ipaddr", "", "set target ip address")
	flag.Parse()

	conf := config.Config{
		APIKey: "",
		Debug:  true,
	}

	svc, err := ipregistry.New(conf)
	if err != nil {
		panic(err)
	}

	resp, err := svc.SingleIP(ipaddr)
	fmt.Printf("[%+v]\n", resp)
	if err != nil {
		panic(err)
	}
}
