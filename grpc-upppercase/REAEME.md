# GRPC简单的Server Client通信示例

目录结构：

```bash
.
├── client
│   └── client.go
├── gen
│   ├── uppercase_grpc.pb.go
│   └── uppercase.pb.go
├── go.mod
├── go.sum
├── Makefile
├── proto
│   └── uppercase.proto
├── REAEME.md
└── server
    └── server.go
```

## 具体文件内容

1. Makefile

```makefile
.PHONY: gen

gen:
 protoc \
  --proto_path=proto \
  --go_out=gen \
  --go_opt=paths=source_relative \
  --go-grpc_out=gen \
  --go-grpc_opt=paths=source_relative \
  uppercase.proto

```

其中 `proto_path` 告诉 protoc compiler 在哪里读取 `.proto` 文件。`go_out` 告诉 compiler 输出到哪个目录。 `go_opt=paths=source_relative` 一般会固定使用这个设置，它表示输出的文件和输入文件的使用相同的目录结构。比如 `foo.proto` 将会在 out 目录生成 `foo.pb.go`， 而 `bar/foo.proto` 会输出在 out 目录下的 `bar/foo.pb.go`。 其中 out 目录需要手动建立， 但是中间的目录 compiler 会自动创建。

## 安装Go的gRPC和Protocol Buffers插件

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

其中 `protoc-gen-go` 用于生成类型和相应的序列化和反序列化代码， `protoc-gen-go-grpc` 用于生成 rpc 服务代码。注意一定要把插件添加到 PATH 环境变量里面， compiler 才能找到这些插件。

## 安装相应的go module

```bash
go get google.golang.org/grpc
go get google.golang.org/protobuf
```

## 运行程序

```bash
go run server/server.go
go run client/client.go
```

## 参考资料

1. [gRPC教程](https://juejin.cn/post/7191008929986379836)
