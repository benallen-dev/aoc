#!/usr/bin/env bash

# This script is used to create a new day's directory

# Check for the correct number of arguments
if [ $# -ne 2 ]; then
	echo "Usage: $0 <year> <day>"
	exit 1
fi

year=$1
# Pad the day with a 0 if it is less than 10
if [ $2 -lt 10 ]; then
	day=0$2
else
	day=$2
fi

# if we're not in the root directory, cd to it
cd $(git rev-parse --show-toplevel)
mkdir -p ./cmd/$year/$day
cp -r ./template/* ./cmd/$year/$day

touch ./cmd/$year/$day/input.txt
touch ./cmd/$year/$day/input_example.txt

echo "done"
