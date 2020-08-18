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

// Package approve 审批
package approve

import (
	"bytes"
	"net/http"

	"github.com/fastwego/feishu"
)

const (
	apiGet                    = "/approval/openapi/v2/approval/get"
	apiInstanceList           = "/approval/openapi/v2/instance/list"
	apiInstanceGet            = "/approval/openapi/v2/instance/get"
	apiInstanceCreate         = "/approval/openapi/v2/instance/create"
	apiApprove                = "/approval/openapi/v2/instance/approve"
	apiReject                 = "/approval/openapi/v2/instance/reject"
	apiTransfer               = "/approval/openapi/v2/instance/transfer"
	apiCancel                 = "/approval/openapi/v2/instance/cancel"
	apiUpload                 = "/approval/openapi/v2/file/upload"
	apiExternalInstanceCreate = "/approval/openapi/v2/external/instance/create"
	apiCreate                 = "/approval/openapi/v2/approval/create"
	apiInstanceCc             = "/approval/openapi/v2/instance/cc"
	apiSubscribe              = "/approval/openapi/v2/subscription/subscribe"
	apiUnsubscribe            = "/approval/openapi/v2/subscription/unsubscribe"
)

/*
查看审批定义


根据 Approval Code 获取某个审批定义的详情，用于构造创建审批实例的请求。

**权限说明** ：需要获取 访问审批应用 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uADNyUjLwQjM14CM0ITN

POST https://www.feishu.cn/approval/openapi/v2/approval/get
*/
func Get(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiGet, bytes.NewReader(payload), header)
}

/*
批量获取审批实例ID


根据 approval_code 批量获取审批实例的 instance_code，用于拉取租户下某个审批定义的全部审批实例。
默认以审批创建时间排序。

**权限说明** ：需要获取 访问审批应用 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQDOyUjL0gjM14CN4ITN

POST https://www.feishu.cn/approval/openapi/v2/instance/list
*/
func InstanceList(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiInstanceList, bytes.NewReader(payload), header)
}

/*
获取单个审批实例详情


通过审批实例 Instance Code  获取审批实例详情。Instance Code 由 [批量获取审批实例](https://open.feishu.cn/document/ukTMukTMukTM/uQDOyUjL0gjM14CN4ITN) 接口获取。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uEDNyUjLxQjM14SM0ITN

POST https://www.feishu.cn/approval/openapi/v2/instance/get
*/
func InstanceGet(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiInstanceGet, bytes.NewReader(payload), header)
}

/*
创建审批实例


创建一个审批实例，调用方需对审批定义的表单有详细了解，将按照定义的表单结构，将表单 Value 通过接口传入。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uIDNyUjLyQjM14iM0ITN

POST https://www.feishu.cn/approval/openapi/v2/instance/create
*/
func InstanceCreate(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiInstanceCreate, bytes.NewReader(payload), header)
}

/*
审批任务同意


对于单个审批任务进行同意操作。同意后审批流程会流转到下一个审批人。


**权限说明** ：需要获取 访问审批应用 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMDNyUjLzQjM14yM0ITN

POST https://www.feishu.cn/approval/openapi/v2/instance/approve
*/
func Approve(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiApprove, bytes.NewReader(payload), header)
}

/*
审批任务拒绝


对于单个审批任务进行拒绝操作。拒绝后审批流程结束。

**权限说明** ：需要获取 访问审批应用 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQDNyUjL0QjM14CN0ITN

POST https://www.feishu.cn/approval/openapi/v2/instance/reject
*/
func Reject(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiReject, bytes.NewReader(payload), header)
}

/*
审批任务转交


对于单个审批任务进行转交操作。转交后审批流程流转给被转交人。

**权限说明** ：需要获取 访问审批应用 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uUDNyUjL1QjM14SN0ITN

POST https://www.feishu.cn/approval/openapi/v2/instance/transfer
*/
func Transfer(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiTransfer, bytes.NewReader(payload), header)
}

/*
审批实例撤回


对于单个审批实例进行撤销操作。撤销后审批流程结束。

**权限说明** ：需要获取 访问审批应用 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYDNyUjL2QjM14iN0ITN

POST https://www.feishu.cn/approval/openapi/v2/instance/cancel
*/
func Cancel(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiCancel, bytes.NewReader(payload), header)
}

/*
上传文件


当审批表单中有图片或附件控件时，开发者需在创建审批实例前通过审批上传文件接口将文件上传到审批系统。

**权限说明** ：需要获取 访问审批应用 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uUDOyUjL1gjM14SN4ITN

POST https://www.feishu.cn/approval/openapi/v2/file/upload
*/
func Upload(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiUpload, bytes.NewReader(payload), header)
}

/*
三方审批实例同步


用于第三方审批的实例同步。
使用前需确保已经在审批后台创建第三方审批。

> 审批实例：员工发起审批时产生的对象，详情参见 [开发指南](/ssl:ttdoc/ugTM5UjL4ETO14COxkTN/ukDNyUjL5QjM14SO0ITN)
>
**权限说明** ：需要获取 访问审批应用 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uczM3UjL3MzN14yNzcTN

POST https://www.feishu.cn/approval/openapi/v2/external/instance/create
*/
func ExternalInstanceCreate(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiExternalInstanceCreate, bytes.NewReader(payload), header)
}

/*
创建审批定义


用于通过接口创建简单的审批定义，可以灵活指定定义的基础信息、表单和流程等。不推荐企业自建应用使用，如有需要尽量联系管理员在审批管理后台创建定义。

**权限说明** ：需要获取 访问审批应用 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uUzNyYjL1cjM24SN3IjN

POST https://www.feishu.cn/approval/openapi/v2/approval/create
*/
func Create(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiCreate, bytes.NewReader(payload), header)
}

/*
审批实例抄送


通过接口可以将当前审批实例抄送给其他人。

**权限说明** ：需要获取 访问审批应用 权限。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uADOzYjLwgzM24CM4MjN

POST https://www.feishu.cn/approval/openapi/v2/instance/cc
*/
func InstanceCc(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiInstanceCc, bytes.NewReader(payload), header)
}

/*
订阅审批事件


订阅 approval_code 后，可以收到该审批定义对应实例的事件通知。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ucDOyUjL3gjM14yN4ITN

POST https://www.feishu.cn/approval/openapi/v2/subscription/subscribe
*/
func Subscribe(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiSubscribe, bytes.NewReader(payload), header)
}

/*
取消订阅审批事件


取消订阅 approval_code 后，无法再收到该审批定义对应实例的事件通知。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ugDOyUjL4gjM14CO4ITN

POST https://www.feishu.cn/approval/openapi/v2/subscription/unsubscribe
*/
func Unsubscribe(ctx *feishu.App, payload []byte) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiUnsubscribe, bytes.NewReader(payload), header)
}
