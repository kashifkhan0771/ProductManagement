#!/bin/bash
if [ "$(docker ps -q -f name=product-management-mongo-db)" ]; then
  docker rm -f product-management-mongo-db
fi
