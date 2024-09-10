#!/bin/bash

# Create cluster
kind create cluster --name=go-mail-api-cluster --config=infra/kind/kind-config.yml
kubectl cluster-info --context kind-go-mail-api-cluster

# Build the go-mail-api image and load into the (kind) cluster
IMAGE_NAME="go-mail-api-app" # dev
IMAGE_TAG="latest"
# Check if the image exists locally
if [[ "$(docker images -q ${IMAGE_NAME}:${IMAGE_TAG} 2> /dev/null)" == "" ]]; then
  echo "Image not found locally. Building the image..."
  docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .
else
  echo "Image already exists locally. Skipping build."
fi
# Load the image into the kind cluster
kind load docker-image "${IMAGE_NAME}:${IMAGE_TAG}" --name go-mail-api-cluster

# Build the postgres image and load into the (kind) cluster
POSTGRES_IMAGE_NAME="postgres-mail-db" # dev
POSTGRES_IMAGE_TAG="latest"
# Check if the image exists locally
if [[ "$(docker images -q ${POSTGRES_IMAGE_NAME}:${POSTGRES_IMAGE_TAG} 2> /dev/null)" == "" ]]; then
  echo "Image not found locally. Building the image..."
  docker build -t ${POSTGRES_IMAGE_NAME}:${POSTGRES_IMAGE_TAG} ./postgres
else
  echo "Image already exists locally. Skipping build."
fi
# Load the image into the kind cluster
kind load docker-image ${POSTGRES_IMAGE_NAME}:${POSTGRES_IMAGE_TAG} --name go-mail-api-cluster