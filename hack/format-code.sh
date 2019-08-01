#!/usr/bin/env bash

pushd $GOPATH/src/github.com/nightfury1204/movie-listing-app

echo "gofmt -s -w *.go cmds models pkg routes"
gofmt -s -w *.go cmds models pkg routes

echo "goimports -w *.go cmds models pkg routes"
goimports -w *.go cmds models pkg routes

popd
