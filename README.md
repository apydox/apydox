# apydox
An API documentation hub that'll make your team happy

## Contributing

### Prerequisites

- [protoc](https://grpc.io/docs/protoc-installation/) - protoc is used for the plugin system and is required to generate gRPC Go code from .proto files.
- [Go protoc/gRPC plugins](https://grpc.io/docs/languages/go/quickstart/)

### The plugin system

apydox uses the go-plugin package by Hashicorp to provide the plugin architecture provided for some of the components within apydox.
This uses gRPC to communicate between processes running on the same machine, plugins MUST be running on the same host it's unsafe to have remote plugins and is not supporterd by go-plugin.

### Generating code from proto files

We use [protocol buffer files (*.proto)](https://developers.google.com/protocol-buffers) to define gRPC services and generate Go code from those definitions.

Enter in to the directory containing the .proto file you want to generate Go code for in your terminal.

Then run the following command, replacing `{name}` with actual name of the protobuf definition file:
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative {name}.proto
```
