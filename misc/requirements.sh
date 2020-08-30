#!/bin/bash

echo "INSTALLING PACKAGES"

file="requirements.txt"

while read -r line
do
go get -v $line
go install -v $line
done < $file

printf "All packages installed correctly"
