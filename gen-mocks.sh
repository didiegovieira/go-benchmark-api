#!/bin/bash

rm -rf test/mocks/*

REGEX="(type ([a-zA-Z]+){1} interface)"

for INTERFACE_PATH in $(find . -regex '.+\.go')
do
  if grep -q " interface {" "$INTERFACE_PATH"; then
     INTERFACE_NAME=""

     str=$(cat $INTERFACE_PATH)
     if [[ $str =~ $REGEX ]]
     then
         INTERFACE_NAME=${BASH_REMATCH[2]}
     fi

     IFS='/' read -r -a array <<< $(echo "$INTERFACE_PATH" | sed 's/_interface\.go/_mock.go/g')
     MOCK_NAME="Mock${array[-2]^}${INTERFACE_NAME}"

     echo "Generating mock for $INTERFACE_PATH with name $MOCK_NAME"

     mockgen -source="$INTERFACE_PATH" -destination=test/mocks/"${array[-2]}_${array[-1]}" -package mocks -mock_names="$INTERFACE_NAME=$MOCK_NAME"
  fi
done