![GolangCI](https://github.com/lawzava/go-postal/workflows/golangci/badge.svg?branch=main)
[![Version](https://img.shields.io/badge/version-v1.1.1-green.svg)](https://github.com/lawzava/go-postal/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/lawzava/go-postal)](https://goreportcard.com/report/github.com/lawzava/go-postal)
[![Coverage Status](https://coveralls.io/repos/github/lawzava/go-postal/badge.svg?branch=main)](https://coveralls.io/github/lawzava/go-postal?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/lawzava/go-postal.svg)](https://pkg.go.dev/github.com/lawzava/go-postal)

# go-postal

Minimalistic library for keeping up to date with Postal (Zip) Code checks and US state finder/parser.

## Installation

```
go get github.com/lawzava/go-postal
```

## Usage

```go
package main

import "github.com/lawzava/go-postal"

func main() {
	postal.IsValid("70100") // true
	postal.IsValid("101023") // false
}
```
