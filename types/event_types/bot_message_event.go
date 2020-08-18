// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package event_types

const (
	EventTypeMessage = "message"
)

type EventMessage struct {
	BaseEvent
	Event struct {
		MsgType string `json:"msg_type"`
	} `json:"event"`
}

const EventTypeAddBot = "add_bot"

/*

{
    "ts": "1502199207.7171419", //  事件发送的时间，一般近似于事件发生的时间。
    "uuid": "bc447199585340d1f3728d26b1c0297a",  // 事件的唯一标识
    "token": "41a9425ea7df4536a7623e38fa321bae", // 即Verification Token
    "type": "event_callback", // 此事件此处始终为event_callback
    "event": {
        "app_id": "cli_9c8609450f78d102",
        "chat_i18n_names": { // 群名称国际化字段
            "en_us": "英文标题",
            "zh_cn": "中文标题"
        },
        "chat_name": "群名称",
        "chat_owner_employee_id": "ca51d83b",// 群主的employee_id（即“用户ID”。如果群主是机器人则没有这个字段，仅企业自建应用返回）
        "chat_owner_name": "xxx", // 群主姓名
        "chat_owner_open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb", // 群主的open_id
        "open_chat_id": "oc_e520983d3e4f5ec7306069bffe672aa3",  // 群聊的id
        "operator_employee_id": "ca51d83b", // 操作者的emplolyee_id ，仅企业自建应用返回
        "operator_name": "yyy", // 操作者姓名
        "operator_open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb",//操作者的open_id
        "owner_is_bot": false, //群主是否是机器人
        "tenant_key": "736588c9260f175d",  // 企业标识
        "type": "add_bot" // 事件类型
    }
}


*/

type EventAddBot struct {
	BaseEvent
	Event struct {
		AppID         string `json:"app_id"`
		ChatI18NNames struct {
			EnUs string `json:"en_us"`
			ZhCn string `json:"zh_cn"`
		} `json:"chat_i18n_names"`
		ChatName            string `json:"chat_name"`
		ChatOwnerEmployeeID string `json:"chat_owner_employee_id"`
		ChatOwnerName       string `json:"chat_owner_name"`
		ChatOwnerOpenID     string `json:"chat_owner_open_id"`
		OpenChatID          string `json:"open_chat_id"`
		OperatorEmployeeID  string `json:"operator_employee_id"`
		OperatorName        string `json:"operator_name"`
		OperatorOpenID      string `json:"operator_open_id"`
		OwnerIsBot          bool   `json:"owner_is_bot"`
		TenantKey           string `json:"tenant_key"`
		Type                string `json:"type"`
	} `json:"event"`
}

const EventTypeRemoveBot = "remove_bot"

/*

{
    "ts": "1502199207.7171419", //  事件发送的时间，一般近似于事件发生的时间。
    "uuid": "bc447199585340d1f3728d26b1c0297a",  // 事件的唯一标识
    "token": "41a9425ea7df4536a7623e38fa321bae", // 即Verification Token
    "type": "event_callback", // 此事件此处始终为event_callback
    "event": {
        "app_id": "cli_9c8609450f78d102",
        "chat_i18n_names": { // 群名称国际化字段
            "en_us": "英文标题",
            "zh_cn": "中文标题"
        },
        "chat_name": "群名称",
        "chat_owner_employee_id": "ca51d83b",// 群主的employee_id（即“用户ID”。如果群主是机器人则没有这个字段，仅企业自建应用返回）
        "chat_owner_name": "xxx", // 群主姓名
        "chat_owner_open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb", // 群主的open_id
        "open_chat_id": "oc_e520983d3e4f5ec7306069bffe672aa3",  // 群聊的id
        "operator_employee_id": "ca51d83b", // 操作者的emplolyee_id ，仅企业自建应用返回
        "operator_name": "yyy", // 操作者姓名
        "operator_open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb",//操作者的open_id
        "owner_is_bot": false, //群主是否是机器人
        "tenant_key": "736588c9260f175d",  // 企业标识
        "type": "remove_bot" // 移除机器人：remove_bot

    }
}

*/

