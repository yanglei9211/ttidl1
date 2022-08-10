namespace go h2.example

struct H2Req {
	1: string Name (api.query="name")
}

struct H2Resp {
	1: string RespBody
}

struct AddReq {
	1 : i32 x (api.query="x")
	2 : i32 y (api.query="y")
}

struct AddResp {
	1: i32 Ans
}

service H2Service {
	H2Resp H2Method(1: H2Req request) (api.get="/test");
}

service AddService {
    AddResp AddMethod(1: AddReq request) (api.get="/add");
    AddResp SubMethod(1: AddReq request) (api.post="/sub");

}

struct ItemReq {
    1: string subject (api.query="subject")
    2: string item_id (api.query="item_id")
}

struct Q {
    1: string desc
    2: string ans
    3: list<string> tag_ids
}

struct Data {
    1: i32 type
    2: string stem
    3: list<Q> qs
}

struct Item {
    1: string item_id
    2: Data data
}

struct ItemResp {
    1: string item_json
    2: Item item
}

service ItemService {
    ItemResp GetItemMethod(1: ItemReq request) (api.get="get_item")
}
