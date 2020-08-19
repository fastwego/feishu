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

// Package group_manage 机器人/群信息和群管理
package group_manage

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/fastwego/feishu"
)

const (
	apiChatCreate    = "/open-apis/chat/v4/create/"
	apiChatList      = "/open-apis/chat/v4/list"
	apiChatInfo      = "/open-apis/chat/v4/info"
	apiChatUpdate    = "/open-apis/chat/v4/update/"
	apiChatterAdd    = "/open-apis/chat/v4/chatter/add/"
	apiChatterDelete = "/open-apis/chat/v4/chatter/delete/"
	apiDisband       = "/open-apis/chat/v4/disband"
)

/*
创建群

机器人创建群并拉指定用户进群。

**权限说明** ：需要启用机器人能力


See: https://open.feishu.cn/document/ukTMukTMukTM/ukDO5QjL5gTO04SO4kDN

POST https://open.feishu.cn/open-apis/chat/v4/create/
*/
func ChatCreate(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiChatCreate, bytes.NewReader(payload), header)
}

/*
获取群列表

获取机器人所在的群列表。

**权限说明** ：需要启用机器人能力


See: https://open.feishu.cn/document/ukTMukTMukTM/uITO5QjLykTO04iM5kDN

GET https://open.feishu.cn/open-apis/chat/v4/list
*/
func ChatList(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiChatList+"?"+params.Encode(), header)
}

/*
获取群信息

获取群名称、群主 ID、成员列表 ID 等群基本信息。

**权限说明** ：需要启用机器人能力；机器人必须在群里


See: https://open.feishu.cn/document/ukTMukTMukTM/uMTO5QjLzkTO04yM5kDN

GET https://open.feishu.cn/open-apis/chat/v4/info?chat_id=oc_eb9e82d5657777ebf1bb5b9024f549ef
*/
func ChatInfo(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiChatInfo+"?"+params.Encode(), header)
}

/*
更新群信息

更新群名称、群配置、转让群主等。

**权限说明** ：需要启用机器人能力；机器人必须是群主，才能执行修改群配置和转让群主的操作（机器人创建的群，机器人默认是群主。）


See: https://open.feishu.cn/document/ukTMukTMukTM/uYTO5QjL2kTO04iN5kDN

POST https://open.feishu.cn/open-apis/chat/v4/update/
*/
func ChatUpdate(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiChatUpdate, bytes.NewReader(payload), header)
}

/*
拉用户进群

机器人拉用户进群，机器人必须在群里。

**权限说明** ：需要启用机器人能力；机器人必须在群里


See: https://open.feishu.cn/document/ukTMukTMukTM/ucTO5QjL3kTO04yN5kDN

POST https://open.feishu.cn/open-apis/chat/v4/chatter/add/
*/
func ChatterAdd(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiChatterAdd, bytes.NewReader(payload), header)
}

/*
移除用户出群

机器人移除用户出群。

**权限说明** ：需要启用机器人能力；机器人必须是群主（机器人创建的群，机器人默认是群主。）


See: https://open.feishu.cn/document/ukTMukTMukTM/uADMwUjLwADM14CMwATN

POST https://open.feishu.cn/open-apis/chat/v4/chatter/delete/
*/
func ChatterDelete(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiChatterDelete, bytes.NewReader(payload), header)
}

/*
解散群

机器人解散群。

**权限说明** ：需要启用机器人能力；机器人必须是群主（机器人创建的群，机器人默认是群主。）


See: https://open.feishu.cn/document/ukTMukTMukTM/uUDN5QjL1QTO04SN0kDN

POST https://open.feishu.cn/open-apis/chat/v4/disband
*/
func Disband(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiDisband, bytes.NewReader(payload), header)
}