type EventRemoveBot struct {
	BaseEvent
	Event struct {
		AppID         string `json:"app_id"`
		ChatI18NNames struct {
			EnUs string `json:"en_us"`
			ZhCn string `json:"zh_cn"`
		} `json:"chat_i18n_names"`
		ChatName            string `json:"chat_name"`
		ChatOwnerEmployeeID string `json:"chat_owner_employee_id"`
		ChatOwnerName       string `json:"chat_owner_name"`
		ChatOwnerOpenID     string `json:"chat_owner_open_id"`
		OpenChatID          string `json:"open_chat_id"`
		OperatorEmployeeID  string `json:"operator_employee_id"`
		OperatorName        string `json:"operator_name"`
		OperatorOpenID      string `json:"operator_open_id"`
		OwnerIsBot          bool   `json:"owner_is_bot"`
		TenantKey           string `json:"tenant_key"`
		Type                string `json:"type"`
	} `json:"event"`
}

const EventTypeP2PChatCreate = "p2p_chat_create"

/*

{
    "ts": "1502199207.7171419", //  事件发送的时间，一般近似于事件发生的时间。
    "uuid": "bc447199585340d1f3728d26b1c0297a",  // 事件的唯一标识
    "token": "41a9425ea7df4536a7623e38fa321bae", // 即Verification Token
    "type": "event_callback", // 此事件此处始终为event_callback
    "event": {
        "app_id": "cli_9c8609450f78d102",
        "chat_id": "oc_26b66a5eb603162b849f91bcd8815b20", //机器人和用户的会话id
        "operator": { // 会话的发起人。可能是用户，也可能是机器人。
            "open_id": "ou_2d2c0399b53d06fd195bb393cd1e38f2" // 员工对此应用的唯一标识，同一员工对不同应用的open_id不同
            "user_id": "gfa21d92"  // 即“用户ID”，仅企业自建应用会返回
        },
        "tenant_key": "736588c9260f175c",  // 企业标识
        "type": "p2p_chat_create",  // 事件类型
        "user": {  // 会话的用户
            "name": "张三",
            "open_id": "ou_7dede290d6a27698b969a7fd70ca53da",  // 员工对此应用的唯一标识，同一员工对不同应用的open_id不同
            "user_id": "gfa21d92" // 即“用户ID”，仅企业自建应用会返回
        }
    }
}

*/

type EventP2PChatCreate struct {
	BaseEvent
	Event struct {
		AppID    string `json:"app_id"`
		ChatID   string `json:"chat_id"`
		Operator struct {
			OpenID string `json:"open_id"`
			UserID string `json:"user_id"`
		} `json:"operator"`
		TenantKey string `json:"tenant_key"`
		Type      string `json:"type"`
		User      struct {
			Name   string `json:"name"`
			OpenID string `json:"open_id"`
			UserID string `json:"user_id"`
		} `json:"user"`
	} `json:"event"`
}

const MsgTypeText = "text"

/*

{
    "uuid": "41b5f371157e3d5341b38b20396e77e3",
    "token": "2g7als3DgPW6Xp1xEpmcvgVhQG621bFY",//校验Token
    "ts": "1550038209.428520",  //时间戳
    "type": "event_callback",//事件回调此处固定为event_callback
    "event": {
        "type": "message", // 事件类型
        "app_id": "cli_xxx",
        "tenant_key": "xxx", //企业标识
        "root_id": "",
        "parent_id": "",
        "open_chat_id": "oc_5ce6d572455d361153b7cb51da133945",
        "chat_type": "private",//私聊private，群聊group
        "msg_type": "text",    //消息类型
        "open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb",
        "open_message_id": "om_36686ee62209da697d8775375d0c8e88",
        "is_mention": false,
        "text": "<at open_id="xxx">@小助手</at> 消息内容 <at open_id="yyy">@张三</at>",      // 消息文本，可能包含被@的人/机器人。
        "text_without_at_bot":"消息内容 <at open_id="yyy">@张三</at>" //消息内容，会过滤掉at你的机器人的内容
   }
}

*/

