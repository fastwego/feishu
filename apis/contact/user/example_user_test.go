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

package user_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/contact/user"
)

func ExampleBatchGet() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := user.BatchGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleDepartmentUserList() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := user.DepartmentUserList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleDepartmentUserDetailList() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := user.DepartmentUserDetailList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleAdd() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := user.Add(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelete() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := user.Delete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdate() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := user.Update(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleRoleList() {
	var ctx *feishu.App

	resp, err := user.RoleList(ctx)

	fmt.Println(resp, err)
}

func ExampleRoleMembers() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := user.RoleMembers(ctx, params)

	fmt.Println(resp, err)
}

func ExampleBatchGetId() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := user.BatchGetId(ctx, params)

	fmt.Println(resp, err)
}

func ExampleUnionIdBatchGetList() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := user.UnionIdBatchGetList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleSearch() {
	var ctx *feishu.App

	params := url.Values{}
	accessToken := ""
	resp, err := user.Search(ctx, params, accessToken)

	fmt.Println(resp, err)
}
