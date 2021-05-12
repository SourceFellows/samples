#!/bin/bash

echo "You have to kill this container because delve ignores SIGINT in headless mode"

dlv exec --api-version 2 --headless --log --listen :4040 /app/main