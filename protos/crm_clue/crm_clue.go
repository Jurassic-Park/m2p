
syntax = "proto3";

package crm_clue_proto;

service CrmClue {
    rpc FindByPagination(QuerySchema) returns(FindRes);
    rpc FindOne(CrmClueSchema) returns(FindOneRes);
    rpc Create(CrmClueSchema) returns(FindOneRes);
    rpc Update(UpdateSchema) returns(FindRes);
}
//数据库结构
message CrmClueSchema{
	int32 id = 1;
    int64 clue_day_id = 2;
    uint32 user_id = 3;
    string hotline = 4;
    int32 crm_from = 5;
    int32 teleid = 6;
    int32 customer_level = 7;
    uint32 clue_type = 8;
    int32 clue_sub_type = 9;
    uint32 clue_source = 10;
    int32 expect_prov = 11;
    int32 expect_city = 12;
    uint32 dist_id = 13;
    int32 settle_prov = 14;
    int32 settle_city = 15;
    int32 credit_status = 16;
    uint32 is_buytime = 17;
    string remark = 18;
    int32 is_financial = 19;
    int32 is_urgent = 20;
    string followup_time = 21;
    string first_visit_time = 22;
    string visit_time = 23;
    int32 is_revisit = 24;
    string next_revisit_time = 25;
    string next_contact_time = 26;
    int32 is_danger = 27;
    int32 is_change_saler = 28;
    int32 local_store = 29;
    string deal_final_time = 30;
    uint32 salerid = 31;
    int32 original_salerid = 32;
    string assistants = 33;
    uint32 deal_status = 34;
    int32 deal_sub_status = 35;
    string deal_ids = 36;
    uint32 deal_price = 37;
    int32 buy_price = 38;
    int32 from_clue = 39;
    uint32 shop_status = 40;
    string need_brand = 41;
    string need_series = 42;
    uint32 zhigou = 43;
    int32 create_from = 44;
    uint32 is_delete = 45;
    string create_time = 46;
    uint32 create_operator = 47;
    string update_time = 48;
    uint32 update_operator = 49;
    string lastrevisit_time = 50;
    int32 lastrevisit_operator = 51;
    int32 transfering = 52;
    string agreement_photo = 53;
    string driving_license = 54;
    int32 is_delay_return = 55;
    int32 clue_class = 56;
    uint32 clue_third_type = 57;
    int32 media_interview = 58;
    int32 business_cooperation = 59;
    int32 notice_sales_contact = 60;
    uint32 team_id = 61;
    int32 init_store = 62;
    int32 dcc_customer_level = 63;
    int32 overtime_type = 64;
    string first_call_time = 65;
    int32 c_pay_status = 66;
    int32 is_dcc = 67;
}
//更新结构
message UpdateSchema {
    CrmClueSchema conditions = 1;
    CrmClueSchema modifies = 2;
}
//查询结构
message QuerySchema{
    CrmClueSchema conditions = 1;
    int32 page_num = 2;
    int32 page_size = 3;
}
//查询返回对象
message FindOneRes{
    int32 code = 1;
    string msg = 2;
    CrmClueSchema data = 3;
}
//查询返回string
message FindRes{
    int32 code = 1;
    string msg = 2;
    string data= 3;
}