#! /bin/bash
rm -rf ./choose-your-adventure
go build
./choose-your-adventure "$@"
