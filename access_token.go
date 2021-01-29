// Copyright 2021 FastWeGo
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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/faabiosr/cachego"
)

type AccessTokenManager interface {
	GetAccessToken() (accessToken string, err error)
}

type getRefreshRequestFunc func() *http.Request

type DefaultAccessTokenManager struct {
	Id                    string
	GetRefreshRequestFunc getRefreshRequestFunc
	Cache                 cachego.Cache
}

// 防止多个 goroutine 并发刷新冲突
var getAccessTokenLock sync.Mutex

// GetAccessToken
func (m *DefaultAccessTokenManager) GetAccessToken() (accessToken string, err error) {

	cacheKey := m.getCacheKey()
	accessToken, err = m.Cache.Fetch(cacheKey)
	if accessToken != "" {
		return
	}

	getAccessTokenLock.Lock()
	defer getAccessTokenLock.Unlock()

	accessToken, err = m.Cache.Fetch(cacheKey)
	if accessToken != "" {
		return
	}

	req := m.GetRefreshRequestFunc()

	// 添加 serverUrl
	if !strings.HasPrefix(req.URL.String(), "http") {
		parse, _ := url.Parse(ServerUrl)
		req.URL.Host = parse.Host
		req.URL.Scheme = parse.Scheme
	}

	if Logger != nil {
		Logger.Printf("%s %s", req.Method, req.URL.String())
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	defer response.Body.Close()

	var result = struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`

		AppAccessToken    string `json:"app_access_token"`
		TenantAccessToken string `json:"tenant_access_token"`

		Expire int `json:"expire"`
	}{}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		err = fmt.Errorf("unmarshal error %s", string(resp))
		return
	}

	if result.AppAccessToken == "" && result.TenantAccessToken == "" {
		err = fmt.Errorf("%s", string(resp))
		return
	}

	accessToken = result.AppAccessToken
	if result.TenantAccessToken != "" {
		accessToken = result.TenantAccessToken
	}

	err = m.Cache.Save(cacheKey, accessToken, time.Duration(result.Expire)*time.Second)
	if err != nil {
		return
	}

	if Logger != nil {
		Logger.Printf("refreshAccessToken: %s ExpiresIn %d\n", accessToken, result.Expire)
	}

	return
}

// getCacheKey
func (m *DefaultAccessTokenManager) getCacheKey() (key string) {
	return "access_token:" + m.Id
}
