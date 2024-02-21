# Punycode Encoder and Decoder

This repository contains two simple Go programs for encoding and decoding Punycode strings. Punycode is used to represent Unicode domain names with non-ASCII characters in ASCII-compatible encoding.

## punyencode.go

`punyencode.go` is a Go program that encodes a non-ASCII domain name into its Punycode representation.

### Usage

```bash
go run punyencode.go <non-ASCII>

Example:

go run punyencode.go bücher.ch


This will output the Punycode representation of the non-ASCII domain name bücher.ch.


punydecode.go
punydecode.go is a Go program that decodes a Punycode string into its non-ASCII domain name.

Usage:

go run punydecode.go bücher.ch


go run punydecode.go xn--bcher-kva.ch




