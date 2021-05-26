#!/bin/bash

## UNCOMMENT THIS LINES IF YOU WANT TO BUILD FROM SOURCE
#echo Removing pre-built binary
#rm mangodl
#echo Installing the needed dependencies
#go get -d ./...
#echo Building...
#go build


# install the executable to /usr/bin
sudo install -m755 mangodl /usr/bin/mangodl