# tiktokBackend

开发流程：先 定义好idl文件。thrift格式

使用kitex命令生成代码。
kitex -module "github.com/demonwuwen/tiktokBackend" -service a.b.c idl/base.thrift
go mod tidy

使用hz命令生成代码
hz new -idl idl/hello.thrift
go mod tidy 