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

package approval_test

import (
	"fmt"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/capabilities/approval"
)

func ExampleGet() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := approval.Get(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleInstanceList() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := approval.InstanceList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleInstanceGet() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := approval.InstanceGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleInstanceCreate() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := approval.InstanceCreate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleApprove() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := approval.Approve(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleReject() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := approval.Reject(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTransfer() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := approval.Transfer(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCancel() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := approval.Cancel(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleExternalInstanceCreate() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := approval.ExternalInstanceCreate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleExternalInstanceCheck() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := approval.ExternalInstanceCheck(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMessageSend() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := approval.MessageSend(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMessageUpdate() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := approval.MessageUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleApprovalCreate() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := approval.ApprovalCreate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleInstanceCc() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := approval.InstanceCc(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSubscribe() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := approval.Subscribe(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUnsubscribe() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := approval.Unsubscribe(ctx, payload)

	fmt.Println(resp, err)
}
