#!/bin/sh

FILES=`ls . | egrep "*.json"`
for file in $FILES
do
    echo "Testing $file"
    newman run "$file" --suppress-exit-code 1
done
