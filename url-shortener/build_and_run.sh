#! /bin/bash
rm -rf ./url-shortener
go build
./url-shortener "$@"
