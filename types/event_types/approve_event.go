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

const EventTypeLeaveApproval = "leave_approval"

/*

{
     "ts": "1502199207.7171419", // 时间戳
     "uuid": "bc447199585340d1f3728d26b1c0297a",
     "token": "41a9425ea7df4536a7623e38fa321bae", // 校验Toke
     "type": "event_callback", // 事件回调此处固定为event_callback
     "event": {
         "app_id": "cli_xxx",
         "tenant_key":"xxx",
         "type": "leave_approval",
         "instance_code": "xxx", // 审批实例Code
         "employee_id": "xxx",  // 用户id
         "start_time": 1502199207, // 审批发起时间
         "end_time": 1502199307,   // 审批结束时间
         "leave_type": "xxx",      // 请假类型
         "leave_unit": 1,          // 请假最小时长：1：一天，2：半天
         "leave_start_time": "2018-12-01 12:00:00", // 请假开始时间
         "leave_end_time": "2018-12-02 12:00:00",   // 请假结束时间
         "leave_interval": 7200,          // 请假时长，单位（秒）
         "leave_reason": "xxx"     // 请假事由
    }
}

*/

type EventLeaveApproval struct {
	BaseEvent
	Event struct {
		AppID          string `json:"app_id"`
		TenantKey      string `json:"tenant_key"`
		Type           string `json:"type"`
		InstanceCode   string `json:"instance_code"`
		EmployeeID     string `json:"employee_id"`
		StartTime      int    `json:"start_time"`
		EndTime        int    `json:"end_time"`
		LeaveType      string `json:"leave_type"`
		LeaveUnit      int    `json:"leave_unit"`
		LeaveStartTime string `json:"leave_start_time"`
		LeaveEndTime   string `json:"leave_end_time"`
		LeaveInterval  int    `json:"leave_interval"`
		LeaveReason    string `json:"leave_reason"`
	} `json:"event"`
}

const EventTypeLeaveApprovalV2 = "leave_approvalV2"

/*

{
     "ts": "1502199207.7171419", // 时间戳
     "uuid": "bc447199585340d1f3728d26b1c0297a",
     "token": "41a9425ea7df4536a7623e38fa321bae", // 校验Toke
     "type": "event_callback", // 事件回调此处固定为event_callback
     "event": {
         "app_id": "cli_xxx",
         "tenant_key":"xxx",
         "type": "leave_approvalV2",
         "instance_code": "xxx", // 审批实例Code
         "user_id": "xxx",  // 用户id
         "start_time": 1564590532,  // 审批发起时间
         "end_time": 1564590532, // 审批结束时间
         "leave_name": "@i18n@123456",  // 假期名称
         "leave_unit": "DAY",  // 请假最小时长
         "leave_start_time": "2019-10-01 00:00:00",// 请假开始时间
         "leave_end_time":"2019-10-02 00:00:00",// 请假结束时间
         "leave_detail": [
            ["2019-10-01 00:00:00","2019-10-02 00:00:00"]
         ], // 具体请假明细
        "leave_interval": 86400,  // 请假时长，单位（秒）
        "leave_reason": "abc", // 请假事由
        "i18n_resources":[ {
                "locale":"en_us",
                "is_default":true,
                "texts":{
                    "@i18n@123456":"Holiday"
                }
        }  ] // 国际化文案
}


*/

type EventLeaveApprovalV2 struct {
	BaseEvent
	Event struct {
		AppID          string     `json:"app_id"`
		TenantKey      string     `json:"tenant_key"`
		Type           string     `json:"type"`
		InstanceCode   string     `json:"instance_code"`
		UserID         string     `json:"user_id"`
		StartTime      int        `json:"start_time"`
		EndTime        int        `json:"end_time"`
		LeaveName      string     `json:"leave_name"`
		LeaveUnit      string     `json:"leave_unit"`
		LeaveStartTime string     `json:"leave_start_time"`
		LeaveEndTime   string     `json:"leave_end_time"`
		LeaveDetail    [][]string `json:"leave_detail"`
		LeaveInterval  int        `json:"leave_interval"`
		LeaveReason    string     `json:"leave_reason"`
		I18NResources  []struct {
			Locale    string `json:"locale"`
			IsDefault bool   `json:"is_default"`
			Texts     struct {
				I18N123456 string `json:"@i18n@123456"`
			} `json:"texts"`
		} `json:"i18n_resources"`
	} `json:"event"`
}

const EventTypeWorkApproval = "work_approval"

/*

{
     "ts": "1502199207.7171419", //时间戳
     "uuid": "bc447199585340d1f3728d26b1c0297a",
     "token": "41a9425ea7df4536a7623e38fa321bae", //校验Toke
     "type": "event_callback", //事件回调此处固定为event_callback
     "event": {
         "app_id": "cli_xxx",
         "tenant_key":"xxx",
         "type":"work_approval",
         "instance_code": "xxx", // 审批实例Code
         "employee_id": "xxx",  // 用户id
         "start_time": 1502199207, // 审批发起时间
         "end_time": 1502199307,   // 审批结束时间
         "work_type": "xxx",      // 加班类型
         "work_start_time": "2018-12-01 12:00:00", // 加班开始时间
         "work_end_time": "2018-12-02 12:00:00",   // 加班结束时间
         "work_interval": 7200,          // 加班时长，单位（秒）
         "work_reason": "xxx"     // 加班事由
    }
}

*/

