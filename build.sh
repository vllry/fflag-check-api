#!/usr/bin/env bash

protoc -I ./ ./models.proto --go_out=plugins=grpc:.