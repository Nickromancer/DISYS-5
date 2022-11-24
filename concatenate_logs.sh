#!/usr/bin/env bash
echo "Logs (the numbers refers to the ports of the nodes. Refer to the readme.md for explanation.)" > log.txt
find . -name 'logfile.*' -exec echo -e '\n' >> log.txt \; -exec echo {} >> log.txt \; -exec cat {} >> log.txt \;