type EventWorkApproval struct {
	BaseEvent
	Event struct {
		AppID         string `json:"app_id"`
		TenantKey     string `json:"tenant_key"`
		Type          string `json:"type"`
		InstanceCode  string `json:"instance_code"`
		EmployeeID    string `json:"employee_id"`
		StartTime     int    `json:"start_time"`
		EndTime       int    `json:"end_time"`
		WorkType      string `json:"work_type"`
		WorkStartTime string `json:"work_start_time"`
		WorkEndTime   string `json:"work_end_time"`
		WorkInterval  int    `json:"work_interval"`
		WorkReason    string `json:"work_reason"`
	} `json:"event"`
}

const EventTypeShiftApproval = "shift_approval"

/*

{
     "ts": "1502199207.7171419", //时间戳
     "uuid": "bc447199585340d1f3728d26b1c0297a",
     "token": "41a9425ea7df4536a7623e38fa321bae", //校验Toke
     "type": "event_callback", //事件回调此处固定为event_callback
     "event": {
         "app_id": "cli_xxx",
         "tenant_key":"xxx",
         "type":"shift_approval",
         "instance_code": "xxx", // 审批实例Code
         "employee_id": "xxx",  // 用户id
         "start_time": 1502199207, // 审批发起时间
         "end_time": 1502199307,   // 审批结束时间
         "shift_time": "2018-12-01 12:00:00",    // 换班时间
         "return_time": "2018-12-02 12:00:00",   // 还班时间
         "shift_reason": "xxx"     // 换班事由
    }
}

*/

type EventShiftApproval struct {
	BaseEvent
	Event struct {
		AppID        string `json:"app_id"`
		TenantKey    string `json:"tenant_key"`
		Type         string `json:"type"`
		InstanceCode string `json:"instance_code"`
		EmployeeID   string `json:"employee_id"`
		StartTime    int    `json:"start_time"`
		EndTime      int    `json:"end_time"`
		ShiftTime    string `json:"shift_time"`
		ReturnTime   string `json:"return_time"`
		ShiftReason  string `json:"shift_reason"`
	} `json:"event"`
}

const EventTypeRemedyApproval = "remedy_approval"

/*

{
     "ts": "1502199207.7171419", //时间戳
     "uuid": "bc447199585340d1f3728d26b1c0297a",
     "token": "41a9425ea7df4536a7623e38fa321bae", //校验Toke
     "type": "event_callback", //事件回调此处固定为event_callback
     "event": {
         "app_id": "cli_xxx",
         "tenant_key":"xxx",
         "type":"remedy_approval",
         "instance_code": "xxx", // 审批实例Code
         "employee_id": "xxx",  // 用户id
         "start_time": 1502199207, // 审批发起时间
         "end_time": 1502199307,   // 审批结束时间
         "remedy_time": "2018-12-01 12:00:00",   // 补卡时间
         "remedy_reason": "xxx"    // 补卡原因
    }
}

*/

type EventRemedyApproval struct {
	BaseEvent
	Event struct {
		AppID        string `json:"app_id"`
		TenantKey    string `json:"tenant_key"`
		Type         string `json:"type"`
		InstanceCode string `json:"instance_code"`
		EmployeeID   string `json:"employee_id"`
		StartTime    int    `json:"start_time"`
		EndTime      int    `json:"end_time"`
		RemedyTime   string `json:"remedy_time"`
		RemedyReason string `json:"remedy_reason"`
	} `json:"event"`
}

const EventTypeTripApproval = "trip_approval"

/*

{
     "ts": "1502199207.7171419", //时间戳
     "uuid": "bc447199585340d1f3728d26b1c0297a",
     "token": "41a9425ea7df4536a7623e38fa321bae", //校验Toke
     "type": "event_callback", //事件回调此处固定为event_callback
     "event": {
         "app_id": "cli_xxx",
         "tenant_key":"xxx",
         "type":"trip_approval",
         "instance_code": "xxx", // 审批实例Code
         "employee_id": "xxx",  // 用户id
         "start_time": 1502199207, // 审批发起时间
         "end_time": 1502199307,   // 审批结束时间
         "schedules": [{         // Schedule 结构数组
                "trip_start_time": "2018-12-01 12:00:00", // 行程开始时间
                "trip_end_time":  "2018-12-01 12:00:00", // 行程结束时间
                "trip_interval":  3600, // 行程时长（秒）
                "departure":  "xxx",       // 出发地
                "destination":  "xxx",       // 目的地
                "transportation":  "xxx",       // 目的地
                "trip_type": "单程", // 单程/往返
                "remark": "备注"    // 备注
                },
         ]
         "trip_interval": 3600,    // 行程总时长（秒）
         "trip_reason": "xxx"     // 出差事由
         "trip_peers": ["xxx", "yyy"],   // 同行人
    }
}

*/

type EventTripApproval struct {
	BaseEvent
	Event struct {
		AppID        string `json:"app_id"`
		TenantKey    string `json:"tenant_key"`
		Type         string `json:"type"`
		InstanceCode string `json:"instance_code"`
		EmployeeID   string `json:"employee_id"`
		StartTime    int    `json:"start_time"`
		EndTime      int    `json:"end_time"`
		Schedules    []struct {
			TripStartTime  string `json:"trip_start_time"`
			TripEndTime    string `json:"trip_end_time"`
			TripInterval   int    `json:"trip_interval"`
			Departure      string `json:"departure"`
			Destination    string `json:"destination"`
			Transportation string `json:"transportation"`
			TripType       string `json:"trip_type"`
			Remark         string `json:"remark"`
		} `json:"schedules"`
		TripInterval int      `json:"trip_interval"`
		TripReason   string   `json:"trip_reason"`
		TripPeers    []string `json:"trip_peers"`
	} `json:"event"`
}
