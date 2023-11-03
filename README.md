# gRPC Master Class

## Section 1: Course Overview

What is gRPC?

-   Remote procedure call
-   Built on top of HTTP 2.0
-   Supports streaming of data

## Section 2: Code Download

[Github](https://github.com/Clement-Jean/grpc-go-course)

## Section 3: Internals Deep Dive

Why gRPC?

-   Efficiency serialization/deserialization over JSON

HTTP 1.0

-   opens a TCP connection for every request
-   plaintext headers (bigger latency for each call)

HTTP 2.0

-   one TCP connection
-   server push
-   multiplexing
-   headers/data compressed to binary
-   SSL connection required by default

Types of API

1. Unary
2. Server streaming
3. Client streaming
4. Bi-directional streaming

# gRPC Set Up

https://grpc.io/docs/languages/go/quickstart/

```
protoc -Igreet/proto --go_out=. --go_opt=module=github.com/tuvo1106/grpc_master_class --go-grpc_out=. --go-grpc_opt=module=github.com/tuvo1106/grpc_master_class greet/proto/dummy.proto
```

alternatively, `make greet`

-   find missing packages
-   `go mod tidy`
