
syntax = "proto3";

package gin_user_fields_proto;

service GinUserFields {
    rpc FindByPagination(QuerySchema) returns(FindRes);
    rpc FindOne(GinUserFieldsSchema) returns(FindOneRes);
    rpc Create(GinUserFieldsSchema) returns(FindOneRes);
    rpc Update(UpdateSchema) returns(FindRes);
}
//数据库结构
message GinUserFieldsSchema{
	    uint32 uid = 1;
    string sexy = 2;
    string birthdy = 3;
    string id_card = 4;
    string truth_name = 5;
    string created = 6;
    string updated = 7;
}
//更新结构
message UpdateSchema {
    GinUserFieldsSchema conditions = 1;
    GinUserFieldsSchema modifies = 2;
}
//查询结构
message QuerySchema{
    GinUserFieldsSchema conditions = 1;
    int32 page_num = 2;
    int32 page_size = 3;
}
//查询返回对象
message FindOneRes{
    int32 code = 1;
    string msg = 2;
    GinUserFieldsSchema data = 3;
}
//查询返回string
message FindRes{
    int32 code = 1;
    string msg = 2;
    string data= 3;
}