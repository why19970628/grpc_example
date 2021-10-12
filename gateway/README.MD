# grpc-gateway

## docs
https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/

## Testing the gRPC-Gateway

`cd gateway`

`$  go mod tidy && go mod vendor`

Now we can start the server:

`$ go run main.go`

Then we use cURL to send HTTP requests:

`$ curl -X POST -k http://localhost:8090/v1/example/echo -d '{"value": " hello"}'`

`{"message":"hello world"}`
