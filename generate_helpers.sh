#!/usr/bin/env bash

go run ./cmd/generator/generator.go --out-dir=$PWD --tmpl-dir=$PWD/cmd/generator
goimports -w $PWD
