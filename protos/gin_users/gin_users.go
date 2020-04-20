
syntax = "proto3";

package gin_users_proto;

service GinUsers {
    rpc FindByPagination(QuerySchema) returns(FindRes);
    rpc FindOne(GinUsersSchema) returns(FindOneRes);
    rpc Create(GinUsersSchema) returns(FindOneRes);
    rpc Update(UpdateSchema) returns(FindRes);
}
//数据库结构
message GinUsersSchema{
	uint32 id = 1;
    string username = 2;
    string fullname = 3;
    string password = 4;
    string mobile = 5;
    string email = 6;
    string create_time = 7;
    string update_time = 8;
}
//更新结构
message UpdateSchema {
    GinUsersSchema conditions = 1;
    GinUsersSchema modifies = 2;
}
//查询结构
message QuerySchema{
    GinUsersSchema conditions = 1;
    int32 page_num = 2;
    int32 page_size = 3;
}
//查询返回对象
message FindOneRes{
    int32 code = 1;
    string msg = 2;
    GinUsersSchema data = 3;
}
//查询返回string
message FindRes{
    int32 code = 1;
    string msg = 2;
    string data= 3;
}