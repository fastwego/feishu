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

/*
创建 企业自建应用 实例
*/
func NewApp(config AppConfig) (app *App) {
	app = newApp(config)
	app.GetTenantAccessTokenHandler = app.GetTenantAccessToken
	app.GetAppAccessTokenHandler = app.GetAppAccessToken
	return
}

// 防止多个 goroutine 并发刷新冲突
var getInternalTenantAccessTokenLock sync.Mutex

/*
从 应用 实例 的 TenantAccessToken 管理器 获取 access_token

如果没有 access_token 或者 已过期，那么刷新

*/
func (app *App) GetTenantAccessToken() (accessToken string, err error) {

	cacheKey := "tenant_access_token:" + app.Config.AppId
	accessToken, err = app.Cache.Fetch(cacheKey)
	if accessToken != "" {
		return
	}

	getInternalTenantAccessTokenLock.Lock()
	defer getInternalTenantAccessTokenLock.Unlock()

	accessToken, err = app.Cache.Fetch(cacheKey)
	if accessToken != "" {
		return
	}

	var expiresIn int

	accessToken, expiresIn, err = app.refreshTenantAccessToken()
	if err != nil {
		return
	}

	// 提前 5min 过期 提供冗余时间
	//Token 有效期为 2 小时，在此期间调用该接口 token 不会改变。当 token 有效期小于 10 分的时候，再次请求获取 token 的时候，会生成一个新的 token，与此同时老的 token 依然有效。
	d := time.Duration(expiresIn-300) * time.Second
	_ = app.Cache.Save(cacheKey, accessToken, d)

	if app.Logger != nil {
		app.Logger.Printf("%s %s %d\n", "refreshTenantAccessTokenFromServer", accessToken, expiresIn)
	}

	return
}

/*
从服务器获取新的 TenantAccessToken

See: https://open.feishu.cn/document/ukTMukTMukTM/uIjNz4iM2MjLyYzM
*/
func (app *App) refreshTenantAccessToken() (accessToken string, expiresIn int, err error) {

	params := struct {
		AppID     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
	}{
		AppID:     app.Config.AppId,
		AppSecret: app.Config.AppSecret,
	}
	data, err := json.Marshal(params)
	if err != nil {
		return
	}

	url := FeishuServerUrl + "/open-apis/auth/v3/tenant_access_token/internal/"

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

// 防止多个 goroutine 并发刷新冲突
var getAppAccessTokenLock sync.Mutex

/*
从 应用 实例 的 AppAccessToken 管理器 获取 access_token

如果没有 access_token 或者 已过期，那么刷新
*/
func (app *App) GetAppAccessToken() (accessToken string, err error) {
	cacheKey := "app_access_token:" + app.Config.AppId
	accessToken, err = app.Cache.Fetch(cacheKey)
	if accessToken != "" {
		return
	}

	getAppAccessTokenLock.Lock()
	defer getAppAccessTokenLock.Unlock()

	accessToken, err = app.Cache.Fetch(cacheKey)
	if accessToken != "" {
		return
	}

	accessToken, expiresIn, err := app.refreshAppAccessToken()
	if err != nil {
		return
	}

	// 提前过期 提供冗余时间
	d := time.Duration(expiresIn-300) * time.Second
	_ = app.Cache.Save(cacheKey, accessToken, d)

	if app.Logger != nil {
		app.Logger.Printf("%s %s %d\n", "refreshAppAccessTokenFromServer", accessToken, expiresIn)
	}

	return
}

/*
从服务器获取新的 AppAccessToken

See: https://open.feishu.cn/document/ukTMukTMukTM/uADN14CM0UjLwQTN
*/
func (app *App) refreshAppAccessToken() (accessToken string, expiresIn int, err error) {

	params := struct {
		AppID     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
	}{
		AppID:     app.Config.AppId,
		AppSecret: app.Config.AppSecret,
	}
	url := FeishuServerUrl + "/open-apis/auth/v3/app_access_token/internal/"
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
