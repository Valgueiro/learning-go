# learning go

Based on: [Learning Go](https://learning.oreilly.com/library/view/learning-go/9781492077206)

## First steps

1. Set GOPATH:

```
	export GOPATH=$(pwd)
	export PATH=$PATH:$GOPATH/bin
```
2. `go run <file.go>`

## Install deps

Go uses the repository itself, there is no dependency manager

`go install github.com/rakyll/hey@latest`
