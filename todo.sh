#!/bin/bash

exec 3> TODO_auto.md

echo "# TODO" >&3
echo >&3

egrep -R TODO . \
        --exclude-dir=tmp \
        --exclude-dir=.git \
        --exclude=todo.sh \
        --exclude=TODO*.md | \
sed -E 's|(\./[^:]*):.*\[TODO\] *(.*)|- [[LINK](\1)]  \2   |g' >&3

exec 3>&-