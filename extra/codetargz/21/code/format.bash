#!/usr/bin/env bash
find . -name '*.go' -exec gofmt -w {} \;
