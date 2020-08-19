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

// Package bot_manage 机器人/机器人信息和管理
package bot_manage

import (
	"bytes"
	"net/http"

	"github.com/fastwego/feishu"
)

const (
	apiInfo   = "/open-apis/bot/v3/info/"
	apiAdd    = "/open-apis/bot/v4/add"
	apiRemove = "/open-apis/bot/v4/remove"
)

/*
获取机器人信息

获取机器人的基本信息

**权限说明** ：需要启用机器人能力


See: https://open.feishu.cn/document/ukTMukTMukTM/uAjMxEjLwITMx4CMyETM

GET https://open.feishu.cn/open-apis/bot/v3/info/
*/
func Info(ctx *feishu.App) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiInfo, header)
}

/*
拉机器人进群

拉机器人进群

**权限说明** ：需要启用机器人能力；机器人的owner需要已经在群里


See: https://open.feishu.cn/document/ukTMukTMukTM/uYDO04iN4QjL2gDN

POST https://open.feishu.cn/open-apis/bot/v4/add
*/
func Add(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiAdd, bytes.NewReader(payload), header)
}

/*
将机器人移出群

将机器人移出群。

**权限说明** ：需要启用机器人能力


See: https://open.feishu.cn/document/ukTMukTMukTM/ucDO04yN4QjL3gDN

POST https://open.feishu.cn/open-apis/bot/v4/remove
*/
func Remove(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiRemove, bytes.NewReader(payload), header)
}
