#!/bin/zsh

# SPDX-License-Identifier: MIT

for file in ./raw_data/**/messages.json; do
    new_filename=$(echo "$file" | sed 's/.json/.jsonl/g')
    jq -c '.[]' "$file" > "$new_filename"
done
