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

// Package user_group 用户群组
package user_group

import (
	"net/http"
	"net/url"

	"github.com/fastwego/feishu"
)

const (
	apiGroupList = "/open-apis/user/v4/group_list"
	apiMembers   = "/open-apis/chat/v4/members"
	apiSearch    = "/open-apis/chat/v4/search"
)

/*
获取用户所在的群列表

获取用户所在的群列表。

**权限说明** ：应用需要“读取群信息”权限；


See: https://open.feishu.cn/document/ukTMukTMukTM/uQzMwUjL0MDM14CNzATN

GET https://open.feishu.cn/open-apis/user/v4/group_list?page_size=2&page_token=6592161138799017988
*/
func GroupList(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiGroupList+"?"+params.Encode(), header)
}

/*
获取群成员列表

如果用户在群中，则返回该群的成员列表。

**权限说明** ：应用需要“读取群信息”权限；


See: https://open.feishu.cn/document/ukTMukTMukTM/uUzMwUjL1MDM14SNzATN

GET https://open.feishu.cn/open-apis/chat/v4/members?chat_id=oc_92c3f700c2ae31369cefee459fb93870&page_token=0&page_size=3
*/
func Members(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiMembers+"?"+params.Encode(), header)
}

/*
搜索用户所在的群列表

搜索用户所在的群列表。

**权限说明** ：应用需要“读取群信息”权限；


See: https://open.feishu.cn/document/ukTMukTMukTM/uUTOyUjL1kjM14SN5ITN

GET https://open.feishu.cn/open-apis/chat/v4/search?page_size=10&page_token=24&query=rd
*/
func Search(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiSearch+"?"+params.Encode(), header)
}
