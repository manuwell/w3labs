#!/bin/bash
set -x

echo "Setting the envs"
set -a  
source ../.env.local
set +a
echo "Envs set"

echo "Running the server ... "
go run ../cmd/server
echo "Server shutdown"