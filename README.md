# zero-demo
go-zero Demo

1、在rpc包中由于业务的分类，会出现多个.proto的文件 所以不能再用之前的rpcgen工具，rpcgen是配置在.zshrc文件中用alias来弄的。

对于多个.proto文件

protoc -I ./ --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. xxx.proto

然后对还有service的.proto文件 

goctl rpc protoc xxxx.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero

2、如果在rpc服务想用grpcui来调试的话。.yaml文件中需要加上Mode: dev不然这个工具会一直连不上
