#!/bin/sh

GOOS=windows GOARCH=amd64 go build -o "$1_windows_amd64.exe" "$1.go"
GOOS=windows GOARCH=386   go build -o "$1_windows_386.exe" "$1.go"
GOOS=linux GOARCH=amd64   go build -o "$1_linux_amd64" "$1.go"
GOOS=linux GOARCH=386     go build -o "$1_linux_386" "$1.go"
GOOS=darwin GOARCH=amd64  go build -o "$1_darwin_amd64" "$1.go"
GOOS=darwin GOARCH=386    go build -o "$1_darwin_386" "$1.go"
