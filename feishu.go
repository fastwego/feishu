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

/*
飞书 开放平台 SDK

See: https://open.feishu.cn/
*/
package feishu

import (
	"log"
	"os"

	"github.com/faabiosr/cachego"
	"github.com/faabiosr/cachego/file"
)

// GetTenantAccessTokenFunc 获取 tenant_access_token 方法接口
type GetTenantAccessTokenFunc func() (accessToken string, err error)
type GetAppAccessTokenFunc func() (accessToken string, err error)

/*
App 实例
*/
type App struct {
	Config                      AppConfig
	Cache                       cachego.Cache
	GetTenantAccessTokenHandler GetTenantAccessTokenFunc
	GetAppAccessTokenHandler    GetAppAccessTokenFunc
	Client                      Client
	Server                      Server
	Logger                      *log.Logger
}

/*
应用 配置
*/
type AppConfig struct {
	AppId             string
	AppSecret         string
	VerificationToken string
	EncryptKey        string
}

func newApp(config AppConfig) (app *App) {
	instance := App{
		Config: config,
		Cache:  file.New(os.TempDir()),
	}

	instance.Client = Client{Ctx: &instance}
	instance.Server = Server{Ctx: &instance}

	instance.Logger = log.New(os.Stdout, "[feishu] ", log.LstdFlags|log.Llongfile)

	return &instance
}

/*
SetTenantAccessTokenCacheDriver 设置 TenantAccessToken 缓存器 默认为文件缓存：目录 os.TempDir()

驱动接口类型 为 cachego.Cache
*/
func (app *App) SetCacheDriver(driver cachego.Cache) {
	app.Cache = driver
}

/*
SetGetTenantAccessTokenHandler 设置 TenantAccessToken 获取方法。默认 从本地缓存获取（过期从微信接口刷新）

如果有多实例服务，可以设置为 Redis 或 RPC 等中控服务器 获取 就可以避免 TenantAccessToken 刷新冲突
*/
func (app *App) SetGetTenantAccessTokenHandler(f GetTenantAccessTokenFunc) {
	app.GetTenantAccessTokenHandler = f
}

func (app *App) SetGetAppAccessTokenHandler(f GetAppAccessTokenFunc) {
	app.GetAppAccessTokenHandler = f
}

/*
SetLogger 日志记录 默认输出到 os.Stdout

可以新建 logger 输出到指定文件

如果不想开启日志，可以 SetLogger(nil)
*/
func (app *App) SetLogger(logger *log.Logger) {
	app.Logger = logger
}
