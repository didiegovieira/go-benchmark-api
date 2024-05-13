#!/bin/bash

rm -rf tests/mocks/*

for INTERFACE_PATH in $(find . -regex '.+_interface\.go')
do
  echo "Generating mock for $INTERFACE_PATH"
  IFS='/' read -r -a array <<< $(echo "$INTERFACE_PATH" | sed 's/_interface\.go/_mock.go/g')
  mockgen -source="$INTERFACE_PATH" -destination=tests/mocks/"${array[-2]}_${array[-1]}" -package mocks
done