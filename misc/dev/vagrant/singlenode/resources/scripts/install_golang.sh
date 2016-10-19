#!/bin/bash

set -e

source $HOME_DIR/.bash_profile

echo "Installing Go..."
if [ ! -f /cache/go1.6.linux-amd64.tar.gz ]; then
    wget -q -P /cache https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
    tar xf /cache/go1.6.linux-amd64.tar.gz -C /usr/local
fi

echo "Installing Glide..."
if [ ! -f /cache/glide-v0.12.2-linux-amd64.tar.gz ]; then
    wget -q -P /cache https://github.com/Masterminds/glide/releases/download/v0.12.2/glide-v0.12.2-linux-amd64.tar.gz
    mkdir -p $GOPATH/bin || true
    tar -xf /cache/glide-v0.12.2-linux-amd64.tar.gz linux-amd64/glide --strip-components=1 -C $GOPATH/bin/glide
fi

echo "Installing Godep..."
if [ ! -f /cache/.godep_v74 ]; then
    wget -q -P /cache https://github.com/tools/godep/releases/download/v74/godep_linux_amd64
    mkdir -p $GOPATH/bin || true
    chmod +x /cache/godep_linux_amd64
    cp /cache/godep_linux_amd64 $GOPATH/bin/godep
    touch /cache/.godep_v74
fi