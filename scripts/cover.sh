#!/usr/bin/env bash
if [ -z ${LOGZIO_API_TOKEN} ]; then echo "Must supply a LOGZIO_API_TOKEN env var" && exit 1; fi
go test -coverprofile=coverage.out && go tool cover -html=coverage.out