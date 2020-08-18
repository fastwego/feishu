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

const EventTypeAddUserToChat = "add_user_to_chat"

/*

{
  "ts": "1502199207.7171419", // 事件发送的时间，一般近似于事件发生的时间。
  "uuid": "bc447199585340d1f3728d26b1c0297a", // 事件的唯一标识
  "token": "41a9425ea7df4536a7623e38fa321bae", // 即Verification Token
  "type": "event_callback", // 此事件此处始终为event_callback
  "event": {
    "app_id": "cli_9c8609450f78d102",
    "chat_id": "oc_9e9619b938c9571c1c3165681cdaead5", // 群聊的id
    "operator": {
      // 用户进出群的操作人。用户主动退群的话，operator 就是user自己
      "open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb", // 员工对此应用的唯一标识，同一员工对不同应用的open_id不同
      "user_id": "ca51d83b" // 即“用户ID”，仅企业自建应用会返回
    },
    "tenant_key": "736588c9260f175d",
    "type": "add_user_to_chat", // 事件类型，add_user_to_chat/remove_user_from_chat/revoke_add_user_from_chat
    "users": [{
        "name": "James",
        "open_id": "ou_706adeb944ab1473b9fb3e7da2a40b68",
        "user_id": "51g97a4g"
      },
      {
        "name": "张三",
        "open_id": "ou_7885357f9922aaa34001b190109e0b48",
        "user_id": "6e125386"
      }
    ]
  }
}

*/

type EventAddUserToChat struct {
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
		Users     []struct {
			Name   string `json:"name"`
			OpenID string `json:"open_id"`
			UserID string `json:"user_id"`
		} `json:"users"`
	} `json:"event"`
}

const EventTypeChatDisband = "chat_disband"

/*

{
	"ts": "1502199207.7171419", // 事件发送的时间，一般近似于事件发生的时间。
	"uuid": "bc447199585340d1f3728d26b1c0297a", // 事件的唯一标识
	"token": "41a9425ea7df4536a7623e38fa321bae", // 即Verification Token
	"type": "event_callback", // 此事件此处始终为event_callback
	"event": {
		"app_id": "cli_9c8609450f78d102",
		"chat_id": "oc_9f2df3c095c9395334bb6e70ced0fa83",
		"operator": {
			"open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb",
			"user_id": "ca51d83b"
		},
		"tenant_key": "736588c9260f175d",
		"type": "chat_disband"
	}
}


*/

type EventChatDisband struct {
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
	} `json:"event"`
}
