# How to grow Go code as a bonsai: the style, the rules, the linters (Definition 2021 Hackaton)

## Build

`go build -buildmode=plugin plugin/plugin.go`

## Run

* `go run ./cmd/main.go ./testdata/*`
* `golangci-lint run -v --timeout 10m0s ./testdata/config/test.go`

## Developers
[Mercuryo.io](https://mercuryo.io)
## Contribute
PRs are welcome!