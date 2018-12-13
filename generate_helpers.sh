#!/usr/bin/env bash

go run ./cmd/generator/generator.go --out-dir=$PWD
goimports -w $PWD
