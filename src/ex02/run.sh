#!/usr/bin/env bash

go test -benchmem github.com/litleleprikon/golangconf2020/src/ex02 -run=^$ -bench '^(BenchmarkSample)$' -v