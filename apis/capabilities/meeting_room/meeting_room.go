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

// Package meeting_room 会议室
package meeting_room

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/fastwego/feishu"
)

const (
	apiBuildingList                = "/open-apis/meeting_room/building/list"
	apiBuildingBatchGet            = "/open-apis/meeting_room/building/batch_get"
	apiMeetingRoomList             = "/open-apis/meeting_room/room/list"
	apiMeetingRoomBatchGet         = "/open-apis/meeting_room/room/batch_get"
	apiMeetingRoomFreeBusyBatchGet = "/open-apis/meeting_room/freebusy/batch_get"
	apiInstanceReply               = "/open-apis/meeting_room/instance/reply"
	apiBuildingCreate              = "/open-apis/meeting_room/building/create"
	apiBuildingUpdate              = "/open-apis/meeting_room/building/update"
	apiBuildingDelete              = "/open-apis/meeting_room/building/delete"
	apiBuildingBatchGetById        = "/open-apis/meeting_room/building/batch_get_id"
	apiMeetingRoomCreate           = "/open-apis/meeting_room/room/create"
	apiMeetingRoomUpdate           = "/open-apis/meeting_room/room/update"
	apiMeetingRoomDelete           = "/open-apis/meeting_room/room/delete"
	apiMeetingRoomBatchGetById     = "/open-apis/meeting_room/room/batch_get_id"
	apiCountryList                 = "/open-apis/meeting_room/country/list"
	apiDistrictList                = "/open-apis/meeting_room/district/list"
)

/*
获取建筑物列表


该接口用于获取本企业下的建筑物（办公大楼）。

**权限说明** ：需要获取 获取会议室信息 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ugzNyUjL4cjM14CO3ITN

GET https://open.feishu.cn/open-apis/meeting_room/building/list?page_size=1&page_token=0&order_by=name-asc&fields=*
*/
func BuildingList(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiBuildingList+"?"+params.Encode(), header)
}

/*
查询建筑物详情


该接口用于获取指定建筑物的详细信息。

**权限说明** ：需要获取 获取会议室信息 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukzNyUjL5cjM14SO3ITN

GET https://open.feishu.cn/open-apis/meeting_room/building/batch_get?building_ids=omb_8ec170b937536a5d87c23b418b83f9bb&building_ids=omb_38570e4f0fd9ecf15030d3cc8b388f3a&fields=*
*/
func BuildingBatchGet(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiBuildingBatchGet+"?"+params.Encode(), header)
}

/*
获取会议室列表


该接口用于获取指定建筑下的会议室。

**权限说明** ：需要获取 获取会议室信息 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uADOyUjLwgjM14CM4ITN

GET https://open.feishu.cn/open-apis/meeting_room/room/list?building_id=omb_8ec170b937536a5d87c23b418b83f9bb&page_size=1&page_token=0&order_by=name-asc&fields=*
*/
func MeetingRoomList(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiMeetingRoomList+"?"+params.Encode(), header)
}

/*
查询会议室详情


该接口用于获取指定会议室的详细信息。

**权限说明** ：需要获取 获取会议室信息 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uEDOyUjLxgjM14SM4ITN

GET https://open.feishu.cn/open-apis/meeting_room/room/batch_get?room_ids=omm_eada1d61a550955240c28757e7dec3af&room_ids=omm_83d09ad4f6896e02029a6a075f71c9d1&fields=*
*/
func MeetingRoomBatchGet(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiMeetingRoomBatchGet+"?"+params.Encode(), header)
}

/*
会议室忙闲查询


该接口用于获取指定会议室的忙闲日程实例列表。非重复日程只有唯一实例；重复日程可能存在多个实例，依据重复规则和时间范围扩展。

**权限说明** ：需要获取 获取会议室信息 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uIDOyUjLygjM14iM4ITN

GET https://open.feishu.cn/open-apis/meeting_room/freebusy/batch_get?room_ids=omm_83d09ad4f6896e02029a6a075f71c9d1&room_ids=omm_eada1d61a550955240c28757e7dec3af&time_min=2019-09-04T08:45:00%2B08:00&time_max=2019-09-04T09:45:00%2B08:00
*/
func MeetingRoomFreeBusyBatchGet(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiMeetingRoomFreeBusyBatchGet+"?"+params.Encode(), header)
}

