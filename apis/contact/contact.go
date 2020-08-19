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

// Package contact 获取企业自定义用户属性配置
package contact

import (
	"net/http"

	"github.com/fastwego/feishu"
)

const (
	apiTenantCustomAttrGet = "/open-apis/contact/v1/tenant/custom_attr/get"
)

/*
获取企业自定义用户属性配置


该接口用于获取企业配置的自定义用户属性。

**权限说明** ：调用该接口需要申请 获取部门组织架构信息 和 以应用身份访问通讯录 权限

::: note
 1. 调用该接口前，需要先确认企业管理员已经在 **企业管理后台** > **通讯录设置** 自定义信息栏开启了“允许开放平台API调用“。
 2. 调用该接口的应用需要具有当前企业通讯录的读取或者更新权限。
:::


See: https://open.feishu.cn/document/ukTMukTMukTM/ucTN3QjL3UzN04yN1cDN

GET https://open.feishu.cn/open-apis/contact/v1/tenant/custom_attr/get
*/
func TenantCustomAttrGet(ctx *feishu.App) (resp []byte, err error) {

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiTenantCustomAttrGet, header)
}
