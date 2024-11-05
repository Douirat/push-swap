#!/bin/bash

# Build the first executable
go build -o checker ./main/checker/main.go

# Build the second executable
go build -o push-swap ./main/push-swap/main.go

echo "Executables created: push-swap, checker"
 