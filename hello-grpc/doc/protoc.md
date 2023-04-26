

# protoc 使用


protoc 就是 protobuf 的编译器，它把 proto 文件编译成不同的语言

📖 安装，略

📖 使用

$ protoc --help
Usage: protoc [OPTION] PROTO_FILES

  -IPATH, --proto_path=PATH   指定搜索路径
  --plugin=EXECUTABLE:
  
  ....
 
  --cpp_out=OUT_DIR           Generate C++ header and source.

  ....

@<filename>                proto文件的具体位置


1. 搜索路径参数
第一个比较重要的参数就是搜索路径参数，即上述展示的 -IPATH, --proto_path=PATH。
它表示的是我们要在哪个路径下搜索 .proto 文件，这个参数既可以用 -I 指定；
如果不指定该参数，则默认在当前路径下进行搜索；
另外，该参数也可以指定多次，这也意味着我们可以指定多个路径进行搜索。

2. 语言插件参数
语言参数即上述的 --cpp_out= 等，protoc 支持的语言长达13种，且都是比较常见的
运行 help 出现的语言参数，说明protoc本身已经内置该语言对应的编译插件，我们无需安装
Go 语言是由 google 维护，通过 protoc 的插件机制来实现，所以仓库单独维护

3. proto 文件位置参数
proto 文件位置参数，即上述的 @<filename> 参数，指定了我们 proto 文件的具体位置，如 pb/bbs.proto。


📖 语言插件

✨ golang 插件
非内置的语言支持就得自己单独安装语言插件，比如 --go_out= 对应的是 protoc-gen-go，安装命令如下：

# 最新版
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest


注意
protoc-gen-go 要求 pb 文件必须指定 go 包的路径，即
option go_package = "user_growth/pb";


✨ grpc go 插件
在 google.golang.org/protobuf 中，protoc-gen-go 纯粹用来生成 pb 序列化相关的文件，不再承载 gRPC 代码生成功能。

生成 gRPC 相关代码需要安装 grpc-go 相关的插件 protoc-gen-go-grpc

 $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


执行 code gen 命令

$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    easy-demo/hello-grpc/pb/bbs.proto


--go_out、--go-grpc_out=
指定 go、grpc 代码生成的基本路径

--go_opt、--go-grpc_opt
设定插件参数（例，paths=import/source_relative，再来一道路径前缀








