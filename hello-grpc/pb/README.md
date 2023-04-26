
## 生成代码

```shell
$ cd hello-grpc/pb/
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative user.proto
```