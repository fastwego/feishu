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

// Package user 用户管理
package user

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/fastwego/feishu"
)

const (
	apiBatchGet                 = "/open-apis/contact/v1/user/batch_get"
	apiDepartmentUserList       = "/open-apis/contact/v1/department/user/list"
	apiDepartmentUserDetailList = "/open-apis/contact/v1/department/user/detail/list"
	apiAdd                      = "/open-apis/contact/v1/user/add"
	apiDelete                   = "/open-apis/contact/v1/user/delete"
	apiUpdate                   = "/open-apis/contact/v1/user/update"
	apiRoleList                 = "/open-apis/contact/v2/role/list"
	apiRoleMembers              = "/open-apis/contact/v2/role/members"
	apiBatchGetId               = "/open-apis/user/v1/batch_get_id"
	apiUnionIdBatchGetList      = "/open-apis/user/v1/union_id/batch_get/list"
	apiSearch                   = "/open-apis/search/v1/user"
)

/*
批量获取用户信息


该接口用于批量获取用户详细信息。

**权限说明** : 调用该接口需要申请“以应用身份访问通讯录”以及[用户数据权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)。请求的用户如果在当前应用的通讯录授权范围内，会返回该用户的详细信息；否则不会返回。


See: https://open.feishu.cn/document/ukTMukTMukTM/uIzNz4iM3MjLyczM

GET https://open.feishu.cn/open-apis/contact/v1/user/batch_get?employee_ids=2fab1234&employee_ids=2f1234cd
*/
func BatchGet(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiBatchGet+"?"+params.Encode(), header)
}

/*
获取部门用户列表


该接口用于获取部门用户列表。
**权限说明**：调用该接口需要申请 获取部门组织架构信息 和 以应用身份访问通讯录 权限。应用需要有被调用部门的通讯录授权


See: https://open.feishu.cn/document/ukTMukTMukTM/uEzNz4SM3MjLxczM

GET https://open.feishu.cn/open-apis/contact/v1/department/user/list?department_id=TT-1234&page_size=100&fetch_child=true
*/
func DepartmentUserList(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiDepartmentUserList+"?"+params.Encode(), header)
}

/*
获取部门用户详情


该接口用于获取部门用户详情信息。
**权限说明**：调用该接口需要申请 获取部门组织架构信息 和 以应用身份访问通讯录 权限，同时根据需要返回字段申请对应的[用户数据权限]。应用需要有被调用部门的通讯录授权。(https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)


See: https://open.feishu.cn/document/ukTMukTMukTM/uYzN3QjL2czN04iN3cDN

GET https://open.feishu.cn/open-apis/contact/v1/department/user/detail/list?department_id=TT-1234&page_size=10&fetch_child=true
*/
func DepartmentUserDetailList(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiDepartmentUserDetailList+"?"+params.Encode(), header)
}

/*
新增用户


该接口用于向通讯录中新增用户。

**权限说明** ：调用该接口需要申请 更新通讯录 以及 以应用身份访问通讯录 权限。新增用户的所有部门必须都在当前应用的通讯录授权范围内才允许新增用户。应用商店应用无权限调用此接口。


See: https://open.feishu.cn/document/ukTMukTMukTM/uMzNz4yM3MjLzczM

POST https://open.feishu.cn/open-apis/contact/v1/user/add
*/
func Add(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiAdd, bytes.NewReader(payload), header)
}

/*
删除用户


该接口用于从通讯录中删除用户。

**权限说明** : 调用该接口需要申请 更新通讯录 以及 以应用身份访问通讯录 权限。应用需要有待删除用户、待删除用户的所有部门的通讯录权限才能删除该用户。应用商店应用无权限调用接口。



See: https://open.feishu.cn/document/ukTMukTMukTM/uUzNz4SN3MjL1czM

POST https://open.feishu.cn/open-apis/contact/v1/user/delete
*/
func Delete(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiDelete, bytes.NewReader(payload), header)
}

/*
更新用户信息


该接口用于更新通讯录中用户信息。

**权限说明** : 调用该接口需要申请 更新通讯录 以及 以应用身份访问通讯录 权限。应用需要拥有待更新用户的通讯录授权，如果涉及到用户部门变更，还需要同时拥有所有新部门的通讯录授权。应用商店应用无权限调用此接口。



See: https://open.feishu.cn/document/ukTMukTMukTM/uQzNz4CN3MjL0czM

POST https://open.feishu.cn/open-apis/contact/v1/user/update
*/
func Update(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiUpdate, bytes.NewReader(payload), header)
}

/*
获取角色列表


该接口用于获取企业的用户角色列表。

**权限说明** ：调用该接口需要申请 获取角色 和 以应用身份访问通讯录 权限


See: https://open.feishu.cn/document/ukTMukTMukTM/uYzMwUjL2MDM14iNzATN

GET https://open.feishu.cn/open-apis/contact/v2/role/list
*/
func RoleList(ctx *feishu.App) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiRoleList, header)
}

/*
获取角色成员列表


该接口用于获取角色下的用户列表。

**权限说明** ：调用该接口需要申请 获取角色 和 以应用身份访问通讯录 权限。角色内不在当前应用通讯录授权范围内的用户不会返回。


See: https://open.feishu.cn/document/ukTMukTMukTM/uczMwUjL3MDM14yNzATN

GET https://open.feishu.cn/open-apis/contact/v2/role/members?role_id=or_846ea69995a259a27cc690182f27de87&page_size=2&page_token=763bd1e74d05e958
*/
func RoleMembers(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiRoleMembers+"?"+params.Encode(), header)
}

/*
使用手机号或邮箱获取用户 ID

ID

根据用户邮箱或手机号查询用户 open_id 和 user_id，支持批量查询。
调用该接口需要申请  通过手机号或者邮箱获取用户ID  权限。

:::warnning
只能查询到应用可用性范围内的用户 ID，不在范围内的用户会表现为不存在。
:::


See: https://open.feishu.cn/document/ukTMukTMukTM/uUzMyUjL1MjM14SNzITN

GET https://open.feishu.cn/open-apis/user/v1/batch_get_id?emails=lisi@z.com&emails=wangwu@z.com&mobiles=13812345678&mobiles=%2b12126668888
*/
func BatchGetId(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiBatchGetId+"?"+params.Encode(), header)
}

/*
使用统一 ID 获取用户 ID

ID 获取用户 ID
使用统一 ID 获取用户 ID信息。
调用该接口需要具有 获取用户统一ID  权限。


See: https://open.feishu.cn/document/ukTMukTMukTM/uUTO5UjL1kTO14SN5kTN

GET https://open.feishu.cn/open-apis/user/v1/union_id/batch_get/list?union_ids=on_94a1ee5551019f18cd73d9f111898cf2&union_ids=on_42f2ef9d07319a4d96fffd7ef5cbfc79
*/
func UnionIdBatchGetList(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiUnionIdBatchGetList+"?"+params.Encode(), header)
}

/*
搜索用户


以用户身份搜索其他用户的信息，无法搜索到外部企业或已离职的用户。
调用该接口需要申请  搜索用户  权限。


See: https://open.feishu.cn/document/ukTMukTMukTM/uMTM4UjLzEDO14yMxgTN

GET https://open.feishu.cn/open-apis/search/v1/user?query=zhangsan&page_size=20&page_token=20
*/
func Search(ctx *feishu.App, params url.Values, header http.Header) (resp []byte, err error) {

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiSearch+"?"+params.Encode(), header)
}
