#!/bin/bash

if [ $PWD == "/home/codezeros/Desktop/tempSim" ]
then
        rm -rf node4
        cp -r ../cosmos/temp11 node4  && cp config4.yaml node4/config.yaml
        cd node4
        go mod tidy && go run .
else
        echo "This Script Is Only Ment To Run In `/home/codezeros/Desktop/tempSim`"
fi
