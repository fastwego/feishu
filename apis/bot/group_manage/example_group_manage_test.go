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

package group_manage_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/bot/group_manage"
)

func ExampleChatCreate() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := group_manage.ChatCreate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleChatList() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := group_manage.ChatList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleChatInfo() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := group_manage.ChatInfo(ctx, params)

	fmt.Println(resp, err)
}

func ExampleChatUpdate() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := group_manage.ChatUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleChatterAdd() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := group_manage.ChatterAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleChatterDelete() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := group_manage.ChatterDelete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDisband() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := group_manage.Disband(ctx, payload)

	fmt.Println(resp, err)
}
