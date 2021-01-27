#!/usr/bin/env bash

# Copyright 2021 MongoDB Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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
