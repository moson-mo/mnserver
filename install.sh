#!/bin/bash

# copies binary to /usr/local/bin
if [ -e mnserver ]
then
    sudo cp mnserver /usr/local/bin/
else
    sudo -E cp $(go env GOPATH)/bin/mnserver /usr/local/bin/
fi