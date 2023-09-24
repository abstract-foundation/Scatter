#!/bin/bash

while getopts n: flag
do
    case "${flag}" in
        n) network=${OPTARG};;
    esac
done

if [ "`go build cmd/$network/main.go`" = "" ]; then
    time ./main
fi
