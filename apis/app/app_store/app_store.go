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

// Package app_store 应用信息/应用商店
package app_store

import (
	"net/http"
	"net/url"

	"github.com/fastwego/feishu"
)

const (
	apiCheckUser = "/open-apis/pay/v1/paid_scope/check_user"
	apiOrderList = "/open-apis/pay/v1/order/list"
	apiOrderGet  = "/open-apis/pay/v1/order/get"
)

/*
查询用户是否在应用开通范围

该接口用于查询用户是否在企业管理员设置的使用该应用的范围中。如果设置的付费套餐是按人收费或者限制了最大人数，开放平台会引导企业管理员设置“付费功能开通范围”，本接口用于查询用户是否在企业管理员设置的使用该应用的范围中，可以通过此接口，在付费功能点入口判断是否允许某个用户进入使用。

See: https://open.feishu.cn/document/ukTMukTMukTM/uATNwUjLwUDM14CM1ATN

GET https://open.feishu.cn/open-apis/pay/v1/paid_scope/check_user?open_id=ou_5ad573a6411d72b8305fda3a9c15c70e
*/
func CheckUser(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiCheckUser+"?"+params.Encode(), header)
}

/*
查询租户购买的付费方案

该接口用于分页查询应用租户下的已付费订单，每次购买对应一个唯一的订单，订单会记录购买的套餐的相关信息，业务方需要自行处理套餐的有效期和付费方案的升级。


See: https://open.feishu.cn/document/ukTMukTMukTM/uETNwUjLxUDM14SM1ATN

GET https://open.feishu.cn/open-apis/pay/v1/order/list?status=all&page_size=10&page_token=10&tenant_key=2e5c3a3ae38f175f
*/
func OrderList(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiOrderList+"?"+params.Encode(), header)
}

/*
查询订单详情

该接口用于查询某个订单的具体信息


See: https://open.feishu.cn/document/ukTMukTMukTM/uITNwUjLyUDM14iM1ATN

GET https://open.feishu.cn/open-apis/pay/v1/order/get?order_id=6708978506916697671
*/
func OrderGet(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiOrderGet+"?"+params.Encode(), header)
}
