#!/bin/bash

## UNCOMMENT THESE LINES IF YOU WANT TO BUILD FROM SOURCE
#echo Removing pre-built binary
#rm mangodl
#rm /usr/local/bin/mangodl
#echo Installing the needed dependencies
#go get -d ./...
#echo Building...
#go build main.go
#mv main mangodl

# make the file executable
chmod +x mangodl
# install the executable to /usr/bin
sudo install -m755 mangodl /usr/local/bin/mangodl
