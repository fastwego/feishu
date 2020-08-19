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

package main

type Param struct {
	Name string
	Type string
}

type Api struct {
	Name        string
	Description string
	Request     string
	See         string
	FuncName    string
	AccessToken string
	GetParams   []Param
	PathParams  []Param
}

type ApiGroup struct {
	Name    string
	Apis    []Api
	Package string
}

var apiConfig = []ApiGroup{
	{
		Name:    `身份认证`,
		Package: `authen`,
		Apis: []Api{
			{
				Name:        "获取用户授权跳转链接",
				Description: "应用请求用户身份验证时，需按如下方式构造登录链接，并引导用户跳转至此链接。飞书客户端内用户免登，系统浏览器内用户需完成扫码登录。登录成功后会生成登录预授权码 code，并作为参数重定向到重定向URL。",
				Request:     "GET https://open.feishu.cn/open-apis/authen/v1/index?redirect_uri={REDIRECT_URI}&app_id={APPID}&state={STATE}",
				See:         "https://open.feishu.cn/document/ukTMukTMukTM/ukzN4UjL5cDO14SO3gTN",
				FuncName:    "GetRedirectUri",
				GetParams: []Param{
					{Name: `app_id`, Type: `string`},
					{Name: `redirect_uri`, Type: `string`},
					{Name: `state`, Type: `string`},
				},
			},
			{
				Name:        "获取登录用户身份",
				Description: "通过此接口获取登录预授权码 code 对应的登录用户身份。",
				Request:     "POST https://open.feishu.cn/open-apis/authen/v1/access_token",
				See:         "https://open.feishu.cn/document/ukTMukTMukTM/uEDO4UjLxgDO14SM4gTN",
				FuncName:    "GetAccessToken",
			},
			{
				Name:        "刷新access_token",
				Description: "该接口用于在 access_token 过期时用 refresh_token 重新获取 access_token。此时会返回新的 refresh_token，再次刷新 access_token 时需要使用新的 refresh_token。",
				Request:     "POST https://open.feishu.cn/open-apis/authen/v1/refresh_access_token",
				See:         "https://open.feishu.cn/document/ukTMukTMukTM/uQDO4UjL0gDO14CN4gTN",
				FuncName:    "RefreshAccessToken",
			},
			{
				Name:        "获取用户信息（身份验证）",
				Description: "此接口仅用于获取登录用户的信息。调用此接口需要在 Header 中带上 user_access_token。\n\n",
				Request:     "GET https://open.feishu.cn/open-apis/authen/v1/user_info",
				See:         "https://open.feishu.cn/document/ukTMukTMukTM/uIDO4UjLygDO14iM4gTN",
				FuncName:    "GetUserInfo",
				GetParams: []Param{
					{Name: `user_access_token`, Type: `string`},
				},
			},
			{
				Name:        "小程序登录/code2session",
				Description: "通过 login 接口获取到登录凭证后，开发者可以通过服务器发送请求的方式获取 session_key 和 openId",
				Request:     "POST https://open.feishu.cn/open-apis/mina/v2/tokenLoginValidate",
				See:         "https://open.feishu.cn/document/uYjL24iN/ukjM04SOyQjL5IDN",
				FuncName:    "Code2Session",
			},
		},
	},
	{
		Name:    `用户管理`,
		Package: `contact/user`,
		Apis: []Api{
			{
				Name: "批量获取用户信息",
				Description: `
该接口用于批量获取用户详细信息。   

**权限说明** : 调用该接口需要申请“以应用身份访问通讯录”以及[用户数据权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)。请求的用户如果在当前应用的通讯录授权范围内，会返回该用户的详细信息；否则不会返回。
`,
				Request:     "GET https://open.feishu.cn/open-apis/contact/v1/user/batch_get?employee_ids=2fab1234&employee_ids=2f1234cd",
				See:         "https://open.feishu.cn/document/ukTMukTMukTM/uIzNz4iM3MjLyczM",
				FuncName:    "BatchGet",
				AccessToken: "tenant",
				GetParams: []Param{
					{Name: `employee_ids`, Type: `string`},
				},
			},
			{
				Name: "获取部门用户列表",
				Description: `
该接口用于获取部门用户列表。  
**权限说明**：调用该接口需要申请 获取部门组织架构信息 和 以应用身份访问通讯录 权限。应用需要有被调用部门的通讯录授权
`,
				Request:  "GET https://open.feishu.cn/open-apis/contact/v1/department/user/list?department_id=TT-1234&page_size=100&fetch_child=true",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uEzNz4SM3MjLxczM",
				FuncName: "DepartmentUserList",
				GetParams: []Param{
					{Name: `page_size`, Type: `string`},
					{Name: `fetch_child`, Type: `string`},
					{Name: `department_id`, Type: `string`},
				},
			},
			{
				Name: "获取部门用户详情",
				Description: `
该接口用于获取部门用户详情信息。
**权限说明**：调用该接口需要申请 获取部门组织架构信息 和 以应用身份访问通讯录 权限，同时根据需要返回字段申请对应的[用户数据权限]。应用需要有被调用部门的通讯录授权。(https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)
`,
				Request:  "GET https://open.feishu.cn/open-apis/contact/v1/department/user/detail/list?department_id=TT-1234&page_size=10&fetch_child=true",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uYzN3QjL2czN04iN3cDN",
				FuncName: "DepartmentUserDetailList",
				GetParams: []Param{
					{Name: `department_id`, Type: `string`},
					{Name: `page_size`, Type: `string`},
					{Name: `fetch_child`, Type: `string`},
				},
			},
			{
				Name: "新增用户",
				Description: `
该接口用于向通讯录中新增用户。  

**权限说明** ：调用该接口需要申请 更新通讯录 以及 以应用身份访问通讯录 权限。新增用户的所有部门必须都在当前应用的通讯录授权范围内才允许新增用户。应用商店应用无权限调用此接口。
`,
				Request:  "POST https://open.feishu.cn/open-apis/contact/v1/user/add",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uMzNz4yM3MjLzczM",
				FuncName: "Add",
			},
			{
				Name: "删除用户",
				Description: `
该接口用于从通讯录中删除用户。  

**权限说明** : 调用该接口需要申请 更新通讯录 以及 以应用身份访问通讯录 权限。应用需要有待删除用户、待删除用户的所有部门的通讯录权限才能删除该用户。应用商店应用无权限调用接口。

`,
				Request:  "POST https://open.feishu.cn/open-apis/contact/v1/user/delete",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uUzNz4SN3MjL1czM",
				FuncName: "Delete",
			},
			{
				Name: "更新用户信息",
				Description: `
该接口用于更新通讯录中用户信息。  

**权限说明** : 调用该接口需要申请 更新通讯录 以及 以应用身份访问通讯录 权限。应用需要拥有待更新用户的通讯录授权，如果涉及到用户部门变更，还需要同时拥有所有新部门的通讯录授权。应用商店应用无权限调用此接口。

`,
				Request:  "POST https://open.feishu.cn/open-apis/contact/v1/user/update",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uQzNz4CN3MjL0czM",
				FuncName: "Update",
			},
			{
				Name: "获取角色列表",
				Description: `
该接口用于获取企业的用户角色列表。

**权限说明** ：调用该接口需要申请 获取角色 和 以应用身份访问通讯录 权限
`,
				Request:  "GET https://open.feishu.cn/open-apis/contact/v2/role/list",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uYzMwUjL2MDM14iNzATN",
				FuncName: "RoleList",
			},
			{
				Name: "获取角色成员列表",
				Description: `
该接口用于获取角色下的用户列表。

**权限说明** ：调用该接口需要申请 获取角色 和 以应用身份访问通讯录 权限。角色内不在当前应用通讯录授权范围内的用户不会返回。
`,
				Request:  "GET https://open.feishu.cn/open-apis/contact/v2/role/members?role_id=or_846ea69995a259a27cc690182f27de87&page_size=2&page_token=763bd1e74d05e958",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uczMwUjL3MDM14yNzATN",
				FuncName: "RoleMembers",
				GetParams: []Param{
					{Name: `page_token`, Type: `string`},
					{Name: `role_id`, Type: `string`},
					{Name: `page_size`, Type: `string`},
				},
			},
			{
				Name: "使用手机号或邮箱获取用户 ID",
				Description: `ID

根据用户邮箱或手机号查询用户 open_id 和 user_id，支持批量查询。
调用该接口需要申请  通过手机号或者邮箱获取用户ID  权限。

:::warnning
只能查询到应用可用性范围内的用户 ID，不在范围内的用户会表现为不存在。
:::
`,
				Request:  "GET https://open.feishu.cn/open-apis/user/v1/batch_get_id?emails=lisi@z.com&emails=wangwu@z.com&mobiles=13812345678&mobiles=%2b12126668888",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uUzMyUjL1MjM14SNzITN",
				FuncName: "BatchGetId",
				GetParams: []Param{
					{Name: `emails`, Type: `string`},
					{Name: `mobiles`, Type: `string`},
				},
			},
			{
				Name: "使用统一 ID 获取用户 ID",
				Description: `ID 获取用户 ID
使用统一 ID 获取用户 ID信息。
调用该接口需要具有 获取用户统一ID  权限。
`,
				Request:  "GET https://open.feishu.cn/open-apis/user/v1/union_id/batch_get/list?union_ids=on_94a1ee5551019f18cd73d9f111898cf2&union_ids=on_42f2ef9d07319a4d96fffd7ef5cbfc79",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uUTO5UjL1kTO14SN5kTN",
				FuncName: "UnionIdBatchGetList",
				GetParams: []Param{
					{Name: `union_ids`, Type: `string`},
				},
			},
			{
				Name: "搜索用户",
				Description: `
以用户身份搜索其他用户的信息，无法搜索到外部企业或已离职的用户。
调用该接口需要申请  搜索用户  权限。
`,
				Request:     "GET https://open.feishu.cn/open-apis/search/v1/user?query=zhangsan&page_size=20&page_token=20",
				See:         "https://open.feishu.cn/document/ukTMukTMukTM/uMTM4UjLzEDO14yMxgTN",
				FuncName:    "Search",
				AccessToken: "user",
				GetParams: []Param{
					{Name: `query`, Type: `string`},
					{Name: `page_size`, Type: `string`},
					{Name: `page_token`, Type: `string`},
				},
			},
		},
	},
	{
		Name:    `部门管理`,
		Package: `contact/department`,
		Apis: []Api{
			{
				Name: "获取部门详情",
				Description: `该接口用于获取部门详情信息。  

**权限说明** ：调用该接口需要具有 以应用身份访问通讯录  以及 [部门数据权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)。应用需要拥有待查询部门的通讯录授权。
`,
				Request:  "GET https://open.feishu.cn/open-apis/contact/v1/department/info/get?department_id=TT-1234",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uAzNz4CM3MjLwczM",
				FuncName: "InfoGet",
				GetParams: []Param{
					{Name: `department_id`, Type: `string`},
				},
			},
			{
				Name: "获取子部门列表",
				Description: `
该接口用于获取当前部门子部门列表。

**权限说明**：调用该接口需要申请 获取部门组织架构信息 以及 以应用身份访问通讯录 权限。调用该接口需要具有当前部门的授权范围。企业根部门 ID 为 0，当获取根部门子部门列表时，通讯录授权范围必须为全员权限。
`,
				Request:  "GET https://open.feishu.cn/open-apis/contact/v1/department/simple/list?department_id=TT-1234&page_size=10&fetch_child=true",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/ugzN3QjL4czN04CO3cDN",
				FuncName: "SimpleList",
				GetParams: []Param{
					{Name: `fetch_child`, Type: `string`},
					{Name: `department_id`, Type: `string`},
					{Name: `page_size`, Type: `string`},
				},
			},
			{
				Name: "批量获取部门详情",
				Description: `
该接口用于批量获取部门详情，只返回权限范围内的部门。
**权限说明** : 调用该接口需要申请 以应用身份访问通讯录 以及[部门数据权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)。
`,
				Request:  "GET https://open.feishu.cn/open-apis/contact/v1/department/detail/batch_get?department_ids=od-2efe30807a10608754862a63b108828f&department_ids=od-da6427b2adbceb91204d7fa6aeb7e8ff",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uczN3QjL3czN04yN3cDN",
				FuncName: "DetailBatchGet",
				GetParams: []Param{
					{Name: `department_ids`, Type: `string`},
				},
			},
			{
				Name: "新增部门",
				Description: `
该接口用于向通讯录中增加新的部门。

**权限说明** ：调用该接口需要申请 更新通讯录 以及 以应用身份访问通讯录 权限。应用需要拥有待新增部门的父部门的通讯录授权。应用商店应用无权限调用接口。

`,
				Request:  "POST https://open.feishu.cn/open-apis/contact/v1/department/add",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uYzNz4iN3MjL2czM",
				FuncName: "Add",
			},
			{
				Name: "删除部门",
				Description: `该接口用于从通讯录中删除部门。 

**权限说明** : 调用该接口需要申请 更新通讯录 以及 以应用身份访问通讯录 权限。应用需要同时拥有待删除部门及其父部门的通讯录授权。应用商店应用无权限调用该接口。 

`,
				Request:  "POST https://open.feishu.cn/open-apis/contact/v1/department/delete",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/ugzNz4CO3MjL4czM",
				FuncName: "Delete",
			},
			{
				Name: "更新部门信息",
				Description: `该接口用于更新通讯录中部门的信息。

**权限说明**：调用该接口需要申请 更新通讯录 以及 以应用身份访问通讯录 权限。调用该接口需要具有该部门以及更新操作涉及的部门的通讯录权限。应用商店应用无权限调用此接口。

`,
				Request:  "POST https://open.feishu.cn/open-apis/contact/v1/department/update",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uczNz4yN3MjL3czM",
				FuncName: "Update",
			},
		},
	},
	{
		Name:    `异步批量接口`,
		Package: `contact/async_batch`,
		Apis: []Api{
			{
				Name: "批量新增部门",
				Description: `
该接口用于向通讯录中批量新增多个部门。
调用该接口需要具有所有新增部门父部门的授权范围。
应用商店应用无权限调用此接口。
调用该接口需要申请  更新通讯录  以及  以应用身份访问通讯录  权限。
`,
				Request:  "POST https://open.feishu.cn/open-apis/contact/v2/department/batch_add",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uMDOwUjLzgDM14yM4ATN",
				FuncName: "DepartmentBatchAdd",
			},
			{
				Name: "批量新增用户",
				Description: `
该接口用于向通讯录中批量新增多个用户。
调用该接口需要具有待添加用户所在部门的通讯录授权范围。
应用商店应用无权限调用此接口。
调用该接口需要申请  更新通讯录  以及  以应用身份访问通讯录  权限。
`,
				Request:  "POST https://open.feishu.cn/open-apis/contact/v2/user/batch_add",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uIDOwUjLygDM14iM4ATN",
				FuncName: "UserBatchAdd",
			},
			{
				Name: "查询批量任务执行状态",
				Description: `
该接口用于查询通讯录异步任务当前的执行状态以及执行结果。
应用商店应用无权限调用此接口。
调用该接口需要申请  更新通讯录  以及  以应用身份访问通讯录  权限。
`,
				Request:  "GET https://open.feishu.cn/open-apis/contact/v2/task/get?task_id=0123456789abcdef0123456789abcdef",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uUDOwUjL1gDM14SN4ATN",
				FuncName: "TaskGet",
				GetParams: []Param{
					{Name: `task_id`, Type: `string`},
				},
			},
		},
	},
	{
		Name:    `获取企业自定义用户属性配置`,
		Package: `contact`,
		Apis: []Api{
			{
				Name: "获取企业自定义用户属性配置",
				Description: `
该接口用于获取企业配置的自定义用户属性。

**权限说明** ：调用该接口需要申请 获取部门组织架构信息 和 以应用身份访问通讯录 权限

::: note
 1. 调用该接口前，需要先确认企业管理员已经在 **企业管理后台** > **通讯录设置** 自定义信息栏开启了“允许开放平台API调用“。
 2. 调用该接口的应用需要具有当前企业通讯录的读取或者更新权限。
:::
`,
				Request:  "GET https://open.feishu.cn/open-apis/contact/v1/tenant/custom_attr/get",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/ucTN3QjL3UzN04yN1cDN",
				FuncName: "TenantCustomAttrGet",
			},
		},
	},
	{
		Name:    `用户群组`,
		Package: `user_group`,
		Apis: []Api{
			{
				Name: "获取用户所在的群列表",
				Description: `获取用户所在的群列表。

**权限说明** ：应用需要“读取群信息”权限；
`,
				Request:  "GET https://open.feishu.cn/open-apis/user/v4/group_list?page_size=2&page_token=6592161138799017988",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uQzMwUjL0MDM14CNzATN",
				FuncName: "GroupList",
				GetParams: []Param{
					{Name: `page_size`, Type: `string`},
					{Name: `page_token`, Type: `string`},
				},
			},
			{
				Name: "获取群成员列表",
				Description: `如果用户在群中，则返回该群的成员列表。

**权限说明** ：应用需要“读取群信息”权限；
`,
				Request:  "GET https://open.feishu.cn/open-apis/chat/v4/members?chat_id=oc_92c3f700c2ae31369cefee459fb93870&page_token=0&page_size=3",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uUzMwUjL1MDM14SNzATN",
				FuncName: "Members",
				GetParams: []Param{
					{Name: `chat_id`, Type: `string`},
					{Name: `page_token`, Type: `string`},
					{Name: `page_size`, Type: `string`},
				},
			},
			{
				Name: "搜索用户所在的群列表",
				Description: `搜索用户所在的群列表。

**权限说明** ：应用需要“读取群信息”权限；
`,
				Request:  "GET https://open.feishu.cn/open-apis/chat/v4/search?page_size=10&page_token=24&query=rd",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uUTOyUjL1kjM14SN5ITN",
				FuncName: "Search",
				GetParams: []Param{
					{Name: `page_size`, Type: `string`},
					{Name: `page_token`, Type: `string`},
					{Name: `query`, Type: `string`},
				},
			},
		},
	},
	{
		Name:    `应用信息/应用管理`,
		Package: `app/app_manage`,
		Apis: []Api{
			{
				Name: "校验应用管理员",
				Description: `
该接口用于查询用户是否为应用管理员，需要申请 校验用户是否为应用管理员 权限。
`,
				Request:  "GET https://open.feishu.cn/open-apis/application/v3/is_user_admin",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uITN1EjLyUTNx4iM1UTM",
				FuncName: "IsUserAdmin",
				GetParams: []Param{
					{Name: `open_id`, Type: `string`},
				},
			},
			{
				Name: "获取应用管理员管理范围",
				Description: `该接口用于获取应用管理员的管理范围，即该应用管理员能够管理哪些部门。  

`,
				Request:  "GET https://open.feishu.cn/open-apis/contact/v1/user/admin_scope/get?employee_id=2fab1234",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uMzN3QjLzczN04yM3cDN",
				FuncName: "AdminScopeGet",
				GetParams: []Param{
					{Name: `employee_id`, Type: `string`},
				},
			},
			{
				Name: "获取应用在企业内的可用范围",
				Description: `
该接口用于查询应用在该企业内可以被使用的范围，只能被企业自建应用调用且需要“获取应用信息”权限。
`,
				Request:  "GET https://open.feishu.cn/open-apis/application/v2/app/visibility?app_id=cli_9db45f86b7799104&user_page_token=0&user_page_size=100",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uIjM3UjLyIzN14iMycTN",
				FuncName: "AppVisibility",
				GetParams: []Param{
					{Name: `app_id`, Type: `string`},
					{Name: `user_page_token`, Type: `string`},
					{Name: `user_page_size`, Type: `string`},
				},
			},
			{
				Name: "获取用户可用的应用",
				Description: `
该接口用于查询用户可用的应用列表，只能被企业自建应用调用且需要“获取用户可用的应用”权限。
`,
				Request:  "GET https://open.feishu.cn/open-apis/application/v1/user/visible_apps?user_id=79affdge&page_token=0&lang=zh_cn&page_size=5",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uMjM3UjLzIzN14yMycTN",
				FuncName: "VisibleApps",
				GetParams: []Param{
					{Name: `page_size`, Type: `string`},
					{Name: `user_id`, Type: `string`},
					{Name: `page_token`, Type: `string`},
					{Name: `lang`, Type: `string`},
				},
			},
			{
				Name: "获取企业安装的应用",
				Description: `
该接口用于查询企业安装的应用列表，只能被企业自建应用调用且需要“获取应用信息”权限。
`,
				Request:  "GET https://open.feishu.cn/open-apis/application/v3/app/list?page_size=5&page_token=0&lang=zh_cn&status=1",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uYDN3UjL2QzN14iN0cTN",
				FuncName: "AppList",
				GetParams: []Param{
					{Name: `lang`, Type: `string`},
					{Name: `status`, Type: `string`},
					{Name: `page_size`, Type: `string`},
					{Name: `page_token`, Type: `string`},
				},
			},
			{
				Name: "更新应用可用范围",
				Description: `
该接口用于增加或者删除指定应用被哪些人可用，只能被企业自建应用调用且需要“管理应用可见范围”权限。
`,
				Request:  "POST https://open.feishu.cn/open-apis/application/v3/app/update_visibility",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/ucDN3UjL3QzN14yN0cTN",
				FuncName: "UpdateVisibility",
			},
			{
				Name: "查询应用管理员列表",
				Description: `查询应用管理员列表，返回应用的最新10个管理员账户id列表。

**权限说明** ：
本接口需要申请  获取应用管理员ID  权限才能访问。
回包数据中的user_id 需要申请  获取用户 userid  权限才会返回
回包数据中的union_id 需要申请  获取用户统一ID  权限才会返回
`,
				Request:  "GET https://open.feishu.cn/open-apis/user/v4/app_admin_user/list",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/ucDOwYjL3gDM24yN4AjN",
				FuncName: "AppAdminUserList",
			},
		},
	},
	{
		Name:    `应用信息/应用商店`,
		Package: `app/app_store`,
		Apis: []Api{
			{
				Name:        "查询用户是否在应用开通范围",
				Description: `该接口用于查询用户是否在企业管理员设置的使用该应用的范围中。如果设置的付费套餐是按人收费或者限制了最大人数，开放平台会引导企业管理员设置“付费功能开通范围”，本接口用于查询用户是否在企业管理员设置的使用该应用的范围中，可以通过此接口，在付费功能点入口判断是否允许某个用户进入使用。`,
				Request:     "GET https://open.feishu.cn/open-apis/pay/v1/paid_scope/check_user?open_id=ou_5ad573a6411d72b8305fda3a9c15c70e",
				See:         "https://open.feishu.cn/document/ukTMukTMukTM/uATNwUjLwUDM14CM1ATN",
				FuncName:    "CheckUser",
				GetParams: []Param{
					{Name: `open_id`, Type: `string`},
				},
			},
			{
				Name: "查询租户购买的付费方案",
				Description: `该接口用于分页查询应用租户下的已付费订单，每次购买对应一个唯一的订单，订单会记录购买的套餐的相关信息，业务方需要自行处理套餐的有效期和付费方案的升级。
`,
				Request:  "GET https://open.feishu.cn/open-apis/pay/v1/order/list?status=all&page_size=10&page_token=10&tenant_key=2e5c3a3ae38f175f",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uETNwUjLxUDM14SM1ATN",
				FuncName: "OrderList",
				GetParams: []Param{
					{Name: `status`, Type: `string`},
					{Name: `page_size`, Type: `string`},
					{Name: `page_token`, Type: `string`},
					{Name: `tenant_key`, Type: `string`},
				},
			},
			{
				Name: "查询订单详情",
				Description: `该接口用于查询某个订单的具体信息
`,
				Request:  "GET https://open.feishu.cn/open-apis/pay/v1/order/get?order_id=6708978506916697671",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uITNwUjLyUDM14iM1ATN",
				FuncName: "OrderGet",
				GetParams: []Param{
					{Name: `order_id`, Type: `string`},
				},
			},
		},
	},
	{
		Name:    `机器人/群信息和群管理`,
		Package: `bot/group_manage`,
		Apis: []Api{
			{
				Name: "创建群",
				Description: `机器人创建群并拉指定用户进群。  

**权限说明** ：需要启用机器人能力
`,
				Request:  "POST https://open.feishu.cn/open-apis/chat/v4/create/",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/ukDO5QjL5gTO04SO4kDN",
				FuncName: "ChatCreate",
			},
			{
				Name: "获取群列表",
				Description: `获取机器人所在的群列表。

**权限说明** ：需要启用机器人能力
`,
				Request:  "GET https://open.feishu.cn/open-apis/chat/v4/list",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uITO5QjLykTO04iM5kDN",
				FuncName: "ChatList",
				GetParams: []Param{
					{Name: `page_size`, Type: `page_size`},
				},
			},
			{
				Name: "获取群信息",
				Description: `获取群名称、群主 ID、成员列表 ID 等群基本信息。  

**权限说明** ：需要启用机器人能力；机器人必须在群里
`,
				Request:  "GET https://open.feishu.cn/open-apis/chat/v4/info?chat_id=oc_eb9e82d5657777ebf1bb5b9024f549ef",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uMTO5QjLzkTO04yM5kDN",
				FuncName: "ChatInfo",
				GetParams: []Param{
					{Name: `chat_id`, Type: `string`},
				},
			},
			{
				Name: "更新群信息",
				Description: `更新群名称、群配置、转让群主等。

**权限说明** ：需要启用机器人能力；机器人必须是群主，才能执行修改群配置和转让群主的操作（机器人创建的群，机器人默认是群主。）
`,
				Request:  "POST https://open.feishu.cn/open-apis/chat/v4/update/",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uYTO5QjL2kTO04iN5kDN",
				FuncName: "ChatUpdate",
			},
			{
				Name: "拉用户进群",
				Description: `机器人拉用户进群，机器人必须在群里。

**权限说明** ：需要启用机器人能力；机器人必须在群里
`,
				Request:  "POST https://open.feishu.cn/open-apis/chat/v4/chatter/add/",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/ucTO5QjL3kTO04yN5kDN",
				FuncName: "ChatterAdd",
			},
			{
				Name: "移除用户出群",
				Description: `机器人移除用户出群。

**权限说明** ：需要启用机器人能力；机器人必须是群主（机器人创建的群，机器人默认是群主。）
`,
				Request:  "POST https://open.feishu.cn/open-apis/chat/v4/chatter/delete/",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uADMwUjLwADM14CMwATN",
				FuncName: "ChatterDelete",
			},
			{
				Name: "解散群",
				Description: `机器人解散群。

**权限说明** ：需要启用机器人能力；机器人必须是群主（机器人创建的群，机器人默认是群主。）
`,
				Request:  "POST https://open.feishu.cn/open-apis/chat/v4/disband",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uUDN5QjL1QTO04SN0kDN",
				FuncName: "Disband",
			},
		},
	},
	{
		Name:    `机器人/机器人信息和管理`,
		Package: `bot/bot_manage`,
		Apis: []Api{
			{
				Name: "获取机器人信息",
				Description: `获取机器人的基本信息

**权限说明** ：需要启用机器人能力
`,
				Request:  "GET https://open.feishu.cn/open-apis/bot/v3/info/",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uAjMxEjLwITMx4CMyETM",
				FuncName: "Info",
			},
			{
				Name: "拉机器人进群",
				Description: `拉机器人进群

**权限说明** ：需要启用机器人能力；机器人的owner需要已经在群里
`,
				Request:  "POST https://open.feishu.cn/open-apis/bot/v4/add",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uYDO04iN4QjL2gDN",
				FuncName: "Add",
			},
			{
				Name: "将机器人移出群",
				Description: `将机器人移出群。

**权限说明** ：需要启用机器人能力
`,
				Request:  "POST https://open.feishu.cn/open-apis/bot/v4/remove",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/ucDO04yN4QjL3gDN",
				FuncName: "Remove",
			},
		},
	},
	{
		Name:    `消息`,
		Package: `message`,
		Apis: []Api{
			{
				Name: "批量发送消息",
				Description: `
给多个用户或者多个部门发送消息。

**权限说明** ：需要启用机器人能力；机器人需要拥有批量发送消息权限；机器人需要拥有对用户或部门的可见性  

`,
				Request:  "POST https://open.feishu.cn/open-apis/message/v4/batch_send/",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/ucDO1EjL3gTNx4yN4UTM",
				FuncName: "BatchSend",
			},
			{
				Name: "发送消息",
				Description: `
给指定用户或者会话发送文本/图片/富文本/群名片/消息卡片 消息，其中会话包括私聊会话和群会话。

**权限说明** ：需要启用机器人能力；私聊会话时机器人需要拥有对用户的可见性，群会话需要机器人在群里  
`,
				Request:  "POST https://open.feishu.cn/open-apis/message/v4/send/",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uUjNz4SN2MjL1YzM",
				FuncName: "Send",
			},
			{
				Name: "查询消息已读状态",
				Description: `
查询消息已读状态，只能查询最近七天机器人自身发送消息的已读信息。

**权限说明** ：需要启用机器人能力
`,
				Request:  "POST https://open.feishu.cn/open-apis/message/v4/read_info/",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/ukTM2UjL5EjN14SOxYTN",
				FuncName: "ReadInfo",
			},
			{
				Name: "上传图片",
				Description: `
上传图片，获取图片的 image_key。

**权限说明** ：需要启用机器人能力
`,
				Request:  "POST https://open.feishu.cn/open-apis/image/v4/put/",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uEDO04SM4QjLxgDN",
				FuncName: "ImagePut",
			},
			{
				Name: "获取图片",
				Description: `
根据图片的image_key获取图片内容

**权限说明** ：需要启用机器人能力；机器人只能获取自己上传或者收到推送的图片
`,
				Request:  "GET https://open.feishu.cn/open-apis/image/v4/get?image_key=24383920-9321-4ecd-8b33-bf8ce74e84c8",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uYzN5QjL2cTO04iN3kDN",
				FuncName: "ImageGet",
				GetParams: []Param{
					{Name: `image_key`, Type: `string`},
				},
			},
			{
				Name: "获取文件",
				Description: `
根据文件的 file_key 拉取文件内容，当前仅可用来获取用户与机器人单聊发送的文件

**权限说明** ：需要启用机器人能力；机器人只能获取自己上传的文件
`,
				Request:  "GET https://open.feishu.cn/open-apis/open-file/v1/get?file_key=file_36r377cb-c6h2-4b6d-ag67-0ac3e796008g",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uMDN4UjLzQDO14yM0gTN",
				FuncName: "FileGet",
				GetParams: []Param{
					{Name: `file_key`, Type: `string`},
				},
			},
			{
				Name: "应用发送通知给用户",
				Description: `给指定用户发送应用通知

**权限说明** ：需要申请  以应用的身份发消息 的权限；同时需要启用小程序能力，并且支持**PC模式**。
`,
				Request:  "POST https://open.feishu.cn/open-apis/notify/v4/appnotify",
				See:      "https://open.feishu.cn/document/ukTMukTMukTM/uATOyUjLwkjM14CM5ITN",
				FuncName: "AppNotify",
			},
		},
	},
}
