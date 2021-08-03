#!/bin/bash

if [ -f "./main.exe" ]
    then rm main
fi

go build main.go
./main