type EventMessageText struct {
	EventMessage
	Event struct {
		Type             string `json:"type"`
		AppID            string `json:"app_id"`
		TenantKey        string `json:"tenant_key"`
		RootID           string `json:"root_id"`
		ParentID         string `json:"parent_id"`
		OpenChatID       string `json:"open_chat_id"`
		ChatType         string `json:"chat_type"`
		MsgType          string `json:"msg_type"`
		OpenID           string `json:"open_id"`
		OpenMessageID    string `json:"open_message_id"`
		IsMention        bool   `json:"is_mention"`
		Text             string `json:"text"`
		TextWithoutAtBot string `json:"text_without_at_bot"`
	} `json:"event"`
}

const MsgTypePost = "post"

/*

{
    "uuid": "bc447199585340d1f3728d26b1c0297a",
    "token": "2g7als3DgPW6Xp1xEpmcvgVhQG621bFY",
    "ts": "1550032496.527917",
    "type": "event_callback",
    "event": {
        "type": "message",
        "app_id": "cli_xxx",
        "tenant_key": "xxx", //企业标识
        "root_id": "",
        "parent_id": "",
        "open_chat_id": "oc_5ce6d572455d361153b7cb51da133945",//发消息的open_chat_id
        "chat_type": "private",
        "msg_type": "post",// rich_text和post消息这里统一都是post
        "open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb",//发消息的用户open_id
        "open_message_id": "om_a81fa00ee467b1084ffd80b197b58f4b",//消息id
        "is_mention": false,
        "text": "\u003cp\u003e测试1212\u003c/p\u003e\u003cfigure\u003e\u003cimg src=\"http://p4.pstatp.com/origin/daff000d4967d033705b\" origin-width=\"2456\" origin-height=\"2458\"/\u003e\u003c/figure\u003e",//消息内容
        "text_without_at_bot":"消息内容",//消息内容，会过滤掉at你的机器人的内容
        "title": "测试" ,//消息标题
        "image_keys": [ //富文本里面的图片的keys
           "1867eac8-8006-40be-8549-b7beae0d3c4a",
           "434593af-5269-4db4-8b94-65c6dfd4f35e"
         ],
  }
}

*/

type EventMessagePost struct {
	EventMessage
	Event struct {
		Type             string   `json:"type"`
		AppID            string   `json:"app_id"`
		TenantKey        string   `json:"tenant_key"`
		RootID           string   `json:"root_id"`
		ParentID         string   `json:"parent_id"`
		OpenChatID       string   `json:"open_chat_id"`
		ChatType         string   `json:"chat_type"`
		MsgType          string   `json:"msg_type"`
		OpenID           string   `json:"open_id"`
		OpenMessageID    string   `json:"open_message_id"`
		IsMention        bool     `json:"is_mention"`
		Text             string   `json:"text"`
		TextWithoutAtBot string   `json:"text_without_at_bot"`
		Title            string   `json:"title"`
		ImageKeys        []string `json:"image_keys"`
	} `json:"event"`
}

const MsgTypeImage = "image"

/*

{
    "uuid": "c58e17e9e84824a48e51a562cf969fb3",
    "token": "2g7als3DgPW6Xp1xEpmcvgVhQG621bFY",
    "ts": "1550038110.478493",
    "type": "event_callback",
    "event": {
        "type": "message",
        "app_id": "cli_xxx",
        "tenant_key": "xxx", //企业标识
        "root_id": "",
        "parent_id": "",
        "open_chat_id": "oc_5ce6d572455d361153b7cb51da133945",
        "chat_type": "private",
        "msg_type": "image", //图片消息
        "image_height" :"300",
        "image_width" :"300",
        "open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb",
        "open_message_id": "om_340057d660022bf141eb470859c6114c",
        "is_mention": false,
        "image_key": "cd1ce282-94d1-4154-a326-121b07e4721e", // image_key，获取图片内容请查https://open.feishu.cn/document/ukTMukTMukTM/uYzN5QjL2cTO04iN3kDN

  }
}

*/

