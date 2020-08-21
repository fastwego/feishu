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

// Package authen 身份认证

/*
网页授权流程分为四步：

整体流程
第一步: 网页后端发现用户未登录，请求身份验证；
第二步: 用户登录后，开放平台生成登录预授权码，302跳转至重定向地址；
第三步: 网页后端调用获取登录用户身份校验登录预授权码合法性，获取到用户身份；
第四步: 如需其他用户信息，网页后端可调用获取用户信息（身份验证）。
*/
package authen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/fastwego/feishu"
)

const (
	apiGetRedirectUri     = "/open-apis/authen/v1/index"
	apiGetAccessToken     = "/open-apis/authen/v1/access_token"
	apiRefreshAccessToken = "/open-apis/authen/v1/refresh_access_token"
	apiGetUserInfo        = "/open-apis/authen/v1/user_info"
	apiCode2Session       = "/open-apis/mina/v2/tokenLoginValidate"
)

/*
请求身份验证

应用请求用户身份验证时，需按如下方式构造登录链接，并引导用户跳转至此链接。飞书客户端内用户免登，系统浏览器内用户需完成扫码登录。登录成功后会生成登录预授权码 code，并作为参数重定向到重定向URL。

See: https://open.feishu.cn/document/ukTMukTMukTM/ukzN4UjL5cDO14SO3gTN

GET https://open.feishu.cn/open-apis/authen/v1/index?redirect_uri={REDIRECT_URI}&app_id={APPID}&state={STATE}
*/
func GetRedirectUri(app_id string, redirectUri string, state string) (authorizeUrl string) {
	params := url.Values{}
	params.Add("app_id", app_id)
	params.Add("redirect_uri", redirectUri)
	params.Add("state", state)
	return feishu.FeishuServerUrl + apiGetRedirectUri + "?" + params.Encode()
}

type AccessToken struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		AccessToken      string `json:"access_token"`
		AvatarURL        string `json:"avatar_url"`
		AvatarThumb      string `json:"avatar_thumb"`
		AvatarMiddle     string `json:"avatar_middle"`
		AvatarBig        string `json:"avatar_big"`
		ExpiresIn        int    `json:"expires_in"`
		Name             string `json:"name"`
		EnName           string `json:"en_name"`
		OpenID           string `json:"open_id"`
		TenantKey        string `json:"tenant_key"`
		RefreshExpiresIn int    `json:"refresh_expires_in"`
		RefreshToken     string `json:"refresh_token"`
		TokenType        string `json:"token_type"`
	} `json:"data"`
}

/*
获取登录用户身份

通过此接口获取登录预授权码 code 对应的登录用户身份。

See: https://open.feishu.cn/document/ukTMukTMukTM/uEDO4UjLxgDO14SM4gTN

POST https://open.feishu.cn/open-apis/authen/v1/access_token
*/
func GetAccessToken(app_access_token string, code string) (accessTokenResp AccessToken, err error) {

	params := struct {
		AppAccessToken string `json:"app_access_token"`
		GrantType      string `json:"grant_type"`
		Code           string `json:"code"`
	}{
		AppAccessToken: app_access_token,
		GrantType:      "authorization_code",
		Code:           code,
	}

	data, err := json.Marshal(params)
	if err != nil {
		return
	}

	response, err := http.Post(feishu.FeishuServerUrl+apiGetAccessToken, "application/json", bytes.NewReader(data))
	if err != nil {
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &accessTokenResp)
	if err != nil {
		err = fmt.Errorf("%s", string(body))
		return
	}

	return
}

/*
刷新 access_token

该接口用于在 access_token 过期时用 refresh_token 重新获取 access_token。此时会返回新的 refresh_token，再次刷新 access_token 时需要使用新的 refresh_token。

See: https://open.feishu.cn/document/ukTMukTMukTM/uQDO4UjL0gDO14CN4gTN

POST https://open.feishu.cn/open-apis/authen/v1/refresh_access_token
*/
func RefreshAccessToken(app_access_token string, refresh_token string) (accessTokenResp AccessToken, err error) {
	params := struct {
		AppAccessToken string `json:"app_access_token"`
		GrantType      string `json:"grant_type"`
		RefreshToken   string `json:"refresh_token"`
	}{
		AppAccessToken: app_access_token,
		GrantType:      "refresh_token",
		RefreshToken:   refresh_token,
	}

	data, err := json.Marshal(params)
	if err != nil {
		return
	}

	response, err := http.Post(feishu.FeishuServerUrl+apiRefreshAccessToken, "application/json", bytes.NewReader(data))
	if err != nil {
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &accessTokenResp)
	if err != nil {
		err = fmt.Errorf("%s", string(body))
		return
	}

	return
}

type UserInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Name         string `json:"name"`
		AvatarURL    string `json:"avatar_url"`
		AvatarThumb  string `json:"avatar_thumb"`
		AvatarMiddle string `json:"avatar_middle"`
		AvatarBig    string `json:"avatar_big"`
		Email        string `json:"email"`
		UserID       string `json:"user_id"`
		Mobile       string `json:"mobile"`
		Status       int    `json:"status"`
	} `json:"data"`
}

/*
获取用户信息（身份验证）

此接口仅用于获取登录用户的信息。调用此接口需要在 Header 中带上 user_access_token。

See: https://open.feishu.cn/document/ukTMukTMukTM/uIDO4UjLygDO14iM4gTN

GET https://open.feishu.cn/open-apis/authen/v1/user_info
*/
func GetUserInfo(access_token string) (userInfo UserInfo, err error) {
	header := http.Header{}
	header.Set("Authorization", "Bearer "+access_token)
	header.Set("Content-Type", "application/json")

	request, err := http.NewRequest(http.MethodGet, feishu.FeishuServerUrl+apiGetUserInfo, nil)
	if err != nil {
		return
	}

	request.Header = header

	c := &http.Client{}
	response, err := c.Do(request)
	if err != nil {
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		err = fmt.Errorf("%s", string(body))
		return
	}

	return
}

type Session struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		UID          string `json:"uid"`
		OpenID       string `json:"open_id"`
		UnionID      string `json:"union_id"`
		SessionKey   string `json:"session_key"`
		TenantKey    string `json:"tenant_key"`
		EmployeeID   string `json:"employee_id"`
		TokenType    string `json:"token_type"`
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
	} `json:"data"`
}

/*
code2session

本接口应在后端服务器调用

通过 login 接口获取到登录凭证后，开发者可以通过服务器发送请求的方式获取 session_key 和 openId

See: https://open.feishu.cn/document/uYjL24iN/ukjM04SOyQjL5IDN

POST https://open.feishu.cn/open-apis/mina/v2/tokenLoginValidate
*/
func Code2Session(app_access_token string, code string) (sess Session, err error) {
	header := http.Header{}
	header.Set("Authorization", "Bearer "+app_access_token)
	header.Set("Content-Type", "application/json")

	co := struct {
		Code string `json:"code"`
	}{
		Code: code,
	}

	data, err := json.Marshal(co)
	if err != nil {
		return
	}

	request, err := http.NewRequest(http.MethodPost, feishu.FeishuServerUrl+apiCode2Session, bytes.NewReader(data))
	if err != nil {
		return
	}

	request.Header = header

	c := &http.Client{}
	response, err := c.Do(request)
	if err != nil {
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &sess)
	if err != nil {
		err = fmt.Errorf("%s", string(body))
		return
	}

	return
}
