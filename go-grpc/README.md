# GRPC - A high-performance, open-source universal RPC framework (Remote Procedure Call)

## Why GRPC?
- Provides a common API framework to run in any environment.
	- In gRPC, call a remote application as if it were a local object, making it easier for you to create distributed applications and services.
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

## GRPC Features

### [Protocol Buffers](https://developers.google.com/protocol-buffers)
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

### Streaming
Apart from the Unary RPC, gRPC leverages HTTP/2 streaming capabilities:
- Server streaming RPC
```
// Obtain the files inside the folder
// Results are streamed one after another like:
// Server performs multiple calls to "stream.Send(file)" while
// client performs "file = stream.Recv()" until EOF.
rpc OpenFolder(Folder) returns (stream File) {}
```
- Client streaming RPC
```
// Create list of files through client streaming
// Client performs multiple calls to "stream.Send(file)" while
// server performs "file = stream.Recv()" until EOF.
// Finally, client performs "stream.CloseAndRecv()" and the
// server respond with "stream.SendAndClose(Folder)"
rpc CreateFiles(stream File) returns (Folder) {}
```
- Bidirectional streaming RPC
```
// Both client and server exchange their files in a folder
// They can either perform stream.Send() or stream.Recv() based on
// mutual agreement or using multiple routines.
rpc ReconcileFolder(stream File) returns (stream File) {}
```
### HTTP/2
gRPC is built on HTTP/2 to bring many advanced capabilities:
- Binary Framing Layer - request/response is divided into small messages and framed in binary format.
- Streaming - Bidirectional full-duplex streaming
- Flow Control - Enabling detailed control of memory used to buffer inflight messages. Flow control is a mechanism to ensure that a receiver of messages does not get overwhelmed by a fast sender. Flow control prevents data loss, improves performance and increases reliability. It applies to streaming RPCs and is not relevant for unary RPCs.
- Header Compression - Everything in HTTP/2 is encoded to improve performance. Using the HPACK compression method, HTTP/2 only shares the value different from the previous HTTP header packets.
- Processing - With HTTP/2, gRPC supports both synchronous and asynchronous processing. gRPC-go is not supported for Asynchronous calls as of now.

### Channels
Channels are a core concept in gRPC. The HTTP/2 streams allow many simultaneous streams on one connection; channels extend this concept by supporting multiple streams over multiple concurrent connections.

### Authentication
The following authentication mechanisms are built-in to gRPC:
- SSL/TLS
- ALTS
	- Application Layer Transport Security is a mutual authentication and transport encryption system developed by Google.
- Token-based authentication with Google
	- gRPC provides a generic mechanism (described below) to attach metadata based credentials to requests and responses. Additional support for acquiring access tokens (typically OAuth2 tokens) while accessing Google APIs.

Read [GRPC guides](https://grpc.io/docs/guides/) to learn about other features.

## Advantages
	- gRPC is not a replacement for REST. We need to choose the gRPC/REST depends on the requirements:
	- gRPC can work on low bandwidth with high performance.
	- Rapid development with multiple language platforms.
	- By different evaluations, gRPC offers up to 10x faster performance and API-security than REST+JSON communication as it uses Protobuf and HTTP/2
	- gRPC supports client- or server-side streaming semantics
	- The prime feature of gRPC methodology is the native code generation for client/server applications. gRPC frameworks use protoc compiler to generate code from the .proto file.

## Disadvantages
	- Not supported by web browsers unlike REST.
	- gRPC always uses POST request which a threat to web security.
	- Caching not supported as a result. In fact, it seems gRPC won't like caching.
	- Not transparent and in binary form which is difficult to understand
	- Difficult to implement clients without the GRPC stubs. REST can send requests using HTTP or curl commands or any other REST Client for test purposes. Now, grpcurl helps to ease the testing with gRPC reflection support.
	- Not widely adopted yet like REST.