type EventMessageImage struct {
	EventMessage
	Event struct {
		Type          string `json:"type"`
		AppID         string `json:"app_id"`
		TenantKey     string `json:"tenant_key"`
		RootID        string `json:"root_id"`
		ParentID      string `json:"parent_id"`
		OpenChatID    string `json:"open_chat_id"`
		ChatType      string `json:"chat_type"`
		MsgType       string `json:"msg_type"`
		ImageHeight   string `json:"image_height"`
		ImageWidth    string `json:"image_width"`
		OpenID        string `json:"open_id"`
		OpenMessageID string `json:"open_message_id"`
		IsMention     bool   `json:"is_mention"`
		ImageKey      string `json:"image_key"`
	} `json:"event"`
}

const MsgTypeFile = "file"

/*

{
    "uuid": "8bbb4b5ba77c316991c41fa9ef0ced8b",
    "token": "ZYcIUHYh8lmZen6jStLgvcXAXbqn2tle",
    "ts": "1576725930.795192",
    "type": "event_callback",
    "event": {
        "app_id": "cli_xxx",
        "tenant_key": "xxx",
        "chat_type": "private",
        "is_mention": false,
        "msg_type": "file",
        "open_chat_id": "oc_a1e061372f7745a543dsd5h3d6d2349a",
        "open_id": "ou_2b940d169b7a4a0c76633984b08ced73",
        "open_message_id": "om_d85c4sd7b209tbb98g693a958bc7185f",
        "parent_id": "",
        "root_id": "",
        "type": "message",
        "file_key": "file_b4do9r9b-3526-4bc4-a568-65fe3695b05g"
  }
}

*/

type EventMessageFile struct {
	EventMessage
	Event struct {
		AppID         string `json:"app_id"`
		TenantKey     string `json:"tenant_key"`
		ChatType      string `json:"chat_type"`
		IsMention     bool   `json:"is_mention"`
		MsgType       string `json:"msg_type"`
		OpenChatID    string `json:"open_chat_id"`
		OpenID        string `json:"open_id"`
		OpenMessageID string `json:"open_message_id"`
		ParentID      string `json:"parent_id"`
		RootID        string `json:"root_id"`
		Type          string `json:"type"`
		FileKey       string `json:"file_key"`
	} `json:"event"`
}

const MsgTypeMergeForward = "merge_forward"

