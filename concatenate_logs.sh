#!/usr/bin/env bash
echo "# Logs" > log.txt
find . -name 'logfile.*' -exec echo -e '\n' >> log.txt \; -exec echo {} >> log.txt \; -exec cat {} >> log.txt \;