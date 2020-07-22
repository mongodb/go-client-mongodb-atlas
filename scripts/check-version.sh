#!/usr/bin/env bash
#
# A simple script which ensures that the specified tag
# matches the version defined in the VERSION_FILE.
#
set -ueo pipefail
declare -r VERSION_FILE="mongodbatlas/mongodbatlas.go"

if [[ $# -lt 1 ]]; then
    echo
    echo "Error: TAG not specified"
    echo "Usage: check-version.sh TAG"
    echo
    exit 1
fi
tag="$1"

echo
echo "Checking '$tag'..."
if ! grep -q "Version = \"$tag\"" "$VERSION_FILE"; then
    echo
    echo "ERROR: Mismatch between the release version ('$tag') and '$VERSION_FILE'"
    echo "$VERSION_FILE:"
    echo "---------"
    grep "Version = " "$VERSION_FILE"
    echo "---------"
    echo
    exit 1
fi
echo "OK."
echo
