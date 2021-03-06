package templates

var ProtoTpl = `
syntax = "proto3";

package {{PackageName}};

option go_package = "zhiyong/insure/pack/{{PackageName}};{{PackageName}}";

import "google/api/annotations.proto";
import "google/protobuf/empty.proto";
import "common/common.proto";

service {{UCamelTableName}} {
  // 添加/更新
  rpc Save({{UCamelTableName}}Entity) returns (common.Id) {
    option (google.api.http) = {
      post: "/{{SPackageName}}/{{SCamelTableName}}/save"
      body: "*"
    };
  }
  // 删除
  rpc Delete(common.Id) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      post: "/{{SPackageName}}/{{SCamelTableName}}/delete"
      body: "*"
    };
  }
  // 查询
  // page=1&pageSize=2&name=ddd&ok=8&order=no
  rpc Search(common.SearchRequest) returns ({{UCamelTableName}}SearchResponse) {
    option (google.api.http) = {
      get: "/{{SPackageName}}/{{SCamelTableName}}/search/{param}"
    };
  }
  // 单个
  rpc View(common.Id) returns ({{UCamelTableName}}Entity) {
    option (google.api.http) = {
      get: "/{{SPackageName}}/{{SCamelTableName}}/view/{id}"
    };
  }
}

// 实体
message {{UCamelTableName}}Entity {
	{{TableSchema}}
}

// 列表返回
message {{UCamelTableName}}SearchResponse {
  common.SearchPageResponse pageInfo = 1; // 分页信息
  repeated {{UCamelTableName}}Entity data = 2; // 数据
}
`

var ProtoTplB = `
syntax = "proto3";

package {{TableName}}_proto;

service {{ServerName}} {
    rpc FindByPagination(QuerySchema) returns(FindRes);
    rpc FindOne({{ServerName}}Schema) returns(FindOneRes);
    rpc Create({{ServerName}}Schema) returns(FindOneRes);
    rpc Update(UpdateSchema) returns(FindRes);
}
//数据库结构
message {{ServerName}}Schema{
	{{TableSchema}}
}
//更新结构
message UpdateSchema {
    {{ServerName}}Schema conditions = 1;
    {{ServerName}}Schema modifies = 2;
}
//查询结构
message QuerySchema{
    {{ServerName}}Schema conditions = 1;
    int32 page_num = 2;
    int32 page_size = 3;
}
//查询返回对象
message FindOneRes{
    int32 code = 1;
    string msg = 2;
    {{ServerName}}Schema data = 3;
}
//查询返回string
message FindRes{
    int32 code = 1;
    string msg = 2;
    string data= 3;
}`
