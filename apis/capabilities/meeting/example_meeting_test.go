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

package meeting_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/capabilities/meeting"
)

func ExampleBuildingList() {
	var ctx *feishu.App

	params := url.Values{}

	resp, err := meeting.BuildingList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleBuildingBatchGet() {
	var ctx *feishu.App

	params := url.Values{}

	resp, err := meeting.BuildingBatchGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleMeetingRoomList() {
	var ctx *feishu.App

	params := url.Values{}

	resp, err := meeting.MeetingRoomList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleMeetingRoomBatchGet() {
	var ctx *feishu.App

	params := url.Values{}

	resp, err := meeting.MeetingRoomBatchGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleMeetingRoomFreeBusyBatchGet() {
	var ctx *feishu.App

	params := url.Values{}

	resp, err := meeting.MeetingRoomFreeBusyBatchGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleInstanceReply() {
	var ctx *feishu.App

	payload := []byte("{}")

	resp, err := meeting.InstanceReply(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBuildingCreate() {
	var ctx *feishu.App

	payload := []byte("{}")

	resp, err := meeting.BuildingCreate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBuildingUpdate() {
	var ctx *feishu.App

	payload := []byte("{}")

	resp, err := meeting.BuildingUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBuildingDelete() {
	var ctx *feishu.App

	payload := []byte("{}")

	resp, err := meeting.BuildingDelete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBuildingBatchGetById() {
	var ctx *feishu.App

	params := url.Values{}

	resp, err := meeting.BuildingBatchGetById(ctx, params)

	fmt.Println(resp, err)
}

func ExampleMeetingRoomCreate() {
	var ctx *feishu.App

	payload := []byte("{}")

	resp, err := meeting.MeetingRoomCreate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMeetingRoomUpdate() {
	var ctx *feishu.App

	payload := []byte("{}")

	resp, err := meeting.MeetingRoomUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMeetingRoomDelete() {
	var ctx *feishu.App

	payload := []byte("{}")

	resp, err := meeting.MeetingRoomDelete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMeetingRoomBatchGetById() {
	var ctx *feishu.App

	params := url.Values{}

	resp, err := meeting.MeetingRoomBatchGetById(ctx, params)

	fmt.Println(resp, err)
}

func ExampleCountryList() {
	var ctx *feishu.App

	resp, err := meeting.CountryList(ctx)

	fmt.Println(resp, err)
}

func ExampleDistrictList() {
	var ctx *feishu.App

	params := url.Values{}

	resp, err := meeting.DistrictList(ctx, params)

	fmt.Println(resp, err)
}
