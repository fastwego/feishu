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
	EventTypeEventReply = "event_reply"
)

/*
{
    "ts": "1502199207.7171419", //时间戳
    "uuid": "bc447199585340d1f3728d26b1c0297a",
    "token": "41a9425ea7df4536a7623e38fa321bae", // 校验 Token
    "type": "event_callback", // 事件回调此处固定为 event_callback
    "event": {
        "app_id" : "cli_xxxxxx",
        "type" : "event_reply", // 事件类型
        "tenant_key" : "xxxxx",
        "event_id": "xxxx", //日程id
        "attendee" : {
           "open_id":"xxx", // 回复此日程的参与人的open_id
           "employee_id":"yyy", // 回复此日程的参与人的employee_id，仅自建应用才会返回
           "union_id": "zzz" // 用户在ISV下的唯一标识，申请了"获取用户统一ID"权限后返回
        },
        "status": "accept", // 对日程的回复：accept 接受, tentative 待定, decline 拒绝
        "display_name":"xxx",
        "reply_timestmap":"1589097034" // 回复时间
    }
}

*/
type EventEventReply struct {
	BaseEvent
	Event struct {
		AppID     string `json:"app_id"`
		Type      string `json:"type"`
		TenantKey string `json:"tenant_key"`
		EventID   string `json:"event_id"`
		Attendee  struct {
			OpenID     string `json:"open_id"`
			EmployeeID string `json:"employee_id"`
			UnionID    string `json:"union_id"`
		} `json:"attendee"`
		Status         string `json:"status"`
		DisplayName    string `json:"display_name"`
		ReplyTimestmap string `json:"reply_timestmap"`
	} `json:"event"`
}
