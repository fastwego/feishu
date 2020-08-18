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

// Package appstore 应用信息/应用商店
package appstore

import (
	"net/http"
	"net/url"

	"github.com/fastwego/feishu"
)

const (
	apiOrderList = "/open-apis/pay/v1/order/list"
	apiOrderGet  = "/open-apis/pay/v1/order/get"
)

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
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiOrderList+"?"+params.Encode(), header)
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
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiOrderGet+"?"+params.Encode(), header)
}
