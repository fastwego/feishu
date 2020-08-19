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

package app_store_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/app/app_store"
)

func ExampleCheckUser() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := app_store.CheckUser(ctx, params)

	fmt.Println(resp, err)
}

func ExampleOrderList() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := app_store.OrderList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleOrderGet() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := app_store.OrderGet(ctx, params)

	fmt.Println(resp, err)
}
