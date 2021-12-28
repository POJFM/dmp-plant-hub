#!/bin/bash

# CC=arm-linux-gnueabihf-gcc;CGO_ENABLED=1;GOOS=linux;GOARCH=arm;GOARM=7

export CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7

go build -o build/build_magni

scp build/build_magni pi@Magni:/home/pi/codium-debug/build_magni

echo k0k0s | ssh -tt pi@Magni "cd /home/pi/codium-debug; chmod +x build_magni; sudo ./build_magni"