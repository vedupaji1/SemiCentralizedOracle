#!/bin/bash

if [ $PWD == "/home/codezeros/Desktop/tempSim" ]
then
        rm -rf node2
        cp -r ../cosmos/temp11 node2  && cp config2.yaml node2/config.yaml
        cd node2
        go mod tidy && go run .
else
        echo "This Script Is Only Ment To Run In `/home/codezeros/Desktop/tempSim`"
fi
