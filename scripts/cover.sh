#!/usr/bin/env bash
if [ -z ${LOGZIO_API_TOKEN} ]; then echo "Must supply a LOGZIO_API_TOKEN env var" && exit 1; fi
PREFIX=github.com/logzio/logzio_terraform_client
go test ${PREFIX}/alerts ${PREFIX}/client ${PREFIX}/endpoints ${PREFIX}/users -coverprofile=coverage.out && go tool cover -html=coverage.out