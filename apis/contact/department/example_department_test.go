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

package department_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/contact/department"
)

func ExampleDepartmentInfoGet() {
	var ctx *feishu.App

	params := url.Values{}

	resp, err := department.DepartmentInfoGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleDepartmentSimpleList() {
	var ctx *feishu.App

	params := url.Values{}

	resp, err := department.DepartmentSimpleList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleDepartmentDetailBatchGet() {
	var ctx *feishu.App

	params := url.Values{}

	resp, err := department.DepartmentDetailBatchGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleDepartmentAdd() {
	var ctx *feishu.App

	payload := []byte("{}")

	resp, err := department.DepartmentAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDepartmentDelete() {
	var ctx *feishu.App

	payload := []byte("{}")

	resp, err := department.DepartmentDelete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDepartmentUpdate() {
	var ctx *feishu.App

	payload := []byte("{}")

	resp, err := department.DepartmentUpdate(ctx, payload)

	fmt.Println(resp, err)
}
