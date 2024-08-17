#!/usr/bin/env bash

script_path=$(
    cd "$(dirname "${0}")" || exit 1
    pwd
)

mkarchiso "$script_path/archisoprofile/xfce"

