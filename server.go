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

package feishu

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fastwego/feishu/types/event_types"

	"github.com/fastwego/feishu/util"
)

/*
响应 推送事件 的服务器
*/
type Server struct {
	Ctx *App
}

// Challenge 服务器接口校验
func (s *Server) Challenge(writer http.ResponseWriter, event event_types.EventChallenge) {
	data, err := json.Marshal(struct {
		Challenge string `json:"challenge"`
	}{
		Challenge: event.Challenge,
	})
	if err != nil {
		return
	}
	writer.Write(data)
}

// ParseEvent 解析 推送过来的 事件
func (s *Server) ParseEvent(request *http.Request) (event interface{}, err error) {

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return
	}

	if s.Ctx.Logger != nil {
		s.Ctx.Logger.Println(string(body))
	}

	// 是否加密消息
	if s.Ctx.Config.EncryptKey != "" {
		var encrypt = struct {
			Encrypt string `json:"encrypt"`
		}{}
		err = json.Unmarshal(body, &encrypt)
		if err != nil {
			return
		}

		// 解密
		if encrypt.Encrypt != "" {
			body, err = util.AESDecrypt(encrypt.Encrypt, []byte(s.Ctx.Config.EncryptKey))
			if err != nil {
				return
			}
		}
	}

	cb := event_types.Callback{}
	err = json.Unmarshal(body, &cb)
	if err != nil {
		return
	}

	// 校验 Token
	if len(s.Ctx.Config.VerificationToken) > 0 && cb.Token != s.Ctx.Config.VerificationToken {
		err = fmt.Errorf("cb.Token %s != s.Ctx.Config.VerificationToken %s", cb.Token, s.Ctx.Config.VerificationToken)
		return
	}

	switch cb.Type {
	case event_types.CallbackTypeUrlVerification: // 回调 url 校验
		e := event_types.EventChallenge{}
		err = json.Unmarshal(body, &e)
		if err != nil {
			return
		}
		return e, nil
	case event_types.CallbackTypeEventCallback: // 事件
		e := event_types.BaseEvent{}
		err = json.Unmarshal(body, &e)
		if err != nil {
			return
		}
		switch e.Event.Type {

		// 应用事件
		case event_types.EventTypeAppOpen: //首次开通应用
			e := event_types.EventAppOpen{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil
		case event_types.EventTypeAppStatusChange: // 应用停启用
			e := event_types.EventAppStatusChange{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil
		case event_types.EventTypeOrderPaid: // 应用商店应用购买
			e := event_types.EventOrderPaid{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil
		case event_types.EventTypeAppTicket: // app ticket 推送事件
			e := event_types.EventAppTicket{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil

		case event_types.EventTypeAppUninstalled: // 应用卸载
			e := event_types.EventAppUninstalled{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil

		// 审批事件
		case event_types.EventTypeLeaveApproval: // 请假审批
			e := event_types.EventLeaveApproval{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil
		case event_types.EventTypeLeaveApprovalV2: // 请假审批
			e := event_types.EventLeaveApprovalV2{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil
		case event_types.EventTypeWorkApproval: // 加班审批
			e := event_types.EventWorkApproval{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil
		case event_types.EventTypeShiftApproval: // 换班审批
			e := event_types.EventShiftApproval{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil
		case event_types.EventTypeRemedyApproval: // 补卡审批
			e := event_types.EventRemedyApproval{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil
		case event_types.EventTypeTripApproval: // 出差审批
			e := event_types.EventTripApproval{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil

		// 日历事件
		case event_types.EventTypeEventReply:
			e := event_types.EventEventReply{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil

		// 通讯录事件
		case event_types.EventTypeUserAdd:
			e := event_types.EventUserAdd{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil
		case event_types.EventTypeDeptAdd:
			e := event_types.EventDeptAdd{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil
		case event_types.EventTypeUserStatusChange:
			e := event_types.EventUserStatusChange{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil
		case event_types.EventTypeContactScopeChange:
			e := event_types.EventContactScopeChange{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil

		// 群聊事件
		case event_types.EventTypeAddUserToChat:
			e := event_types.EventAddUserToChat{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil
		case event_types.EventTypeChatDisband:
			e := event_types.EventChatDisband{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil

		// 机器人和消息会话事件
		case event_types.EventTypeAddBot: // 机器人进群
			e := event_types.EventAddBot{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil
		case event_types.EventTypeRemoveBot: // 机器人被移出群
			e := event_types.EventRemoveBot{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil
		case event_types.EventTypeP2PChatCreate: // 用户和机器人的会话首次被创建
			e := event_types.EventP2PChatCreate{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil
		case event_types.EventTypeMessageRead: // 消息已读
			e := event_types.EventMessageRead{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			return e, nil
		case event_types.EventTypeMessage: // 消息事件
			e := event_types.EventMessage{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				return
			}
			switch e.Event.MsgType {
			case event_types.MsgTypeText: // 文本消息
				e := event_types.EventMessageText{}
				err = json.Unmarshal(body, &e)
				if err != nil {
					return
				}
				return e, nil
			case event_types.MsgTypeImage: // 图片消息
				e := event_types.EventMessageImage{}
				err = json.Unmarshal(body, &e)
				if err != nil {
					return
				}
				return e, nil
			case event_types.MsgTypeFile: // 文件消息
				e := event_types.EventMessageFile{}
				err = json.Unmarshal(body, &e)
				if err != nil {
					return
				}
				return e, nil
			case event_types.MsgTypePost: // post x消息
				e := event_types.EventMessagePost{}
				err = json.Unmarshal(body, &e)
				if err != nil {
					return
				}
				return e, nil
			case event_types.MsgTypeMergeForward: // 合并转发
				e := event_types.EventMessageMergeForward{}
				err = json.Unmarshal(body, &e)
				if err != nil {
					return
				}
				return e, nil
			}
		}
	}

	err = errors.New("unknown event")
	return
}
