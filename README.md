# reverse

reverse connection base on grpc.

![image](./imgs/reverse.png)

## usage

client is for server side client.

service is for server side service.

rpc is reverse rpc server.

1. reverse client build a persistent connection to reverse server when reverse client start.
2. client listen on one port and send grpc request to reverse server with ip, port, reverse client id and service name which run on reverse client
3. reverse server send request to reverse client
4. reverse client connect to client and connect to server and forward the data between the two connections

Don't need to change server code.

## example

[example](./example/README.md)
