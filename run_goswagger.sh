#!/bin/bash

SCRIPT_PATH="$(cd $(dirname $0) && pwd)"

SWAGGER_GEN="docker run --rm -it  --user $(id -u):$(id -g) -e GOPATH=$(go env GOPATH):/go -v "$SCRIPT_PATH:/code" -w /code quay.io/goswagger/swagger"
$SWAGGER_GEN generate server --target . --name Hta --spec swagger.yaml --principal schemas.User
