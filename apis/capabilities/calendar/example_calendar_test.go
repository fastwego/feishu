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

package calendar_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/capabilities/calendar"
)

func ExampleGetCalendarById() {
	var ctx *feishu.App

	params := url.Values{}

	resp, err := calendar.GetCalendarById(ctx, params)

	fmt.Println(resp, err)
}

func ExampleCalendarList() {
	var ctx *feishu.App

	resp, err := calendar.CalendarList(ctx)

	fmt.Println(resp, err)
}

func ExampleCreateCalendars() {
	var ctx *feishu.App

	payload := []byte("{}")

	resp, err := calendar.CreateCalendars(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDeleteCalendarById() {
	var ctx *feishu.App

	params := url.Values{}

	resp, err := calendar.DeleteCalendarById(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGetEventById() {
	var ctx *feishu.App

	params := url.Values{}

	resp, err := calendar.GetEventById(ctx, params)

	fmt.Println(resp, err)
}

func ExampleCreateEvent() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}

	resp, err := calendar.CreateEvent(ctx, payload, params)

	fmt.Println(resp, err)
}

func ExampleAttendees() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}

	resp, err := calendar.Attendees(ctx, payload, params)

	fmt.Println(resp, err)
}

func ExampleAcl() {
	var ctx *feishu.App

	params := url.Values{}

	resp, err := calendar.Acl(ctx, params)

	fmt.Println(resp, err)
}

func ExampleDeleteAclByRuleId() {
	var ctx *feishu.App

	params := url.Values{}

	resp, err := calendar.DeleteAclByRuleId(ctx, params)

	fmt.Println(resp, err)
}

func ExampleFreeBusyQuery() {
	var ctx *feishu.App

	payload := []byte("{}")

	resp, err := calendar.FreeBusyQuery(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSharedCalendarQuery() {
	var ctx *feishu.App

	params := url.Values{}

	resp, err := calendar.SharedCalendarQuery(ctx, params)

	fmt.Println(resp, err)
}

func ExampleSharedCalendarEvents() {
	var ctx *feishu.App

	params := url.Values{}

	resp, err := calendar.SharedCalendarEvents(ctx, params)

	fmt.Println(resp, err)
}
