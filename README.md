![GolangCI](https://github.com/lawzava/go-postal/workflows/golangci/badge.svg?branch=main)
[![Version](https://img.shields.io/badge/version-v1.2.0-green.svg)](https://github.com/lawzava/go-postal/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/lawzava/go-postal)](https://goreportcard.com/report/github.com/lawzava/go-postal)
[![Coverage Status](https://coveralls.io/repos/github/lawzava/go-postal/badge.svg?branch=main)](https://coveralls.io/github/lawzava/go-postal?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/lawzava/go-postal.svg)](https://pkg.go.dev/github.com/lawzava/go-postal)

# go-postal

Minimalistic library that checks the format of US ZIP codes (5-digit or ZIP+4) and finds the US state for a ZIP code from its prefix range.

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
