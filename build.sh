#!/bin/bash

# protoc -I ./pb --go_out=plugins=grpc:./pb ./pb/*.proto

docker build -t local/gcs -f Dockerfile.gcs .
docker build -t local/api -f Dockerfile.api .

# Database
kubectl apply -f database.yaml

# API
kubectl apply -f api.yaml

# gRPC server
kubectl apply -f gcs.yaml