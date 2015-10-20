# go-http-max-connections-demo

## Description

Httpd max connections demo by Golang.

## Usage

Unlimited (default) httpd connection mode

```bash
$ go-http-max-connections-demo default
(another shell)
$ curl http://127.0.0.1:8080/
Default
```

Limited httpd connection mode.

```bash
$ go-http-max-connections-demo limited [-m MAX_CONNECTIONS]
(another shell)
$ curl http://127.0.0.1:8080/
Limited
```

If you want to stop this demo, please type CTRL+C.

## Install

To install, use `go get`:

```bash
$ go get -d github.com/koemu/go-http-max-connections-demo
```

## Contribution

1. Fork ([https://github.com/koemu/go-http-max-connections-demo/fork](https://github.com/koemu/go-http-max-connections-demo/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[Yuichiro Saito](https://github.com/koemu)
