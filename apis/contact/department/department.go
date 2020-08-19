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

// Package department 部门管理
package department

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/fastwego/feishu"
)

const (
	apiInfoGet        = "/open-apis/contact/v1/department/info/get"
	apiSimpleList     = "/open-apis/contact/v1/department/simple/list"
	apiDetailBatchGet = "/open-apis/contact/v1/department/detail/batch_get"
	apiAdd            = "/open-apis/contact/v1/department/add"
	apiDelete         = "/open-apis/contact/v1/department/delete"
	apiUpdate         = "/open-apis/contact/v1/department/update"
)

/*
获取部门详情

该接口用于获取部门详情信息。

**权限说明** ：调用该接口需要具有 以应用身份访问通讯录  以及 [部门数据权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)。应用需要拥有待查询部门的通讯录授权。


See: https://open.feishu.cn/document/ukTMukTMukTM/uAzNz4CM3MjLwczM

GET https://open.feishu.cn/open-apis/contact/v1/department/info/get?department_id=TT-1234
*/
func InfoGet(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiInfoGet+"?"+params.Encode(), header)
}

/*
获取子部门列表


该接口用于获取当前部门子部门列表。

**权限说明**：调用该接口需要申请 获取部门组织架构信息 以及 以应用身份访问通讯录 权限。调用该接口需要具有当前部门的授权范围。企业根部门 ID 为 0，当获取根部门子部门列表时，通讯录授权范围必须为全员权限。


See: https://open.feishu.cn/document/ukTMukTMukTM/ugzN3QjL4czN04CO3cDN

GET https://open.feishu.cn/open-apis/contact/v1/department/simple/list?department_id=TT-1234&page_size=10&fetch_child=true
*/
func SimpleList(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiSimpleList+"?"+params.Encode(), header)
}

/*
批量获取部门详情


该接口用于批量获取部门详情，只返回权限范围内的部门。
**权限说明** : 调用该接口需要申请 以应用身份访问通讯录 以及[部门数据权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)。


See: https://open.feishu.cn/document/ukTMukTMukTM/uczN3QjL3czN04yN3cDN

GET https://open.feishu.cn/open-apis/contact/v1/department/detail/batch_get?department_ids=od-2efe30807a10608754862a63b108828f&department_ids=od-da6427b2adbceb91204d7fa6aeb7e8ff
*/
func DetailBatchGet(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiDetailBatchGet+"?"+params.Encode(), header)
}

/*
新增部门


该接口用于向通讯录中增加新的部门。

**权限说明** ：调用该接口需要申请 更新通讯录 以及 以应用身份访问通讯录 权限。应用需要拥有待新增部门的父部门的通讯录授权。应用商店应用无权限调用接口。



See: https://open.feishu.cn/document/ukTMukTMukTM/uYzNz4iN3MjL2czM

POST https://open.feishu.cn/open-apis/contact/v1/department/add
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
删除部门

该接口用于从通讯录中删除部门。

**权限说明** : 调用该接口需要申请 更新通讯录 以及 以应用身份访问通讯录 权限。应用需要同时拥有待删除部门及其父部门的通讯录授权。应用商店应用无权限调用该接口。



See: https://open.feishu.cn/document/ukTMukTMukTM/ugzNz4CO3MjL4czM

POST https://open.feishu.cn/open-apis/contact/v1/department/delete
*/
func Delete(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiDelete, bytes.NewReader(payload), header)
}

/*
更新部门信息

该接口用于更新通讯录中部门的信息。

**权限说明**：调用该接口需要申请 更新通讯录 以及 以应用身份访问通讯录 权限。调用该接口需要具有该部门以及更新操作涉及的部门的通讯录权限。应用商店应用无权限调用此接口。



See: https://open.feishu.cn/document/ukTMukTMukTM/uczNz4yN3MjL3czM

POST https://open.feishu.cn/open-apis/contact/v1/department/update
*/
func Update(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiUpdate, bytes.NewReader(payload), header)
}
