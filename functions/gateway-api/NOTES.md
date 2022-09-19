# Implementation Nodes

## Protobuf 

```shell
# in base dir git clone git@github.com:envoyproxy/envoy.git
protoc --go_out=. \
    --proto_path ../../envoy/api/envoy/config/ \
    --experimental_allow_proto3_optional \
    io/treactor/v1alpha/atom.proto \
    io/treactor/v1alpha/node.proto
```
