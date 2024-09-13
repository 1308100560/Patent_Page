#!/bin/zsh
git add .
git commit -m "update"
git push
curl -X POST http://10.112.138.178:8008/
