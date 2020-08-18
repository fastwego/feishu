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

// Package message 消息
package message

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/fastwego/feishu"
)

const (
	apiBatchSend = "/open-apis/message/v4/batch_send/"
	apiSend      = "/open-apis/message/v4/send/"
	apiReadInfo  = "/open-apis/message/v4/read_info/"
	apiImagePut  = "/open-apis/image/v4/put/"
	apiImageGet  = "/open-apis/image/v4/get"
	apiFileGet   = "/open-apis/open-file/v1/get"
	apiAppNotify = "/open-apis/notify/v4/appnotify"
)

/*
批量发送消息


给多个用户或者多个部门发送消息。

**权限说明** ：需要启用机器人能力；机器人需要拥有批量发送消息权限；机器人需要拥有对用户或部门的可见性



See: https://open.feishu.cn/document/ukTMukTMukTM/ucDO1EjL3gTNx4yN4UTM

POST https://open.feishu.cn/open-apis/message/v4/batch_send/
*/
func BatchSend(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiBatchSend, bytes.NewReader(payload), header)
}

/*
发送文本消息


给指定用户或者会话发送文本消息，其中会话包括私聊会话和群会话。

**权限说明** ：需要启用机器人能力；私聊会话时机器人需要拥有对用户的可见性，群会话需要机器人在群里


See: https://open.feishu.cn/document/ukTMukTMukTM/uUjNz4SN2MjL1YzM

POST https://open.feishu.cn/open-apis/message/v4/send/
*/
func Send(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiSend, bytes.NewReader(payload), header)
}

/*
查询消息已读状态


查询消息已读状态，只能查询最近七天机器人自身发送消息的已读信息。

**权限说明** ：需要启用机器人能力


See: https://open.feishu.cn/document/ukTMukTMukTM/ukTM2UjL5EjN14SOxYTN

POST https://open.feishu.cn/open-apis/message/v4/read_info/
*/
func ReadInfo(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiReadInfo, bytes.NewReader(payload), header)
}

/*
上传图片


上传图片，获取图片的 image_key。

**权限说明** ：需要启用机器人能力


See: https://open.feishu.cn/document/ukTMukTMukTM/uEDO04SM4QjLxgDN

POST https://open.feishu.cn/open-apis/image/v4/put/
*/
func ImagePut(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiImagePut, bytes.NewReader(payload), header)
}

/*
获取图片


根据图片的image_key获取图片内容

**权限说明** ：需要启用机器人能力；机器人只能获取自己上传或者收到推送的图片


See: https://open.feishu.cn/document/ukTMukTMukTM/uYzN5QjL2cTO04iN3kDN

GET https://open.feishu.cn/open-apis/image/v4/get?image_key=24383920-9321-4ecd-8b33-bf8ce74e84c8
*/
func ImageGet(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiImageGet+"?"+params.Encode(), header)
}

/*
获取文件


根据文件的 file_key 拉取文件内容，当前仅可用来获取用户与机器人单聊发送的文件

**权限说明** ：需要启用机器人能力；机器人只能获取自己上传的文件


See: https://open.feishu.cn/document/ukTMukTMukTM/uMDN4UjLzQDO14yM0gTN

GET https://open.feishu.cn/open-apis/open-file/v1/get?file_key=file_36r377cb-c6h2-4b6d-ag67-0ac3e796008g
*/
func FileGet(ctx *feishu.App, params url.Values) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiFileGet+"?"+params.Encode(), header)
}

/*
应用发送通知给用户

给指定用户发送应用通知

**权限说明** ：需要申请  以应用的身份发消息 的权限；同时需要启用小程序能力，并且支持**PC模式**。


See: https://open.feishu.cn/document/ukTMukTMukTM/uATOyUjLwkjM14CM5ITN

POST https://open.feishu.cn/open-apis/notify/v4/appnotify
*/
func AppNotify(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiAppNotify, bytes.NewReader(payload), header)
}
