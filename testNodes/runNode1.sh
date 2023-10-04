#!/bin/bash

if [ $PWD == "/home/codezeros/Desktop/tempSim" ]
then
	rm -rf node1
	cp -r ../cosmos/temp11 node1  && cp config1.yaml node1/config.yaml
	cd node1
	go mod tidy && go run .
else
	echo "This Script Is Only Ment To Run In `/home/codezeros/Desktop/tempSim`"
fi
