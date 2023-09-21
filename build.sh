#!/bin/bash

while getopts n: flag
do
    case "${flag}" in
        n) network=${OPTARG};;
    esac
done

if [ "$network" = "Mainnet" ] || [ "$network" = "mainnet" ]; then
    go build cmd/mainnet/main.go
    ./main
fi
if [ "$network" = "Devnet" ] || [ "$network" = "devnet" ]; then
    go build cmd/devnet/main.go
    ./main
fi