/*

{
    "uuid": "d465df13905458e361ff39af85d96d09",
    "token": "2g7als3DgPW6Xp1xEpmcvgVhQG621bFY",
    "ts": "1550111453.764646",
    "type": "event_callback",
    "event": {
        "type": "message",
        "app_id": "cli_xxx",
        "tenant_key": "xxx", //企业标识
        "root_id": "",
        "parent_id": "",
        "open_chat_id": "oc_5ce6d572455d361153b7cb51da133945",
        "msg_type": "merge_forward",
        "open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb",
        "open_message_id": "om_b3961b120d67259e7495d8eb69488189",
        "is_mention": false,
        "chat_type": "private",
        "chat_id": "6642174187597201422",
        "user": "6610187460791558158",
        "msg_list": [
        {
            "root_id": "",
            "parent_id": "",
            "open_chat_id": "oc_b74c59c68d0f2d0ac65846272499d651",
            "msg_type": "image",
            "open_id": "",
            "open_message_id": "be1000265b014075a869134b20c87633",
            "is_mention": false,
            "image_key": "99295878-5e85-41a3-bb00-0ad63b5b156d",
            "image_url": "https://oapi-staging.zjurl.cn/open-apis/api/v2/file/f/99295878-5e85-41a3-bb00-0ad63b5b156d",
            "create_time": 1550044148
       },
       {
            "root_id": "",
            "parent_id": "",
            "open_chat_id": "oc_b74c59c68d0f2d0ac65846272499d651",
            "msg_type": "text",
            "open_id": "",
            "open_message_id": "om_a96c620f2aa036e3c08abebaec7f09d1",
            "is_mention": false,
            "text": "文本1",
            "create_time": 1550044749
       },
       {
            "root_id": "",
            "parent_id": "",
            "open_chat_id": "oc_b74c59c68d0f2d0ac65846272499d651",
            "msg_type": "post",
            "open_id": "",
            "open_message_id": "om_5d5b1732aa9b997dff23d63146427bb2",
            "is_mention": false,
            "text": "\u003cp\u003e富文本内容\u003c/p\u003e\u003cfigure\u003e\u003cimg src=\"http://p2.pstatp.com/origin/dad90010d9c8fc72f0b0\" origin-width=\"888\" origin-height=\"888\"/\u003e\u003c/figure\u003e",
            "title": "富文本",
            "create_time": 1550044772
       }
       ]
   }
}

*/

type EventMessageMergeForward struct {
	EventMessage
	Event struct {
		Type          string `json:"type"`
		AppID         string `json:"app_id"`
		TenantKey     string `json:"tenant_key"`
		RootID        string `json:"root_id"`
		ParentID      string `json:"parent_id"`
		OpenChatID    string `json:"open_chat_id"`
		MsgType       string `json:"msg_type"`
		OpenID        string `json:"open_id"`
		OpenMessageID string `json:"open_message_id"`
		IsMention     bool   `json:"is_mention"`
		ChatType      string `json:"chat_type"`
		ChatID        string `json:"chat_id"`
		User          string `json:"user"`
		MsgList       []struct {
			RootID        string `json:"root_id"`
			ParentID      string `json:"parent_id"`
			OpenChatID    string `json:"open_chat_id"`
			MsgType       string `json:"msg_type"`
			OpenID        string `json:"open_id"`
			OpenMessageID string `json:"open_message_id"`
			IsMention     bool   `json:"is_mention"`
			ImageKey      string `json:"image_key,omitempty"`
			ImageURL      string `json:"image_url,omitempty"`
			CreateTime    int    `json:"create_time"`
			Text          string `json:"text,omitempty"`
			Title         string `json:"title,omitempty"`
		} `json:"msg_list"`
	} `json:"event"`
}

const EventTypeMessageRead = "message_read"

/*

{
    "ts": "1502199207.7171419", //  事件发送的时间，一般近似于事件发生的时间。
    "uuid": "bc447199585340d1f3728d26b1c0297a",  // 事件的唯一标识
    "token": "41a9425ea7df4536a7623e38fa321bae", // 即Verification Token
    "type": "event_callback", // 此事件此处始终为event_callback
	"event": {
		"app_id": "cli_9c8609450f78d102",
		"open_chat_id": "oc_e520983d3e4f5ec7306069bffe672aa3",
		"open_id": "ou_2d2c0399b53d06fd195bb393cd1e38f2",
		"open_message_ids": ["om_dc13264520392913993dd051dba21dcf"],   // 已读消息列表
		"tenant_key": "xxx",
		"type": "message_read"
	}
}

*/

type EventMessageRead struct {
	BaseEvent
	Event struct {
		AppID          string   `json:"app_id"`
		OpenChatID     string   `json:"open_chat_id"`
		OpenID         string   `json:"open_id"`
		OpenMessageIds []string `json:"open_message_ids"`
		TenantKey      string   `json:"tenant_key"`
		Type           string   `json:"type"`
	} `json:"event"`
}
