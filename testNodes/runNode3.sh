#!/bin/bash

if [ $PWD == "/home/codezeros/Desktop/tempSim" ]
then
        rm -rf node3
        cp -r ../cosmos/temp11 node3  && cp config3.yaml node3/config.yaml
        cd node3
        go mod tidy && go run .
else
        echo "This Script Is Only Ment To Run In `/home/codezeros/Desktop/tempSim`"
fi
