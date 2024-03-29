#!/bin/bash

# CC=arm-linux-gnueabihf-gcc;CGO_ENABLED=1;GOOS=linux;GOARCH=arm;GOARM=7

#export CC=arm-linux-gnueabi-gcc CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7
export CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7

go build -o build/build_magni

scp build/build_magni root@4.2.0.225:/root/server-debug/

echo kokos | ssh -tt root@4.2.0.225 "cd /root/server-debug; chmod +x build_magni; ./build_magni"