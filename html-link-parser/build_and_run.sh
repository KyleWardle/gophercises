#! /bin/bash
rm -rf ./html-link-parser
go build
./html-link-parser "$@"
