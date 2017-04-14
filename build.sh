#!/bin/bash

# go get github.com/jteeuwen/go-bindata/...
# go get github.com/elazarl/go-bindata-assetfs/...

cp public/node_modules/layout-grid/dist/js/layout-grid.min.js public/dist/js/
cp public/node_modules/layout-grid/dist/css/layout-grid.min.css public/dist/css/

# GOPATH=~/go/bin
# PATH="$PATH:$GOPATH/bin"

go-bindata-assetfs -ignore='^(public/node_modules|public/dl)' ./public/...

go build .
