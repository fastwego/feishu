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

package meeting_room_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/capabilities/meeting_room"
)

func ExampleBuildingList() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := meeting_room.BuildingList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleBuildingBatchGet() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := meeting_room.BuildingBatchGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleMeetingRoomList() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := meeting_room.MeetingRoomList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleMeetingRoomBatchGet() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := meeting_room.MeetingRoomBatchGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleMeetingRoomFreeBusyBatchGet() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := meeting_room.MeetingRoomFreeBusyBatchGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleInstanceReply() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := meeting_room.InstanceReply(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBuildingCreate() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := meeting_room.BuildingCreate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBuildingUpdate() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := meeting_room.BuildingUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBuildingDelete() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := meeting_room.BuildingDelete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBuildingBatchGetById() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := meeting_room.BuildingBatchGetById(ctx, params)

	fmt.Println(resp, err)
}

func ExampleMeetingRoomCreate() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := meeting_room.MeetingRoomCreate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMeetingRoomUpdate() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := meeting_room.MeetingRoomUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMeetingRoomDelete() {
	var ctx *feishu.App

	payload := []byte("{}")
	resp, err := meeting_room.MeetingRoomDelete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMeetingRoomBatchGetById() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := meeting_room.MeetingRoomBatchGetById(ctx, params)

	fmt.Println(resp, err)
}

func ExampleCountryList() {
	var ctx *feishu.App

	resp, err := meeting_room.CountryList(ctx)

	fmt.Println(resp, err)
}

func ExampleDistrictList() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := meeting_room.DistrictList(ctx, params)

	fmt.Println(resp, err)
}
