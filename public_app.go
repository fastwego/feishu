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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type GetAppTicketFunc func() (appTicket string, err error)
type ReceiveAppTicketFunc func(appTicket string) (err error)

type PublicApp struct {
	*App
	TenantKey               string
	GetAppTicketHandler     GetAppTicketFunc
	ReceiveAppTicketHandler ReceiveAppTicketFunc
}

// 应用市场应用
func NewPublicApp(config AppConfig, tenantKey string) (publicApp *PublicApp) {
	publicApp = &PublicApp{
		App: newApp(config),
	}
	publicApp.TenantKey = tenantKey
	publicApp.App.GetTenantAccessTokenHandler = publicApp.GetTenantAccessToken
	publicApp.App.GetAppAccessTokenHandler = publicApp.GetAppAccessToken

	publicApp.GetAppTicketHandler = publicApp.GetAppTicket
	publicApp.ReceiveAppTicketHandler = publicApp.ReceiveAppTicket
	return
}

// GetAppTicket 获取 AppTicket
func (ctx *PublicApp) GetAppTicket() (appTicket string, err error) {
	appTicket, err = ctx.Cache.Fetch("app_ticket:" + ctx.Config.AppId)
	if appTicket != "" {
		return
	}

	// 请求推送
	ctx.RequestAppTicket()
	return
}

// RequestAppTicket 请求重新发送 app_ticket
func (ctx *PublicApp) RequestAppTicket() {

	params := struct {
		AppID     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
	}{
		AppID:     ctx.Config.AppId,
		AppSecret: ctx.Config.AppSecret,
	}

	data, err := json.Marshal(params)
	if err != nil {
		return
	}

	url := FeishuServerUrl + "/open-apis/auth/v3/app_ticket/resend/"

	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if ctx.Logger != nil {
		ctx.Logger.Println(url, string(data), string(body))
	}

}

// ReceiveAppTicket 接收 app_ticket 并 存储到 缓存
func (ctx *PublicApp) ReceiveAppTicket(ticket string) (err error) {
	return ctx.Cache.Save("app_ticket:"+ctx.Config.AppId, ticket, 0)
}

// 防止多个 goroutine 并发刷新冲突
var getPublicAppAccessTokenLock sync.Mutex

/*
从 应用 实例 的 AppAccessToken 管理器 获取 access_token

如果没有 access_token 或者 已过期，那么刷新

获得新的 access_token 后 过期时间设置为 0.9 * expiresIn 提供一定冗余
*/
func (ctx *PublicApp) GetAppAccessToken() (accessToken string, err error) {
	cacheKey := "app_access_token:" + ctx.Config.AppId
	accessToken, err = ctx.Cache.Fetch(cacheKey)
	if accessToken != "" {
		return
	}

	getPublicAppAccessTokenLock.Lock()
	defer getPublicAppAccessTokenLock.Unlock()

	accessToken, err = ctx.Cache.Fetch(cacheKey)
	if accessToken != "" {
		return
	}

	accessToken, expiresIn, err := ctx.refreshAppAccessToken()
	if err != nil {
		return
	}

	// 提前过期 提供冗余时间
	d := time.Duration(expiresIn-300) * time.Second
	_ = ctx.Cache.Save(cacheKey, accessToken, d)

	if ctx.Logger != nil {
		ctx.Logger.Printf("%s %s %d\n", "refreshAppAccessTokenFromServer", accessToken, expiresIn)
	}

	return
}

/*
从服务器获取新的 AppAccessToken

See: https://open.feishu.cn/document/ukTMukTMukTM/uEjNz4SM2MjLxYzM
*/
func (ctx *PublicApp) refreshAppAccessToken() (accessToken string, expiresIn int, err error) {

	ticket, err := ctx.GetAppTicketHandler()
	if err != nil {
		return
	}

	params := struct {
		AppID     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
		AppTicket string `json:"app_ticket"`
	}{
		AppID:     ctx.Config.AppId,
		AppSecret: ctx.Config.AppSecret,
		AppTicket: ticket,
	}

	url := FeishuServerUrl + "/open-apis/auth/v3/app_access_token/"

	data, err := json.Marshal(params)
	if err != nil {
		return
	}

	response, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("POST %s RETURN %s", url, response.Status)
		return
	}

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var result = struct {
		Code           int    `json:"code"`
		Msg            string `json:"msg"`
		AppAccessToken string `json:"app_access_token"`
		Expire         int    `json:"expire"`
	}{}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		err = fmt.Errorf("Unmarshal error %s", string(resp))
		return
	}

	if result.AppAccessToken == "" {
		err = fmt.Errorf("%s", string(resp))
		return
	}

	return result.AppAccessToken, result.Expire, nil
}

// 防止多个 goroutine 并发刷新冲突
var getPublicTenantAccessTokenLock sync.Mutex

/*
从 应用 实例 的 TenantAccessToken 管理器 获取 access_token

如果没有 access_token 或者 已过期，那么刷新

*/
func (ctx *PublicApp) GetTenantAccessToken() (accessToken string, err error) {
	cacheKey := "tenant_access_token:" + ctx.TenantKey + ":" + ctx.Config.AppId

	accessToken, err = ctx.Cache.Fetch(cacheKey)
	if accessToken != "" {
		return
	}
	getPublicTenantAccessTokenLock.Lock()
	defer getPublicTenantAccessTokenLock.Unlock()

	accessToken, err = ctx.Cache.Fetch(cacheKey)
	if accessToken != "" {
		return
	}

	var expiresIn int
	accessToken, expiresIn, err = ctx.refreshTenantAccessToken()

	if err != nil {
		return
	}

	// 提前过期 提供冗余时间
	d := time.Duration(expiresIn-300) * time.Second
	_ = ctx.Cache.Save(cacheKey, accessToken, d)

	if ctx.Logger != nil {
		ctx.Logger.Printf("%s %s %d\n", "refreshTenantAccessTokenFromServer", accessToken, expiresIn)
	}

	return
}

/*
See: https://open.feishu.cn/document/ukTMukTMukTM/uMjNz4yM2MjLzYzM
*/
func (ctx *PublicApp) refreshTenantAccessToken() (accessToken string, expiresIn int, err error) {

	appAccessToken, err := ctx.GetAppAccessTokenHandler()
	if err != nil {
		return
	}
	params := struct {
		AppAccessToken string `json:"app_access_token"`
		TenantKey      string `json:"tenant_key"`
	}{
		AppAccessToken: appAccessToken,
		TenantKey:      ctx.TenantKey,
	}
	data, err := json.Marshal(params)
	if err != nil {
		return
	}

	url := FeishuServerUrl + "/open-apis/auth/v3/tenant_access_token/"

	response, err := http.Post(url, "application/json;charset=utf-8", bytes.NewReader(data))
	if err != nil {
		return
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("POST %s RETURN %s", url, response.Status)
		return
	}

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var result = struct {
		Code              int    `json:"code"`
		Msg               string `json:"msg"`
		TenantAccessToken string `json:"tenant_access_token"`
		Expire            int    `json:"expire"`
	}{}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		err = fmt.Errorf("Unmarshal error %s", string(resp))
		return
	}

	if result.TenantAccessToken == "" {
		err = fmt.Errorf("%s", string(resp))
		return
	}

	return result.TenantAccessToken, result.Expire, nil
}
