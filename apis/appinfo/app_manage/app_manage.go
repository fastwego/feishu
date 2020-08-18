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

// Package app_manage 应用信息/应用管理
package app_manage

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/fastwego/feishu"
)

const (
	apiIsUserAdmin      = "/open-apis/application/v3/is_user_admin"
	apiAdminScopeGet    = "/open-apis/contact/v1/user/admin_scope/get"
	apiAppVisibility    = "/open-apis/application/v2/app/visibility"
	apiVisibleApps      = "/open-apis/application/v1/user/visible_apps"
	apiAppList          = "/open-apis/application/v3/app/list"
	apiUpdateVisibility = "/open-apis/application/v3/app/update_visibility"
	apiAppAdminUserList = "/open-apis/user/v4/app_admin_user/list"
)

/*
校验应用管理员


该接口用于查询用户是否为应用管理员，需要申请 校验用户是否为应用管理员 权限。


See: https://open.feishu.cn/document/ukTMukTMukTM/uITN1EjLyUTNx4iM1UTM

GET https://open.feishu.cn/open-apis/application/v3/is_user_admin
*/
func IsUserAdmin(ctx *feishu.App) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiIsUserAdmin, header)
}

/*
获取应用管理员管理范围

该接口用于获取应用管理员的管理范围，即该应用管理员能够管理哪些部门。



See: https://open.feishu.cn/document/ukTMukTMukTM/uMzN3QjLzczN04yM3cDN

GET https://open.feishu.cn/open-apis/contact/v1/user/admin_scope/get?employee_id=2fab1234
*/
func AdminScopeGet(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiAdminScopeGet+"?"+params.Encode(), header)
}

/*
获取应用在企业内的可用范围


该接口用于查询应用在该企业内可以被使用的范围，只能被企业自建应用调用且需要“获取应用信息”权限。


See: https://open.feishu.cn/document/ukTMukTMukTM/uIjM3UjLyIzN14iMycTN

GET https://open.feishu.cn/open-apis/application/v2/app/visibility?app_id=cli_9db45f86b7799104&user_page_token=0&user_page_size=100
*/
func AppVisibility(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiAppVisibility+"?"+params.Encode(), header)
}

/*
获取用户可用的应用


该接口用于查询用户可用的应用列表，只能被企业自建应用调用且需要“获取用户可用的应用”权限。


See: https://open.feishu.cn/document/ukTMukTMukTM/uMjM3UjLzIzN14yMycTN

GET https://open.feishu.cn/open-apis/application/v1/user/visible_apps?user_id=79affdge&page_token=0&lang=zh_cn&page_size=5
*/
func VisibleApps(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiVisibleApps+"?"+params.Encode(), header)
}

/*
获取企业安装的应用


该接口用于查询企业安装的应用列表，只能被企业自建应用调用且需要“获取应用信息”权限。


See: https://open.feishu.cn/document/ukTMukTMukTM/uYDN3UjL2QzN14iN0cTN

GET https://open.feishu.cn/open-apis/application/v3/app/list?page_size=5&page_token=0&lang=zh_cn&status=1
*/
func AppList(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiAppList+"?"+params.Encode(), header)
}

/*
更新应用可用范围


该接口用于增加或者删除指定应用被哪些人可用，只能被企业自建应用调用且需要“管理应用可见范围”权限。


See: https://open.feishu.cn/document/ukTMukTMukTM/ucDN3UjL3QzN14yN0cTN

POST https://open.feishu.cn/open-apis/application/v3/app/update_visibility
*/
func UpdateVisibility(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiUpdateVisibility, bytes.NewReader(payload), header)
}

/*
查询应用管理员列表

查询应用管理员列表，返回应用的最新10个管理员账户id列表。

**权限说明** ：
本接口需要申请  获取应用管理员ID  权限才能访问。
回包数据中的user_id 需要申请  获取用户 userid  权限才会返回
回包数据中的union_id 需要申请  获取用户统一ID  权限才会返回


See: https://open.feishu.cn/document/ukTMukTMukTM/ucDOwYjL3gDM24yN4AjN

GET https://open.feishu.cn/open-apis/user/v4/app_admin_user/list
*/
func AppAdminUserList(ctx *feishu.App) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiAppAdminUserList, header)
}
