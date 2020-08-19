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

func ExampleInfoGet() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := department.InfoGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleSimpleList() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := department.SimpleList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleDetailBatchGet() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := department.DetailBatchGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleAdd() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := department.Add(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelete() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := department.Delete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdate() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := department.Update(ctx, payload)

	fmt.Println(resp, err)
}
