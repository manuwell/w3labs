#!/bin/bash
set -x

echo "Setting the envs"
set -a  
source ../.env.local
set +a
echo "Envs set"


echo "Booting localstack ... "
docker compose up -d 

sleep 5
docker compose logs

echo "Running the server ... "
go run ../cmd/server
echo "Server shutdown"