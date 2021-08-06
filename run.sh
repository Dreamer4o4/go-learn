#!/bin/bash

if [ -f "./main" ]
    then rm main
fi
if [ -f "./main.exe" ]
    then rm main.exe
fi

go build main.go


if [ -f "./main" ]
    then ./main
fi
if [ -f "./main.exe" ]
    then ./main.exe
fi