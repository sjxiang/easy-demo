

# protobuf 基础


定义消息类型
protobuf里最基本的类型就是message，每一个message都会有一个或者多个字段(field)，其中字段包含如下元素


    类型：类型不仅可以是标量类型（int、string等），也可以是复合类型（enum等），也可以是其他message
    字段名：字段名比较推荐的是使用下划线/分隔名称
    字段编号：一个message内每一个字段编号都必须唯一的，在编码后其实传递的是这个编号而不是字段名

    字段规则：消息字段可以是以下字段之一

singular：格式正确的消息可以有零个或一个字段（但不能超过一个）。使用 proto3 语法时，如果未为给定字段指定其他字段规则，则这是默认字段规则

optional：与 singular 相同，不过您可以检查该值是否明确设置

repeated：在格式正确的消息中，此字段类型可以重复零次或多次。系统会保留重复值的顺序

map：这是一个成对的键值对字段

    保留字段：为了避免再次使用到已移除的字段可以设定保留字段。如果任何未来用户尝试使用这些字段标识符，协议缓冲区编译器就会报错

标量值类
标量类型会涉及到不同语言和编码方式，后续有机会深入讲

.proto Type	Go Type	Notes
double	float64	
float	float32	
int32	int32	使用可变长度的编码。对负数的编码效率低下 - 如果您的字段可能包含负值，请改用 sint32。
int64	int64	使用可变长度的编码。对负数的编码效率低下 - 如果字段可能有负值，请改用 sint64。
uint32	uint32	使用可变长度的编码。
uint64	uint64	使用可变长度的编码。
sint32	int32	使用可变长度的编码。有符号整数值。与常规 int32 相比，这些函数可以更高效地对负数进行编码。
sint64	int64	使用可变长度的编码。有符号整数值。与常规 int64 相比，这些函数可以更高效地对负数进行编码。
fixed32	uint32	始终为 4 个字节。如果值通常大于 2^28，则比 uint32 更高效。
fixed64	uint64	始终为 8 个字节。如果值通常大于 2^56，则比 uint64 更高效。
sfixed32	int32	始终为 4 个字节。
sfixed64	int64	始终为 8 个字节。
bool	bool	
string	string	字符串必须始终包含 UTF-8 编码或 7 位 ASCII 文本，并且长度不得超过 232。
bytes	[]byte	可以包含任意长度的 2^32 字节。
复合类型
数组
message SearchResponse {
  repeated Result results = 1;
}

message Result {
  string url = 1;
  string title = 2;
  repeated string snippets = 3;
}
枚举
message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
  enum Corpus {
    UNIVERSAL = 0;
    WEB = 1;
    IMAGES = 2;
    LOCAL = 3;
    NEWS = 4;
    PRODUCTS = 5;
    VIDEO = 6;
  }
  Corpus corpus = 4;
}


服务
定义的method仅能有一个入参和出参数。如果需要传递多个参数需要定义成message

service SearchService {
  rpc Search(SearchRequest) returns (SearchResponse);
}
使用其他消息类型
使用import引用另外一个文件的pb

syntax = "proto3";

import "google/protobuf/wrappers.proto";

package ecommerce;

message Order {
  string id = 1;
  repeated string items = 2;
  string description = 3;
  float price = 4;
  google.protobuf.StringValue destination = 5;
}



