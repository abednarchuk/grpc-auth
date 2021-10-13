#!/bin/bash
protoc -I=. --js_out=import_style=commonjs,binary:. --grpc-web_out=import_style=typescript,mode=grpcweb:. auth.proto