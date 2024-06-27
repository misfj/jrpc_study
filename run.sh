pwd:jrpc/pb
//其余插件在(https://github.com/grpc-ecosystem/grpc-gateway)
//安装protoc-gen-validate
go install github.com/envoyproxy/protoc-gen-validate
//编译proto文件 (https://www.liwenzhou.com/posts/Go/protobuf/) --validate 生成验证代码,需手动安装protoc-gen-validate

protoc  --go_out=. --validate_out="lang=go:." service.proto
//生成grpc代码
protoc --go_out=. --go-grpc_out=.   service.proto
//生成grpc-gateway 代码
protoc --go_out=.  --go-grpc_out=. --grpc-gateway_out=. service.proto
