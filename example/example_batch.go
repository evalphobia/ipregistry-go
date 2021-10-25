package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/evalphobia/ipregistry-go/config"
	"github.com/evalphobia/ipregistry-go/ipregistry"
)

// nolint
func main() {
	var ipaddr string
	flag.StringVar(&ipaddr, "ipaddr", "", "set target ip address list by space separated format")
	flag.Parse()

	conf := config.Config{
		APIKey: "",
		Debug:  true,
	}

	svc, err := ipregistry.New(conf)
	if err != nil {
		panic(err)
	}

	resp, err := svc.BatchIP(strings.Split(ipaddr, " "))
	fmt.Printf("[%+v]\n", resp)
	if err != nil {
		panic(err)
	}
}
