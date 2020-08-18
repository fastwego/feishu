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

const EventTypeUserAdd = "user_add"

/*

{
    "ts": "1502199207.7171419", //  事件发送的时间，一般近似于事件发生的时间。
    "uuid": "bc447199585340d1f3728d26b1c0297a",  // 事件的唯一标识
    "token": "41a9425ea7df4536a7623e38fa321bae", // 即Verification Token
    "type": "event_callback", // 此事件此处始终为event_callback
    "event": {
         "type": "user_add",    // 事件类型，包括user_add, user_update, user_leave
         "app_id": "cli_xxx",   // 应用ID
         "tenant_key": "xxx",   // 企业标识
         "open_id":"xxx" ,  // 员工对此应用的唯一标识，同一员工对不同应用的open_id不同
         "employee_id":"xxx",    // 即“用户ID”，仅企业自建应用会返回
         "union_id": "xxx" // 用户在ISV下的唯一标识，申请了"获取用户统一ID"权限后返回
    }
}

*/

type EventUserAdd struct {
	BaseEvent
	Event struct {
		Type       string `json:"type"`
		AppID      string `json:"app_id"`
		TenantKey  string `json:"tenant_key"`
		OpenID     string `json:"open_id"`
		EmployeeID string `json:"employee_id"`
		UnionID    string `json:"union_id"`
	} `json:"event"`
}

const EventTypeDeptAdd = "dept_add"

/*

{
    "ts": "1502199207.7171419", //  事件发送的时间，一般近似于事件发生的时间。
    "uuid": "bc447199585340d1f3728d26b1c0297a",  // 事件的唯一标识
    "token": "41a9425ea7df4536a7623e38fa321bae", // 即Verification Token
    "type": "event_callback", // 此事件此处始终为event_callback
     "event": {
         "type": "dept_add",  // 事件类型，包括 dept_add, dept_update, dept_delete
         "app_id": "cli_xxx",  // 应用ID
         "tenant_key": "xxx",           // 企业标识
         "open_department_id":"od-xxx"  // 部门的Id
     }
}

*/

type EventDeptAdd struct {
	BaseEvent
	Event struct {
		Type             string `json:"type"`
		AppID            string `json:"app_id"`
		TenantKey        string `json:"tenant_key"`
		OpenDepartmentID string `json:"open_department_id"`
	} `json:"event"`
}

const EventTypeUserStatusChange = "user_status_change"

/*

{
	"ts": "1502199207.7171419", //  事件发送的时间，一般近似于事件发生的时间。
	"uuid": "bc447199585340d1f3728d26b1c0297a",  // 事件的唯一标识
	"token": "41a9425ea7df4536a7623e38fa321bae", // 即Verification Token
	"type": "event_callback", // 此事件此处始终为event_callback
	"event": {
		"type": "user_status_change",    // 事件类型
		"app_id": "cli_xxx",   // 应用ID
		"tenant_key": "xxx",   // 企业标识
		"open_id":"xxx" ,  // 员工对此应用的唯一标识，同一员工对不同应用的open_id不同
		"employee_id":"xxx",    // 即“用户ID”，仅企业自建应用会返回
        "union_id": "xxx",  //用户统一ID，申请了"获取用户统一ID"权限后返回
		"before_status": { // 变化前的状态
			"is_active": false,        // 账号是否已激活
			"is_frozen": false,       // 账号是否冻结
			"is_resigned": false    // 是否离职
		},
		"change_time": "2020-02-21 16:28:48", // 状态更新的时间
		"current_status": { // 变化后的状态
			"is_active": true,
			"is_frozen": false,
			"is_resigned": false
		}
	}
}

*/

type EventUserStatusChange struct {
	BaseEvent
	Event struct {
		Type         string `json:"type"`
		AppID        string `json:"app_id"`
		TenantKey    string `json:"tenant_key"`
		OpenID       string `json:"open_id"`
		EmployeeID   string `json:"employee_id"`
		UnionID      string `json:"union_id"`
		BeforeStatus struct {
			IsActive   bool `json:"is_active"`
			IsFrozen   bool `json:"is_frozen"`
			IsResigned bool `json:"is_resigned"`
		} `json:"before_status"`
		ChangeTime    string `json:"change_time"`
		CurrentStatus struct {
			IsActive   bool `json:"is_active"`
			IsFrozen   bool `json:"is_frozen"`
			IsResigned bool `json:"is_resigned"`
		} `json:"current_status"`
	} `json:"event"`
}

const EventTypeContactScopeChange = "contact_scope_change"

/*

{
    "ts": "1502199207.7171419", //  事件发送的时间，一般近似于事件发生的时间。
    "uuid": "bc447199585340d1f3728d26b1c0297a",  // 事件的唯一标识
    "token": "41a9425ea7df4536a7623e38fa321bae", // 即Verification Token
    "type": "event_callback", // 此事件此处始终为event_callback
     "event": {
         "type": "contact_scope_change", // 事件类型
         "app_id": "cli_xxx",   // 应用ID
         "tenant_key": "xxx", //企业标识
     }
}

*/

type EventContactScopeChange struct {
	BaseEvent
	Event struct {
		Type      string `json:"type"`
		AppID     string `json:"app_id"`
		TenantKey string `json:"tenant_key"`
	} `json:"event"`
}
