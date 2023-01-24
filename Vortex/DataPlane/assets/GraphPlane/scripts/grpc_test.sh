#/bin/bash

curl -X POST -k http://localhost:4000/v1/echo -H "Content-Type: application/grpc" -d '{"value": "foo"}'

curl -X POST -k https://localhost:4000/v1/echo -H "Content-Type: application/grpc" -d '{"value": "foo"}'