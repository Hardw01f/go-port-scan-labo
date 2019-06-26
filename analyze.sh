#!/bin/sh
set -e
set pipefail

cat portscan.log | awk '!match($1,"---"){print $0}' | sort | uniq -c | sort
