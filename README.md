# PQC_Lab web application

## Instructions

- **Protoc**

```
protoc --proto_path=$SRC_DIR/internals/ \
       --go_out=$SRC_DIR/internals/ \
       --go-grpc_out=$SRC_DIR/internals/server.proto
```


