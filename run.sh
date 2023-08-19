#!/bin/bash
go build
./jq -fs "input" "key"
