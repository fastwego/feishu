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

// Package async_batch 异步批量接口
package async_batch

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/fastwego/feishu"
)

const (
	apiDepartmentBatchAdd = "/open-apis/contact/v2/department/batch_add"
	apiUserBatchAdd       = "/open-apis/contact/v2/user/batch_add"
	apiTaskGet            = "/open-apis/contact/v2/task/get"
)

/*
批量新增部门


该接口用于向通讯录中批量新增多个部门。
调用该接口需要具有所有新增部门父部门的授权范围。
应用商店应用无权限调用此接口。
调用该接口需要申请  更新通讯录  以及  以应用身份访问通讯录  权限。


See: https://open.feishu.cn/document/ukTMukTMukTM/uMDOwUjLzgDM14yM4ATN

POST https://open.feishu.cn/open-apis/contact/v2/department/batch_add
*/
func DepartmentBatchAdd(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiDepartmentBatchAdd, bytes.NewReader(payload), header)
}

/*
批量新增用户


该接口用于向通讯录中批量新增多个用户。
调用该接口需要具有待添加用户所在部门的通讯录授权范围。
应用商店应用无权限调用此接口。
调用该接口需要申请  更新通讯录  以及  以应用身份访问通讯录  权限。


See: https://open.feishu.cn/document/ukTMukTMukTM/uIDOwUjLygDM14iM4ATN

POST https://open.feishu.cn/open-apis/contact/v2/user/batch_add
*/
func UserBatchAdd(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiUserBatchAdd, bytes.NewReader(payload), header)
}

/*
查询批量任务执行状态


该接口用于查询通讯录异步任务当前的执行状态以及执行结果。
应用商店应用无权限调用此接口。
调用该接口需要申请  更新通讯录  以及  以应用身份访问通讯录  权限。


See: https://open.feishu.cn/document/ukTMukTMukTM/uUDOwUjL1gDM14SN4ATN

GET https://open.feishu.cn/open-apis/contact/v2/task/get?task_id=0123456789abcdef0123456789abcdef
*/
func TaskGet(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiTaskGet+"?"+params.Encode(), header)
}
