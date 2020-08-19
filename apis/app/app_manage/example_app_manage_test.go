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

package app_manage_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/app/app_manage"
)

func ExampleIsUserAdmin() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := app_manage.IsUserAdmin(ctx, params)

	fmt.Println(resp, err)
}

func ExampleAdminScopeGet() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := app_manage.AdminScopeGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleAppVisibility() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := app_manage.AppVisibility(ctx, params)

	fmt.Println(resp, err)
}

func ExampleVisibleApps() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := app_manage.VisibleApps(ctx, params)

	fmt.Println(resp, err)
}

func ExampleAppList() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := app_manage.AppList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleUpdateVisibility() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := app_manage.UpdateVisibility(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAppAdminUserList() {
	var ctx *feishu.App

	resp, err := app_manage.AppAdminUserList(ctx)

	fmt.Println(resp, err)
}
