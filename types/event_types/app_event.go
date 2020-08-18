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

const EventTypeAppOpen = "app_open"

/*

{
    "ts": "1502199207.7171419", //  事件发送的时间，一般近似于事件发生的时间。
    "uuid": "bc447199585340d1f3728d26b1c0297a",  // 事件的唯一标识
    "token": "41a9425ea7df4536a7623e38fa321bae", // 即Verification Token
    "type": "event_callback", // 此事件此处始终为event_callback
    "event": {
        "app_id": "cli_xxx",  // 开通的应用ID
        "tenant_key":"xxx",   // 开通应用的企业唯一标识
        "type":"app_open",     // 事件类型
        "applicants": [ // 应用的申请者，可能有多个
            {
                "open_id":"xxx" ,  // 员工对此应用的唯一标识，同一员工对不同应用的open_id不同
                "user_id": "b78cfg77"//安装者的user_id （仅企业自建应用返回 ）
            }
        ],
        "installer": {	// 应用的安装者。如果是自动安装，则没有此字段。
            "open_id":"xxx" ,  // 员工对此应用的唯一标识，同一员工对不同应用的open_id不同
            "user_id": "b78cfg77"//安装者的user_id （仅企业自建应用返回 ）
        }
    }
}

*/

type EventAppOpen struct {
	BaseEvent
	Event struct {
		AppID      string `json:"app_id"`
		TenantKey  string `json:"tenant_key"`
		Type       string `json:"type"`
		Applicants []struct {
			OpenID string `json:"open_id"`
			UserID string `json:"user_id"`
		} `json:"applicants"`
		Installer struct {
			OpenID string `json:"open_id"`
			UserID string `json:"user_id"`
		} `json:"installer"`
	} `json:"event"`
}

const EventTypeAppStatusChange = "app_status_change"

/*

{
    "ts": "1502199207.7171419", //  事件发送的时间，一般近似于事件发生的时间。
    "uuid": "bc447199585340d1f3728d26b1c0297a",  // 事件的唯一标识
    "token": "41a9425ea7df4536a7623e38fa321bae", // 即Verification Token
    "type": "event_callback", // 此事件此处始终为event_callback
    "event": {
        "app_id": "cli_xxx",  // 应用ID
        "tenant_key":"xxx",   // 企业唯一标识
        "type": "app_status_change",  // 事件类型
        "status": "start_by_tenant" //应用状态 start_by_tenant: 租户启用; stop_by_tenant: 租户停用; stop_by_platform: 平台停用
    }
}

*/

type EventAppStatusChange struct {
	BaseEvent
	Event struct {
		AppID     string `json:"app_id"`
		TenantKey string `json:"tenant_key"`
		Type      string `json:"type"`
		Status    string `json:"status"`
	} `json:"event"`
}

const EventTypeOrderPaid = "order_paid"

/*

{
    "ts": "1502199207.7171419", //  事件发送的时间，一般近似于事件发生的时间。
    "uuid": "bc447199585340d1f3728d26b1c0297a",  // 事件的唯一标识
    "token": "41a9425ea7df4536a7623e38fa321bae", // 即Verification Token
    "type": "event_callback", // 此事件此处始终为event_callback
    "event": {
        "type":"order_paid",     // 事件类型
        "app_id": "cli_9daeceab98721136", //应用ID
        "order_id": "6704894492631105539", // 用户购买付费方案时对订单ID 可作为唯一标识
        "price_plan_id": "price_9d86fa1333b8110c",  //付费方案ID
        "price_plan_type": "per_seat_per_month", // 用户购买方案类型
        "seats": 20, // 表示购买了多少人份
        "buy_count":1, //套餐购买数量 目前都为1
        "create_time": "1502199207",
        "pay_time": "1502199209",
        "buy_type": "buy", // 购买类型 buy普通购买 upgrade为升级购买
        "src_order_id": "6704894492631105539", // 当前为升级购买时(buy_type 为upgrade)，该字段表示原订单ID，升级后原订单失效，状态变为已升级(业务方需要处理)
        "order_pay_price":10000//订单支付价格 单位分，
        "tenant_key": "2f98c01bc23f6847"//购买应用的企业标示
    }
}

*/

type EventOrderPaid struct {
	BaseEvent
	Event struct {
		Type          string `json:"type"`
		AppID         string `json:"app_id"`
		OrderID       string `json:"order_id"`
		PricePlanID   string `json:"price_plan_id"`
		PricePlanType string `json:"price_plan_type"`
		Seats         int    `json:"seats"`
		BuyCount      int    `json:"buy_count"`
		CreateTime    string `json:"create_time"`
		PayTime       string `json:"pay_time"`
		BuyType       string `json:"buy_type"`
		SrcOrderID    string `json:"src_order_id"`
		OrderPayPrice int    `json:"order_pay_price"`
		TenantKey     string `json:"tenant_key"`
	} `json:"event"`
}

const EventTypeAppTicket = "app_ticket"

/*

{
    "ts": "1502199207.7171419", //  事件发送的时间，一般近似于事件发生的时间。
    "uuid": "bc447199585340d1f3728d26b1c0297a",  // 事件的唯一标识
    "token": "41a9425ea7df4536a7623e38fa321bae", // 即Verification Token
    "type": "event_callback", // 此事件此处始终为event_callback
    "event": {
         "app_id": "cli_xxx",
         "app_ticket":"xxx",
         "type":"app_ticket"
    }
}

*/

type EventAppTicket struct {
	BaseEvent
	Event struct {
		AppID     string `json:"app_id"`
		AppTicket string `json:"app_ticket"`
		Type      string `json:"type"`
	} `json:"event"`
}

const EventTypeAppUninstalled = "app_uninstalled"

/*

{
    "ts": "1502199207.7171419", //时间戳
    "uuid": "bc447199585340d1f3728d26b1c0297a",
    "token": "41a9425ea7df4536a7623e38fa321bae", // 校验 Token
    "type": "event_callback", // 事件回调此处固定为 event_callback
    "event": {
        "app_id": "cli_xxx", // 被卸载的应用ID
        "tenant_key": "xxx",  // 卸载应用的企业ID
        "type": "app_uninstalled"  // 事件类型
     }
}

*/

type EventAppUninstalled struct {
	BaseEvent
	Event struct {
		AppID     string `json:"app_id"`
		TenantKey string `json:"tenant_key"`
		Type      string `json:"type"`
	} `json:"event"`
}
