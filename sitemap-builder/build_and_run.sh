#! /bin/bash
rm -rf ./sitemap-builder
go build
./sitemap-builder "$@"
