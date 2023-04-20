#!/bin/sh
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" $(git rev-parse --show-toplevel)/public/scripts/libs/vendor/wasm
GOOS=js GOARCH=wasm go build -ldflags="-X 'main.secret=AES256Key-32Characters1234567890'" -o $(git rev-parse --show-toplevel)/public/scripts/libs/vendor/wasm/main.wasm $(git rev-parse --show-toplevel)/src/wasm/main.go