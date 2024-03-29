#!/bin/bash

openapi-generator generate \
  -i ./openapi.json \
  -g go \
  -o client/ \
  --git-host github.com \
  --git-user-id mkcr-innovations \
  --git-repo-id integration-app-client/client \
  --additional-properties generateInterfaces=true,packageName=client \
  --global-property apis,models,supportingFiles=utils.go:configuration.go:client.go:README.md:go.mod:go.sum,apiTests=false
