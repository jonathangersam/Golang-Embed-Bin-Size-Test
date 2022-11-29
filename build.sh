#!/bin/bash

# cleanup old binary
rm main1 main2

# print dir contents and file sizes
ls -l

# build it
echo "building..."
go build main1.go
go build main2.go
echo "build done"

# show file sizes
ls -l
