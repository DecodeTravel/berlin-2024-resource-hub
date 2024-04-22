#!/usr/bin/env bash
set -e

checksum=8f5b05adf0a77b2965c9c12c7124d59a6f4c89a5d1d18732129fa26dba0121a0
# pass secret key as argument
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <secret key>"
    exit 1
fi


# run executable address
sh_path=$(dirname "${BASH_SOURCE[0]}")
cd "$sh_path"


# if sha256sum of address is not equal to checksum, print error
if ! echo "$checksum  address" | sha256sum -c -; then
    echo "Error: address sha256sum mismatch"
    exit 1
fi
./address "$1"