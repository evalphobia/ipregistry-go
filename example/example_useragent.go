package main

import (
	"flag"
	"fmt"

	"github.com/evalphobia/ipregistry-go/config"
	"github.com/evalphobia/ipregistry-go/ipregistry"
)

// nolint
func main() {
	var ua string
	flag.StringVar(&ua, "useragent", "", "set target useragent")
	flag.Parse()

	conf := config.Config{
		APIKey: "",
		Debug:  true,
	}

	svc, err := ipregistry.New(conf)
	if err != nil {
		panic(err)
	}

	resp, err := svc.UserAgent([]string{ua})
	fmt.Printf("[%+v]\n", resp)
	if err != nil {
		panic(err)
	}
}