/*
回复会议室日程实例


该接口用于回复会议室日程实例，包括未签到释放和提前结束释放。

**权限说明** ：需要获取 获取会议室信息 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYzN4UjL2cDO14iN3gTN

POST https://open.feishu.cn/open-apis/meeting_room/instance/reply
*/
func InstanceReply(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiInstanceReply, bytes.NewReader(payload), header)
}

/*
创建建筑物


该接口对应管理后台的添加建筑，添加楼层的功能，可用于创建建筑物和建筑物的楼层信息。

**权限说明** ：需要获取 管理会议室信息 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uATNwYjLwUDM24CM1AjN

POST https://open.feishu.cn/open-apis/meeting_room/building/create
*/
func BuildingCreate(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiBuildingCreate, bytes.NewReader(payload), header)
}

/*
更新建筑物


该接口用于编辑建筑信息，添加楼层，删除楼层，编辑楼层信息。

**权限说明** ：需要获取 管理会议室信息 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uETNwYjLxUDM24SM1AjN

POST https://open.feishu.cn/open-apis/meeting_room/building/update
*/
func BuildingUpdate(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiBuildingUpdate, bytes.NewReader(payload), header)
}

/*
删除建筑物


该接口用于删除建筑物（办公大楼）。

**权限说明** ：需要获取 管理会议室信息 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMzMxYjLzMTM24yMzEjN

POST https://open.feishu.cn/open-apis/meeting_room/building/delete
*/
func BuildingDelete(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiBuildingDelete, bytes.NewReader(payload), header)
}

/*
查询建筑物ID


该接口用于根据租户自定义建筑 ID 查询建筑 ID。

**权限说明** ：需要获取 获取会议室信息 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQzMxYjL0MTM24CNzEjN

GET https://open.feishu.cn/open-apis/meeting_room/building/batch_get_id?custom_building_ids=test01&custom_building_ids=test02
*/
func BuildingBatchGetById(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiBuildingBatchGetById+"?"+params.Encode(), header)
}

/*
创建会议室


该接口用于创建会议室。

**权限说明** ：需要获取 管理会议室信息 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uITNwYjLyUDM24iM1AjN

POST https://open.feishu.cn/open-apis/meeting_room/room/create
*/
func MeetingRoomCreate(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiMeetingRoomCreate, bytes.NewReader(payload), header)
}

/*
更新会议室


该接口用于更新会议室。

**权限说明** ：需要获取 管理会议室信息 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMTNwYjLzUDM24yM1AjN

POST https://open.feishu.cn/open-apis/meeting_room/room/update
*/
func MeetingRoomUpdate(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiMeetingRoomUpdate, bytes.NewReader(payload), header)
}

/*
删除会议室


该接口用于删除会议室。

**权限说明** ：需要获取 管理会议室信息 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uUzMxYjL1MTM24SNzEjN

POST https://open.feishu.cn/open-apis/meeting_room/room/delete
*/
func MeetingRoomDelete(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiMeetingRoomDelete, bytes.NewReader(payload), header)
}

/*
查询会议室ID


该接口用于根据租户自定义会议室ID查询会议室ID。

**权限说明** ：需要获取 获取会议室信息 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYzMxYjL2MTM24iNzEjN

GET https://open.feishu.cn/open-apis/meeting_room/room/batch_get_id?custom_room_ids=test01&custom_room_ids=test02
*/
func MeetingRoomBatchGetById(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiMeetingRoomBatchGetById+"?"+params.Encode(), header)
}

/*
获取国家地区列表


新建建筑时需要标明所处国家/地区，该接口用于获得系统预先提供的可供选择的国家 /地区列表

**权限说明** ：需要获取 获取会议室信息 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQTNwYjL0UDM24CN1AjN

GET https://open.feishu.cn/open-apis/meeting_room/country/list
*/
func CountryList(ctx *feishu.App) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiCountryList, header)
}

/*
获取城市列表


新建建筑时需要选择所处国家/地区，该接口用于获得系统预先提供的可供选择的城市列表

**权限说明** ：需要获取 获取会议室信息 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uUTNwYjL1UDM24SN1AjN

GET https://open.feishu.cn/open-apis/meeting_room/district/list?country_id=1814991
*/
func DistrictList(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiDistrictList+"?"+params.Encode(), header)
}
