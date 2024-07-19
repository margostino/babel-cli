#!/bin/bash

# Ensure the script exits if any command fails
set -e

docker compose -f ./docker/docker-compose.yml stop