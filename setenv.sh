#!/bin/bash
if [ $# -ne 2 ]; then
	echo "usage: setenv.sh <year> <day>"
	exit 1
fi

kitten @ send-text "export year=$1; export day=$2"
