# Example

server: grpc server from grpc example
client: grpc client from grpc example

## start server

default port 10000

```bash
go run ./example/server/server.go
```

## start reverse server

default port 10001

```bash
go run ./example/reverse_server/reverse_server.go
```

## start reverse client

```bash
go run ./example/reverse_client/reverse_client.go
```

## start client

use port 10002

```bash
go run ./example/client/client.go
```

## TODO

- tls not work
- client listen wait
- reverse client connection leak
