ipregistry-go
----

[![License: MIT][401]][402] [![GoDoc][101]][102] [![Release][103]][104] [![Build Status][201]][202] [![Coveralls Coverage][203]][204] [![Codecov Coverage][205]][206]
[![Go Report Card][301]][302] [![Code Climate][303]][304] [![BCH compliance][305]][306] [![CodeFactor][307]][308] [![codebeat][309]][310] [![Scrutinizer Code Quality][311]][312] [![FOSSA Status][403]][404]


<!-- Basic -->

[101]: https://godoc.org/github.com/evalphobia/ipregistry-go?status.svg
[102]: https://godoc.org/github.com/evalphobia/ipregistry-go
[103]: https://img.shields.io/github/release/evalphobia/ipregistry-go.svg
[104]: https://github.com/evalphobia/ipregistry-go/releases/latest
[105]: https://img.shields.io/github/downloads/evalphobia/ipregistry-go/total.svg?maxAge=1800
[106]: https://github.com/evalphobia/ipregistry-go/releases
[107]: https://img.shields.io/github/stars/evalphobia/ipregistry-go.svg
[108]: https://github.com/evalphobia/ipregistry-go/stargazers


<!-- Testing -->

[201]: https://github.com/evalphobia/ipregistry-go/workflows/test/badge.svg
[202]: https://github.com/evalphobia/ipregistry-go/actions?query=workflow%3Atest
[203]: https://coveralls.io/repos/evalphobia/ipregistry-go/badge.svg?branch=master&service=github
[204]: https://coveralls.io/github/evalphobia/ipregistry-go?branch=master
[205]: https://codecov.io/gh/evalphobia/ipregistry-go/branch/master/graph/badge.svg
[206]: https://codecov.io/gh/evalphobia/ipregistry-go


<!-- Code Quality -->

[301]: https://goreportcard.com/badge/github.com/evalphobia/ipregistry-go
[302]: https://goreportcard.com/report/github.com/evalphobia/ipregistry-go
[303]: https://codeclimate.com/github/evalphobia/ipregistry-go/badges/gpa.svg
[304]: https://codeclimate.com/github/evalphobia/ipregistry-go
[305]: https://bettercodehub.com/edge/badge/evalphobia/ipregistry-go?branch=master
[306]: https://bettercodehub.com/
[307]: https://www.codefactor.io/repository/github/evalphobia/ipregistry-go/badge
[308]: https://www.codefactor.io/repository/github/evalphobia/ipregistry-go
[309]: https://codebeat.co/badges/142f5ca7-da37-474f-9264-f708ade08b5c
[310]: https://codebeat.co/projects/github-com-evalphobia-ipregistry-go-master
[311]: https://scrutinizer-ci.com/g/evalphobia/ipregistry-go/badges/quality-score.png?b=master
[312]: https://scrutinizer-ci.com/g/evalphobia/ipregistry-go/?branch=master

<!-- License -->
[401]: https://img.shields.io/badge/License-MIT-blue.svg
[402]: LICENSE.md
[403]: https://app.fossa.com/api/projects/git%2Bgithub.com%2Fevalphobia%2Fipregistry-go.svg?type=shield
[404]: https://app.fossa.com/projects/git%2Bgithub.com%2Fevalphobia%2Fipregistry-go?ref=badge_shield


Unofficial golang library for [Ipregistry](https://ipregistry.co/).


# Quick Usage for API

```go
package main

import (
	"fmt"

	"github.com/evalphobia/ipregistry-go/config"
	"github.com/evalphobia/ipregistry-go/ipregistry"
)

func main() {
	conf := config.Config{
        // you can set auth values to config directly, otherwise used from environment variables.
		APIKey:  "<your Ipregistry API Key>",
		Debug:      false,
	}

	svc, err := ipregistry.New(conf)
	if err != nil {
		panic(err)
	}

	// execute score API
	resp, err := svc.SingleIP("8.8.8.8")
	if err != nil {
		panic(err)
	}
	if resp.HasError() {
		panic(fmt.Errorf("code=[%s] message=[%s]", resp.ErrData.Code, resp.ErrData.Message))
	}

	// just print response in json format
	b, _ := json.Marshal(resp)
	fmt.Printf("%s", string(b))
}
```

see example dir for more examples, and see [official API document](https://ipregistry.co/docs/) for more details (especially request/response).


# Environment variables

| Name | Description |
|:--|:--|
| `IPREGISTRY_APIKEY` | [Ipregistry API Key](https://ipregistry.co/docs/authentication). |
