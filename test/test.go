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

// Package test 模拟服务器 测试
package test

import (
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/gorilla/mux"

	"github.com/fastwego/feishu"
)

var MockApp *feishu.App
var MockRouter *mux.Router
var MockSvr *httptest.Server
var onceSetup sync.Once

// 初始化测试环境
func Setup() {
	onceSetup.Do(func() {

		MockApp = feishu.NewApp(feishu.AppConfig{
			AppId:     "APPID",
			AppSecret: "SECRET",
			//VerificationToken: "TOKEN",
			//EncryptKey:        "EncryptKey",
		})

		// Mock Server
		MockRouter = mux.NewRouter()
		MockSvr = httptest.NewServer(MockRouter)
		feishu.FeishuServerUrl = MockSvr.URL // 拦截发往服务器的请求
		feishu.FeishuServerUrl2 = MockSvr.URL

		// Mock access token
		MockRouter.HandleFunc("/open-apis/auth/v3/tenant_access_token/internal/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"tenant_access_token":"ACCESS_TOKEN","expire":7200}`))
		}).Methods("POST")
		MockRouter.HandleFunc("/open-apis/auth/v3/app_access_token/internal/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"app_access_token":"ACCESS_TOKEN","expire":7200}`))
		}).Methods("POST")
	})
}
