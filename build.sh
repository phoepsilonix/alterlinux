#!/usr/bin/env bash

script_path=$(
    cd "$(dirname "${0}")" || exit 1
    pwd
)
go run "${script_path}/alteriso5/main.go" build "$script_path/profile"
