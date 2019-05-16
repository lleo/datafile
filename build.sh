#!/usr/bin/env bash
set -x
BASE_DIR=$GOLLEO/datafile
go build -o $BASE_DIR/dftool/dk $BASE_DIR/dftool/main.go
