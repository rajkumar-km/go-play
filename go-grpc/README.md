# GRPC - A high-performance, open-source universal RPC framework (Remote Procedure Call)

## Why GRPC?
- Provides a common API framework to run in any environment.
	- Define your service using Protocol Buffers, a powerful binary serialization toolset/language, simple IDL, and easy interface updating.
	- We just define the service in a .proto file. gRPC automatically generates the idiomatic client and server stubs for your service in a variety of languages and platforms.
	- Efficiently connect services in and across data centers.
	- Run on any environment from servers inside a large data center to your own tablet — all the complexity of communication between different languages and environments is handled for you by gRPC.
	- Start quickly and scale - Install runtime and dev environments with a single line and also scale to millions of RPCs per second with the framework
- Pluggable support available for:
	- Load balancing
	- Tracing
	- Health checking
	- Authentication - fully integrated pluggable authentication with HTTP/2-based transport
- Distributed computing to connect devices, mobile apps, and browsers to connect backend services.
- Bi-directional streaming
	- RPCs can be defined to stream data either from a server, client, or bi-directional.

## [Protocol Buffers](https://developers.google.com/protocol-buffers)
- Google's protocol buffer is used for binary serialization of data models.		
- Protocol buffers are a combination of
	- The definition language (created in .proto files),
	- The code that the proto compiler generates to interface with data, language-specific runtime libraries
	- The serialization format for data written to a file (or sent across a network connection).
- Benefits
	- Compact data storage helps both network traffic and long-term data storage.
	- Fast parsing and uses HTTP 2 instead of 1.1 to speed up apps - said to be 7 to 8 times faster
	- Easily extensible to add/modify/delete fields with backward and forward compatibility
	- Auto-generated classes (in different programming languages) that reduce the developer's work.
- Disadvantages (compared to REST)
	- Not transparent and in binary form which is difficult to understand
	- Difficult to implement clients without the GRPC stubs. REST can send requests using HTTP or curl commands or any other REST Client for test purposes. Now, grpcurl helps to ease the testing with gRPC reflection support.
	- Not widely adopted yet like REST.

## [Installation](https://grpc.io/docs/languages/go/quickstart/)
1. Install Go
```
# Download go from https://go.dev/dl/ for Linux
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.6.linux-amd64.tar.gz
export PATH="$PATH:/usr/local/go/bin"
```

2. Install Protocol Buffers (https://developers.google.com/protocol-buffers)
```
# Download and unzip protoc from https://github.com/protocolbuffers/protobuf/releases
cp -r protoc-23.4-linux-x86_64/include/ /usr/local/

# If you want to use well known include types:
cp -r protoc-23.4-linux-x86_64/include/google /usr/local/include/
```

3. Install Go Plugins for protocol buffers
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# Add in your ~/.bashrc to preserve the PATH settings
export PATH="$PATH:$(go env GOPATH)/bin"
```
