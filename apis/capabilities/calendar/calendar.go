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

// Package calendar 日历
package calendar

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"

	"github.com/fastwego/feishu"
)

const (
	apiGetCalendarById      = "/open-apis/calendar/v3/calendar_list/:calendarId"
	apiCalendarList         = "/open-apis/calendar/v3/calendar_list"
	apiCreateCalendars      = "/open-apis/calendar/v3/calendars"
	apiDeleteCalendarById   = "/open-apis/calendar/v3/calendars/:calendarId"
	apiGetEventById         = "/open-apis/calendar/v3/calendars/:calendarId/events/:eventId"
	apiCreateEvent          = "/open-apis/calendar/v3/calendars/:calendarId/events"
	apiAttendees            = "/open-apis/calendar/v3/calendars/:calendarId/events/:eventId/attendees"
	apiAcl                  = "/open-apis/calendar/v3/calendars/:calendarId/acl"
	apiDeleteAclByRuleId    = "/open-apis/calendar/v3/calendars/:calendarId/acl/:ruleId"
	apiFreeBusyQuery        = "/open-apis/calendar/v3/freebusy/query"
	apiSharedCalendarQuery  = "/open-apis/calendar/v3/shared_calendar_list/shared_calendar/query"
	apiSharedCalendarEvents = "/open-apis/calendar/v3/shared/calendars/:calendarId/events"
)

/*
获取日历


该接口用于根据日历 ID 获取日历信息。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMDN04yM0QjLzQDN

GET https://open.feishu.cn/open-apis/calendar/v3/calendar_list/:calendarId
*/
func GetCalendarById(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	api := apiGetCalendarById
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(api, header)
}

/*
获取日历列表


该接口用于获取应用在企业内的日历列表。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMTM14yMxUjLzETN

GET https://open.feishu.cn/open-apis/calendar/v3/calendar_list
*/
func CalendarList(ctx *feishu.App) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiCalendarList, header)
}

/*
创建日历


该接口用于为应用在企业内创建一个日历。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQTM14CNxUjL0ETN

POST https://open.feishu.cn/open-apis/calendar/v3/calendars
*/
func CreateCalendars(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiCreateCalendars, bytes.NewReader(payload), header)
}

/*
删除日历


该接口用于删除应用在企业内的指定日历。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uUTM14SNxUjL1ETN

DELETE https://open.feishu.cn/open-apis/calendar/v3/calendars/:calendarId
*/
func DeleteCalendarById(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	api := apiDeleteCalendarById
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPDelete(api, header)
}

/*
获取日程


该接口用于获取指定日历下的指定日程。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ucTM14yNxUjL3ETN

GET https://open.feishu.cn/open-apis/calendar/v3/calendars/:calendarId/events/:eventId
*/
func GetEventById(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	api := apiGetEventById
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(api, header)
}

/*
创建日程


该接口用于在日历中创建一个日程。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ugTM14COxUjL4ETN

POST https://open.feishu.cn/open-apis/calendar/v3/calendars/:calendarId/events
*/
func CreateEvent(ctx *feishu.App, payload []byte, params url.Values) (resp []byte, err error) {

	api := apiCreateEvent
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(api, bytes.NewReader(payload), header)
}

/*
邀请/移除日程参与者


邀请一个或多个用户加入日程；
从日程移除一个或多个用户。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uIjM14iMyUjLyITN

POST https://open.feishu.cn/open-apis/calendar/v3/calendars/:calendarId/events/:eventId/attendees
*/
func Attendees(ctx *feishu.App, payload []byte, params url.Values) (resp []byte, err error) {

	api := apiAttendees
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(api, bytes.NewReader(payload), header)
}

/*
获取访问控制列表


该接口用于查看指定日历的成员列表。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMjM14yMyUjLzITN

GET https://open.feishu.cn/open-apis/calendar/v3/calendars/:calendarId/acl
*/
func Acl(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	api := apiAcl
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(api, header)
}

/*
删除访问控制


该接口用于从日历中移除一个用户。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uUjM14SNyUjL1ITN

DELETE https://open.feishu.cn/open-apis/calendar/v3/calendars/:calendarId/acl/:ruleId
*/
func DeleteAclByRuleId(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	api := apiDeleteAclByRuleId
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPDelete(api, header)
}

/*
查询日历的忙闲状态


该接口用于查询指定日历或用户主日历的忙闲状态。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYjM14iNyUjL2ITN

POST https://open.feishu.cn/open-apis/calendar/v3/freebusy/query
*/
func FreeBusyQuery(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiFreeBusyQuery, bytes.NewReader(payload), header)
}

/*
查询公共日历


该接口用于通过关键字查询公共日历信息。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukDMwYjL5ADM24SOwAjN

GET https://open.feishu.cn/open-apis/calendar/v3/shared_calendar_list/shared_calendar/query?query=ByteDance
*/
func SharedCalendarQuery(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiSharedCalendarQuery+"?"+params.Encode(), header)
}

/*
获取公共日历日程列表


该接口用于获取公共日历下的日程列表。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uIzNwYjLycDM24iM3AjN

GET https://open.feishu.cn/open-apis/calendar/v3/shared/calendars/:calendarId/events
*/
func SharedCalendarEvents(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	api := apiSharedCalendarEvents
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(api, header)
}
