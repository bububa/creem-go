# creem-go creem.io SDK implementation in Golang

[![Go Reference](https://pkg.go.dev/badge/github.com/bububa/creem-go.svg)](https://pkg.go.dev/github.com/bububa/creem-go)
[![Go](https://github.com/bububa/creem-go/actions/workflows/go.yml/badge.svg)](https://github.com/bububa/creem-go/actions/workflows/go.yml)
[![goreleaser](https://github.com/bububa/creem-go/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/bububa/creem-go/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/creem-go.svg)](https://github.com/bububa/creem-go)
[![GoReportCard](https://goreportcard.com/badge/github.com/bububa/creem-go)](https://goreportcard.com/report/github.com/bububa/creem-go)
[![GitHub license](https://img.shields.io/github/license/bububa/creem-go.svg)](https://github.com/bububa/creem-go/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/bububa/creem-go.svg)](https://gitHub.com/bububa/creem-go/releases/)

## Installation

```bash
go install github.com/bububa/creem-go
```

## Usage

### Create Checkout Session

```golang
import (
    "github.com/bububa/creem-go"
    "github.com/bububa/creem-go/checkouts"
)
func main() {
  clt := creem.New(os.Getenv("CREEM_KEY"))
  req := checkouts.CreateRequest {
    ProductID: "xxx",
  }
  var ret checkouts.Checkout
  if err := checkouts.Create(context.Background(), &req, &ret); err != nil {
    panic(err)
  }
  fmt.Println(ret)
}
```

### Get Checkout Session

```golang
import (
    "github.com/bububa/creem-go"
    "github.com/bububa/creem-go/checkouts"
)
func main() {
  clt := creem.New(os.Getenv("CREEM_KEY"))
  checkoutID := "xxx"
  var ret checkouts.Checkout
  if err := checkouts.Get(context.Background(), checkoutID, &ret); err != nil {
    panic(err)
  }
  fmt.Println(ret)
}
```

### Webhook

```golang
package main

import (
    "fmt"
    "net/http"

    "github.com/bububa/creem-go/webhooks"
)

func eventHanderl(ctx context.Context, ev *webhooks.Event) error {
    fmt.Prinf("EVENT: %+v\n", ev)
    return nil
}

func main() {
    handler := webhooks.NewHandler(os.Getenv("CREEM_KEY"), eventHandler)
    http.HandleFunc("/", handler.ServeHTTP)

    port := ":8080"
    fmt.Printf("Starting server on http://localhost%s\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        panic(err)
    }
}
```

For more details please check in pkg.go.dev
