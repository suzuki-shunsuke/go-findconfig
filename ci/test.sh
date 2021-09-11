#!/usr/bin/env bash

set -eu

cd "$(dirname "$0")/.."

repo_name=${1:-}
if [ -z "$repo_name" ]; then
  echo "the repository name is required" >&2
  exit 1
fi

bash scripts/test-code-climate.sh "$repo_name"
