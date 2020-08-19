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

package bot_manage_test

import (
	"fmt"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/bot/bot_manage"
)

func ExampleInfo() {
	var ctx *feishu.App

	resp, err := bot_manage.Info(ctx)

	fmt.Println(resp, err)
}

func ExampleAdd() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := bot_manage.Add(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleRemove() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := bot_manage.Remove(ctx, payload)

	fmt.Println(resp, err)
}
