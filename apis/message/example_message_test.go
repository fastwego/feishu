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

package message_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/message"
)

func ExampleBatchSend() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := message.BatchSend(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSend() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := message.Send(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleReadInfo() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := message.ReadInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleImageGet() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := message.ImageGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleFileGet() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := message.FileGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleAppNotify() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := message.AppNotify(ctx, payload)

	fmt.Println(resp, err)
}
