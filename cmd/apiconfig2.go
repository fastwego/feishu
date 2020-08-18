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

var apiConfig2 = []ApiGroup{
	{
		Name:    `日历`,
		Package: `capabilities/calendar`,
		Apis: []Api{
			{
				Name: "获取日历",
				Description: `
该接口用于根据日历 ID 获取日历信息。
`,
				Request:  "GET https://open.feishu.cn/open-apis/calendar/v3/calendar_list/:calendarId",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMDN04yM0QjLzQDN",
				FuncName: "GetCalendarById",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "获取日历列表",
				Description: `
该接口用于获取应用在企业内的日历列表。
`,
				Request:  "GET https://open.feishu.cn/open-apis/calendar/v3/calendar_list",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMTM14yMxUjLzETN",
				FuncName: "CalendarList",
			},
			{
				Name: "创建日历",
				Description: `
该接口用于为应用在企业内创建一个日历。
`,
				Request:  "POST https://open.feishu.cn/open-apis/calendar/v3/calendars",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQTM14CNxUjL0ETN",
				FuncName: "CreateCalendars",
			},
			{
				Name: "删除日历",
				Description: `
该接口用于删除应用在企业内的指定日历。
`,
				Request:  "DELETE https://open.feishu.cn/open-apis/calendar/v3/calendars/:calendarId",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uUTM14SNxUjL1ETN",
				FuncName: "DeleteCalendarById",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "获取日程",
				Description: `
该接口用于获取指定日历下的指定日程。
`,
				Request:  "GET https://open.feishu.cn/open-apis/calendar/v3/calendars/:calendarId/events/:eventId",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ucTM14yNxUjL3ETN",
				FuncName: "GetEventById",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "创建日程",
				Description: `
该接口用于在日历中创建一个日程。
`,
				Request:  "POST https://open.feishu.cn/open-apis/calendar/v3/calendars/:calendarId/events",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ugTM14COxUjL4ETN",
				FuncName: "CreateEvent",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "邀请/移除日程参与者",
				Description: `
邀请一个或多个用户加入日程；
从日程移除一个或多个用户。
`,
				Request:  "POST https://open.feishu.cn/open-apis/calendar/v3/calendars/:calendarId/events/:eventId/attendees",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uIjM14iMyUjLyITN",
				FuncName: "Attendees",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "获取访问控制列表",
				Description: `
该接口用于查看指定日历的成员列表。
`,
				Request:  "GET https://open.feishu.cn/open-apis/calendar/v3/calendars/:calendarId/acl",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMjM14yMyUjLzITN",
				FuncName: "Acl",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "删除访问控制",
				Description: `
该接口用于从日历中移除一个用户。
`,
				Request:  "DELETE https://open.feishu.cn/open-apis/calendar/v3/calendars/:calendarId/acl/:ruleId",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uUjM14SNyUjL1ITN",
				FuncName: "DeleteAclByRuleId",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "查询日历的忙闲状态",
				Description: `
该接口用于查询指定日历或用户主日历的忙闲状态。
`,
				Request:  "POST https://open.feishu.cn/open-apis/calendar/v3/freebusy/query",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYjM14iNyUjL2ITN",
				FuncName: "FreeBusyQuery",
			},
			{
				Name: "查询公共日历",
				Description: `
该接口用于通过关键字查询公共日历信息。
`,
				Request:  "GET https://open.feishu.cn/open-apis/calendar/v3/shared_calendar_list/shared_calendar/query?query=ByteDance",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukDMwYjL5ADM24SOwAjN",
				FuncName: "SharedCalendarQuery",
				GetParams: []Param{
					{Name: `query`, Type: `string`},
				},
			},
			{
				Name: "获取公共日历日程列表",
				Description: `
该接口用于获取公共日历下的日程列表。
`,
				Request:  "GET https://open.feishu.cn/open-apis/calendar/v3/shared/calendars/:calendarId/events",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uIzNwYjLycDM24iM3AjN",
				FuncName: "SharedCalendarEvents",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
		},
	},
	{
		Name:    `云文档`,
		Package: `capabilities/document`,
		Apis: []Api{

			{
				Name: "新建文件夹",
				Description: `

该接口用于根据 folderToken 在该 folder 下创建文件夹
`,
				Request:     "POST https://open.feishu.cn/open-apis/drive/explorer/v2/folder/:folderToken",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukTNzUjL5UzM14SO1MTN",
				FuncName:    "CreateFolder",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "获取文件夹元信息",
				Description: `
# 获取root folder（我的空间） meta该接口用于获取 我的空间 的元信息
`,
				Request:     "GET https://open.feishu.cn/open-apis/drive/explorer/v2/root_folder/meta",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uAjNzUjLwYzM14CM2MTN",
				FuncName:    "RootFolderMeta",
				AccessToken: "user",
			},
			{
				Name: "获取文件夹下的文档清单",
				Description: `

该接口用于根据 folderToken 获取该文件夹的文档清单，如 doc、sheet、folder
`,
				Request:     "GET https://open.feishu.cn/open-apis/drive/explorer/v2/folder/:folderToken/children",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uEjNzUjLxYzM14SM2MTN",
				FuncName:    "FolderChildren",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "新建文档",
				Description: `
该接口用于根据 folderToken 创建 Doc或 Sheet 。

若没有特定的文件夹用于承载创建的文档，可以先调用「获取文件夹元信息」文档中的「获取 root folder (我的空间) meta」接口，获得我的空间的 token，然后再使用此接口。创建的文档将会在「我的空间」的「归我所有」列表里。
`,
				Request:     "POST https://open.feishu.cn/open-apis/drive/explorer/v2/file/:folderToken",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQTNzUjL0UzM14CN1MTN",
				FuncName:    "CreateFile",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "复制文档",
				Description: `

该接口用于根据文件 token 复制 Doc 或 Sheet  到目标文件夹中。

若没有特定的文件夹用于承载创建的文档，可以先调用「获取文件夹元信息」文档中的「获取 root folder (我的空间) meta」接口，获得我的空间的 token，然后再使用此接口。复制的文档将会在「我的空间」的「归我所有」列表里。
`,
				Request:     "POST https://open.feishu.cn/open-apis/drive/explorer/v2/file/copy/files/:fileToken",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYTNzUjL2UzM14iN1MTN",
				FuncName:    "CopyFile",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "删除文档",
				Description: `Doc 

该接口用于根据 spreadsheetToken 删除对应的 sheet 文档。
`,
				Request:     "DELETE https://open.feishu.cn/open-apis/drive/explorer/v2/file/spreadsheets/:spreadsheetToken",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uATM2UjLwEjN14CMxYTN",
				FuncName:    "DeleteFile",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "增加权限",
				Description: `

该接口用于根据 filetoken 给用户增加文档的权限
`,
				Request:     "POST https://open.feishu.cn/open-apis/drive/permission/member/create",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMzNzUjLzczM14yM3MTN",
				FuncName:    "PermissionMemberCreate",
				AccessToken: "user",
			},
			{
				Name: "转移拥有者",
				Description: `

该接口用于根据文档信息和用户信息转移文档的所有者。
`,
				Request:     "POST https://open.feishu.cn/open-apis/drive/permission/member/transfer",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQzNzUjL0czM14CN3MTN",
				FuncName:    "PermissionMemberTransfer",
				AccessToken: "user",
			},
			{
				Name: "更新文档公共设置",
				Description: `

该接口用于根据 filetoken 更新文档的公共设置
`,
				Request:     "POST https://open.feishu.cn/open-apis/drive/permission/public/update",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukTM3UjL5EzN14SOxcTN",
				FuncName:    "PermissionPublicUpdate",
				AccessToken: "user",
			},
			{
				Name: "获取协作者列表",
				Description: `

该接口用于根据 filetoken 查询协作者，目前包括人('user')和群('chat') 

**你能获取到协作者列表的前提是你对该文档有权限**
`,
				Request:     "POST https://open.feishu.cn/open-apis/drive/permission/member/list",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uATN3UjLwUzN14CM1cTN",
				FuncName:    "PermissionMemberList",
				AccessToken: "user",
			},
			{
				Name: "移除协作者权限",
				Description: `

该接口用于根据 filetoken 移除文档协作者的权限
`,
				Request:     "POST https://open.feishu.cn/open-apis/drive/permission/member/delete",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYTN3UjL2UzN14iN1cTN",
				FuncName:    "PermissionMemberDelete",
				AccessToken: "user",
			},
			{
				Name: "更新协作者权限",
				Description: `

该接口用于根据 filetoken 更新文档协作者的权限
`,
				Request:     "POST https://open.feishu.cn/open-apis/drive/permission/member/update",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ucTN3UjL3UzN14yN1cTN",
				FuncName:    "PermissionMemberUpdate",
				AccessToken: "user",
			},
			{
				Name: "判断协作者是否有某权限",
				Description: `

该接口用于根据 filetoken 判断当前登录用户是否具有某权限
`,
				Request:     "POST https://open.feishu.cn/open-apis/drive/permission/member/permitted",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYzN3UjL2czN14iN3cTN",
				FuncName:    "PermissionMemberPermitted",
				AccessToken: "user",
			},
			{
				Name: "获取文档文本内容",
				Description: `
该接口用于获取文档的纯文本内容，不包含富文本格式信息，主要用于搜索，如导入 es 等。
`,
				Request:     "GET https://open.feishu.cn/open-apis/doc/v2/:docToken/raw_content",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukzNzUjL5czM14SO3MTN",
				FuncName:    "RawContent",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "获取sheet@doc的元数据",
				Description: `sheet@doc 的元数据

该接口用于根据 docToken 获取 sheet@doc 的元数据。
`,
				Request:     "GET https://open.feishu.cn/open-apis/doc/v2/:docToken/sheet_meta",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uADOzUjLwgzM14CM4MTN",
				FuncName:    "SheetMeta",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "获取文档元信息",
				Description: `
该接口用于根据 docToken 获取元数据。
`,
				Request:     "GET https://open.feishu.cn/open-apis/doc/v2/meta/:docToken",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uczN3UjL3czN14yN3cTN",
				FuncName:    "DocMeta",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "获取表格元数据",
				Description: `
该接口用于根据 spreadsheetToken 获取表格元数据。
`,
				Request:     "GET https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/metainfo",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uETMzUjLxEzM14SMxMTN",
				FuncName:    "SpreadSheetsMetainfo",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "操作子表",
				Description: `

该接口用于根据 spreadsheetToken 操作表格，如增加sheet，复制sheet、删除sheet。
`,
				Request:     "POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/sheets_batch_update",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYTMzUjL2EzM14iNxMTN",
				FuncName:    "SheetsBatchUpdate",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "插入数据",
				Description: `

该接口用于根据 spreadsheetToken 和 range 向范围之前增加相应数据的行和相应的数据，相当于数组的插入操作；单次写入不超过5000行，100列，每个格子大小为0.5M。
`,
				Request:     "POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_prepend",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uIjMzUjLyIzM14iMyMTN",
				FuncName:    "SpreadSheetsValuesPrepend",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "追加数据",
				Description: `

该接口用于根据 spreadsheetToken 和 range 遇到空行则进行覆盖追加或新增行追加数据。 空行：默认该行第一个格子是空，则认为是空行；单次写入不超过5000行，100列，每个格子大小为0.5M。
`,
				Request:     "POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_append",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMjMzUjLzIzM14yMyMTN",
				FuncName:    "SpreadSheetsValuesAppend",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "插入行列",
				Description: `

该接口用于根据 spreadsheetToken 和维度信息 插入空行/列 。如 startIndex=3， endIndex=7，则从第 4 行开始开始插入行列，一直到第 7 行，共插入 4 行；单次操作不超过5000行或列。
`,
				Request:     "POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/insert_dimension_range",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQjMzUjL0IzM14CNyMTN",
				FuncName:    "SpreadSheetsInsertDimensionRange",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "增加行列",
				Description: `

该接口用于根据 spreadsheetToken 和长度，在末尾增加空行/列；单次操作不超过5000行或列。
`,
				Request:     "POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/dimension_range",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uUjMzUjL1IzM14SNyMTN",
				FuncName:    "SpreadSheetsDimensionRange",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "增加锁定单元格",
				Description: `

该接口用于根据 spreadsheetToken 和维度信息增加多个范围的锁定单元格；单次操作不超过5000行或列。
`,
				Request:     "POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/protected_dimension",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ugDNzUjL4QzM14CO0MTN",
				FuncName:    "SpreadSheetsProtectedDimension",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "合并单元格",
				Description: `

该接口用于根据 spreadsheetToken 和维度信息合并单元格；单次操作不超过5000行，100列。
`,
				Request:     "POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/merge_cells",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukDNzUjL5QzM14SO0MTN",
				FuncName:    "SpreadSheetsMergeCells",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "拆分单元格",
				Description: `

该接口用于根据 spreadsheetToken 和维度信息拆分单元格；单次操作不超过5000行，100列。
`,
				Request:     "POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/unmerge_cells",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uATNzUjLwUzM14CM1MTN",
				FuncName:    "SpreadSheetsUnmergeCells",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "读取单个范围",
				Description: `

该接口用于根据 spreadsheetToken 和 range 读取表格单个范围的值，返回数据限制为10M。
`,
				Request:     "GET https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values/:range",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ugTMzUjL4EzM14COxMTN",
				FuncName:    "SpreadSheetsRange",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "读取多个范围",
				Description: `

该接口用于根据 spreadsheetToken 和 ranges 读取表格多个范围的值，返回数据限制为10M。
`,
				Request:     "GET https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_batch_get",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukTMzUjL5EzM14SOxMTN",
				FuncName:    "SpreadSheetsValuesBatchGet",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "向多个范围写入数据",
				Description: `

该接口用于根据 spreadsheetToken 和 range 向多个范围写入数据，若范围内有数据，将被更新覆盖；单次写入不超过5000行，100列，每个格子大小为0.5M。
`,
				Request:     "POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_batch_update",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uEjMzUjLxIzM14SMyMTN",
				FuncName:    "SpreadSheetsValuesBatchUpdate",
				AccessToken: "user",
				PathParams: []Param{
					{Name: "", Type: "string"},
				},
			},
			{
				Name: "获取元数据",
				Description: `

该接口用于根据 token 获取各类文件的元数据
`,
				Request:     "POST https://open.feishu.cn/open-apis/suite/docs-api/meta",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMjN3UjLzYzN14yM2cTN",
				FuncName:    "DocsMeta",
				AccessToken: "user",
			},
			{
				Name: "文档搜索",
				Description: `

该接口用于根据搜索条件进行文档搜索
`,
				Request:     "POST https://open.feishu.cn/open-apis/suite/docs-api/search/object",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ugDM4UjL4ADO14COwgTN",
				FuncName:    "SearchObject",
				AccessToken: "user",
			},
			{
				Name: "添加全文评论",
				Description: `

该接口用于根据 filetoken 给文档添加全文评论
`,
				Request:     "POST https://open.feishu.cn/open-apis/comment/add_whole",
				See:         "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ucDN4UjL3QDO14yN0gTN",
				FuncName:    "CommentAddWhole",
				AccessToken: "user",
			},
		},
	},
	{
		Name:    `审批`,
		Package: `capabilities/approve`,
		Apis: []Api{

			{
				Name: "查看审批定义",
				Description: `
根据 Approval Code 获取某个审批定义的详情，用于构造创建审批实例的请求。

**权限说明** ：需要获取 访问审批应用 权限。
`,
				Request:  "POST https://www.feishu.cn/approval/openapi/v2/approval/get",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uADNyUjLwQjM14CM0ITN",
				FuncName: "Get",
			},
			{
				Name: "批量获取审批实例ID",
				Description: `
根据 approval_code 批量获取审批实例的 instance_code，用于拉取租户下某个审批定义的全部审批实例。
默认以审批创建时间排序。

**权限说明** ：需要获取 访问审批应用 权限。
`,
				Request:  "POST https://www.feishu.cn/approval/openapi/v2/instance/list",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQDOyUjL0gjM14CN4ITN",
				FuncName: "InstanceList",
			},
			{
				Name: "获取单个审批实例详情",
				Description: `
通过审批实例 Instance Code  获取审批实例详情。Instance Code 由 [批量获取审批实例](https://open.feishu.cn/document/ukTMukTMukTM/uQDOyUjL0gjM14CN4ITN) 接口获取。
`,
				Request:  "POST https://www.feishu.cn/approval/openapi/v2/instance/get",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uEDNyUjLxQjM14SM0ITN",
				FuncName: "InstanceGet",
			},
			{
				Name: "创建审批实例",
				Description: `
创建一个审批实例，调用方需对审批定义的表单有详细了解，将按照定义的表单结构，将表单 Value 通过接口传入。
`,
				Request:  "POST https://www.feishu.cn/approval/openapi/v2/instance/create",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uIDNyUjLyQjM14iM0ITN",
				FuncName: "InstanceCreate",
			},
			{
				Name: "审批任务同意",
				Description: `
对于单个审批任务进行同意操作。同意后审批流程会流转到下一个审批人。


**权限说明** ：需要获取 访问审批应用 权限。
`,
				Request:  "POST https://www.feishu.cn/approval/openapi/v2/instance/approve",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMDNyUjLzQjM14yM0ITN",
				FuncName: "Approve",
			},
			{
				Name: "审批任务拒绝",
				Description: `
对于单个审批任务进行拒绝操作。拒绝后审批流程结束。

**权限说明** ：需要获取 访问审批应用 权限。
`,
				Request:  "POST https://www.feishu.cn/approval/openapi/v2/instance/reject",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQDNyUjL0QjM14CN0ITN",
				FuncName: "Reject",
			},
			{
				Name: "审批任务转交",
				Description: `
对于单个审批任务进行转交操作。转交后审批流程流转给被转交人。

**权限说明** ：需要获取 访问审批应用 权限。
`,
				Request:  "POST https://www.feishu.cn/approval/openapi/v2/instance/transfer",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uUDNyUjL1QjM14SN0ITN",
				FuncName: "Transfer",
			},
			{
				Name: "审批实例撤回",
				Description: `
对于单个审批实例进行撤销操作。撤销后审批流程结束。

**权限说明** ：需要获取 访问审批应用 权限。
`,
				Request:  "POST https://www.feishu.cn/approval/openapi/v2/instance/cancel",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYDNyUjL2QjM14iN0ITN",
				FuncName: "Cancel",
			},
			{
				Name: "上传文件",
				Description: `
当审批表单中有图片或附件控件时，开发者需在创建审批实例前通过审批上传文件接口将文件上传到审批系统。

**权限说明** ：需要获取 访问审批应用 权限。
`,
				Request:  "POST https://www.feishu.cn/approval/openapi/v2/file/upload",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uUDOyUjL1gjM14SN4ITN",
				FuncName: "Upload",
			},
			{
				Name: "三方审批实例同步",
				Description: `
用于第三方审批的实例同步。
使用前需确保已经在审批后台创建第三方审批。

> 审批实例：员工发起审批时产生的对象，详情参见 [开发指南](/ssl:ttdoc/ugTM5UjL4ETO14COxkTN/ukDNyUjL5QjM14SO0ITN)
> 
**权限说明** ：需要获取 访问审批应用 权限。
`,
				Request:  "POST https://www.feishu.cn/approval/openapi/v2/external/instance/create",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uczM3UjL3MzN14yNzcTN",
				FuncName: "ExternalInstanceCreate",
			},
			{
				Name: "创建审批定义",
				Description: `
用于通过接口创建简单的审批定义，可以灵活指定定义的基础信息、表单和流程等。不推荐企业自建应用使用，如有需要尽量联系管理员在审批管理后台创建定义。

**权限说明** ：需要获取 访问审批应用 权限。
`,
				Request:  "POST https://www.feishu.cn/approval/openapi/v2/approval/create",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uUzNyYjL1cjM24SN3IjN",
				FuncName: "Create",
			},
			{
				Name: "审批实例抄送",
				Description: `
通过接口可以将当前审批实例抄送给其他人。

**权限说明** ：需要获取 访问审批应用 权限。
`,
				Request:  "POST https://www.feishu.cn/approval/openapi/v2/instance/cc",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uADOzYjLwgzM24CM4MjN",
				FuncName: "InstanceCc",
			},
			{
				Name: "订阅审批事件",
				Description: `
订阅 approval_code 后，可以收到该审批定义对应实例的事件通知。
`,
				Request:  "POST https://www.feishu.cn/approval/openapi/v2/subscription/subscribe",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ucDOyUjL3gjM14yN4ITN",
				FuncName: "Subscribe",
			},
			{
				Name: "取消订阅审批事件",
				Description: `
取消订阅 approval_code 后，无法再收到该审批定义对应实例的事件通知。
`,
				Request:  "POST https://www.feishu.cn/approval/openapi/v2/subscription/unsubscribe",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ugDOyUjL4gjM14CO4ITN",
				FuncName: "Unsubscribe",
			},
		},
	},
	{
		Name:    `会议室`,
		Package: `capabilities/meeting`,
		Apis: []Api{

			{
				Name: "获取建筑物列表",
				Description: `
该接口用于获取本企业下的建筑物（办公大楼）。

**权限说明** ：需要获取 获取会议室信息 权限。
`,
				Request:  "GET https://open.feishu.cn/open-apis/meeting_room/building/list?page_size=1&page_token=0&order_by=name-asc&fields=*",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ugzNyUjL4cjM14CO3ITN",
				FuncName: "BuildingList",
				GetParams: []Param{
					{Name: `page_size`, Type: `string`},
					{Name: `page_token`, Type: `string`},
					{Name: `order_by`, Type: `string`},
					{Name: `fields`, Type: `string`},
				},
			},
			{
				Name: "查询建筑物详情",
				Description: `
该接口用于获取指定建筑物的详细信息。

**权限说明** ：需要获取 获取会议室信息 权限。
`,
				Request:  "GET https://open.feishu.cn/open-apis/meeting_room/building/batch_get?building_ids=omb_8ec170b937536a5d87c23b418b83f9bb&building_ids=omb_38570e4f0fd9ecf15030d3cc8b388f3a&fields=*",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukzNyUjL5cjM14SO3ITN",
				FuncName: "BuildingBatchGet",
				GetParams: []Param{
					{Name: `building_ids`, Type: `string`},
					{Name: `fields`, Type: `string`},
				},
			},
			{
				Name: "获取会议室列表",
				Description: `
该接口用于获取指定建筑下的会议室。

**权限说明** ：需要获取 获取会议室信息 权限。
`,
				Request:  "GET https://open.feishu.cn/open-apis/meeting_room/room/list?building_id=omb_8ec170b937536a5d87c23b418b83f9bb&page_size=1&page_token=0&order_by=name-asc&fields=*",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uADOyUjLwgjM14CM4ITN",
				FuncName: "MeetingRoomList",
				GetParams: []Param{
					{Name: `building_id`, Type: `string`},
					{Name: `page_size`, Type: `string`},
					{Name: `page_token`, Type: `string`},
					{Name: `order_by`, Type: `string`},
					{Name: `fields`, Type: `string`},
				},
			},
			{
				Name: "查询会议室详情",
				Description: `
该接口用于获取指定会议室的详细信息。

**权限说明** ：需要获取 获取会议室信息 权限。
`,
				Request:  "GET https://open.feishu.cn/open-apis/meeting_room/room/batch_get?room_ids=omm_eada1d61a550955240c28757e7dec3af&room_ids=omm_83d09ad4f6896e02029a6a075f71c9d1&fields=*",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uEDOyUjLxgjM14SM4ITN",
				FuncName: "MeetingRoomBatchGet",
				GetParams: []Param{
					{Name: `room_ids`, Type: `string`},
					{Name: `fields`, Type: `string`},
				},
			},
			{
				Name: "会议室忙闲查询",
				Description: `
该接口用于获取指定会议室的忙闲日程实例列表。非重复日程只有唯一实例；重复日程可能存在多个实例，依据重复规则和时间范围扩展。

**权限说明** ：需要获取 获取会议室信息 权限。
`,
				Request:  "GET https://open.feishu.cn/open-apis/meeting_room/freebusy/batch_get?room_ids=omm_83d09ad4f6896e02029a6a075f71c9d1&room_ids=omm_eada1d61a550955240c28757e7dec3af&time_min=2019-09-04T08:45:00%2B08:00&time_max=2019-09-04T09:45:00%2B08:00",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uIDOyUjLygjM14iM4ITN",
				FuncName: "MeetingRoomFreeBusyBatchGet",
				GetParams: []Param{
					{Name: `room_ids`, Type: `string`},
					{Name: `time_min`, Type: `string`},
					{Name: `time_max`, Type: `string`},
				},
			},
			{
				Name: "回复会议室日程实例",
				Description: `
该接口用于回复会议室日程实例，包括未签到释放和提前结束释放。

**权限说明** ：需要获取 获取会议室信息 权限。
`,
				Request:  "POST https://open.feishu.cn/open-apis/meeting_room/instance/reply",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYzN4UjL2cDO14iN3gTN",
				FuncName: "InstanceReply",
			},
			{
				Name: "创建建筑物",
				Description: `
该接口对应管理后台的添加建筑，添加楼层的功能，可用于创建建筑物和建筑物的楼层信息。

**权限说明** ：需要获取 管理会议室信息 权限。
`,
				Request:  "POST https://open.feishu.cn/open-apis/meeting_room/building/create",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uATNwYjLwUDM24CM1AjN",
				FuncName: "BuildingCreate",
			},
			{
				Name: "更新建筑物",
				Description: `
该接口用于编辑建筑信息，添加楼层，删除楼层，编辑楼层信息。

**权限说明** ：需要获取 管理会议室信息 权限。
`,
				Request:  "POST https://open.feishu.cn/open-apis/meeting_room/building/update",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uETNwYjLxUDM24SM1AjN",
				FuncName: "BuildingUpdate",
			},
			{
				Name: "删除建筑物",
				Description: `
该接口用于删除建筑物（办公大楼）。

**权限说明** ：需要获取 管理会议室信息 权限。
`,
				Request:  "POST https://open.feishu.cn/open-apis/meeting_room/building/delete",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMzMxYjLzMTM24yMzEjN",
				FuncName: "BuildingDelete",
			},
			{
				Name: "查询建筑物ID",
				Description: `
该接口用于根据租户自定义建筑 ID 查询建筑 ID。

**权限说明** ：需要获取 获取会议室信息 权限。
`,
				Request:  "GET https://open.feishu.cn/open-apis/meeting_room/building/batch_get_id?custom_building_ids=test01&custom_building_ids=test02",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQzMxYjL0MTM24CNzEjN",
				FuncName: "BuildingBatchGetById",
				GetParams: []Param{
					{Name: `custom_building_ids`, Type: `string`},
				},
			},
			{
				Name: "创建会议室",
				Description: `
该接口用于创建会议室。

**权限说明** ：需要获取 管理会议室信息 权限。
`,
				Request:  "POST https://open.feishu.cn/open-apis/meeting_room/room/create",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uITNwYjLyUDM24iM1AjN",
				FuncName: "MeetingRoomCreate",
			},
			{
				Name: "更新会议室",
				Description: `
该接口用于更新会议室。

**权限说明** ：需要获取 管理会议室信息 权限。
`,
				Request:  "POST https://open.feishu.cn/open-apis/meeting_room/room/update",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMTNwYjLzUDM24yM1AjN",
				FuncName: "MeetingRoomUpdate",
			},
			{
				Name: "删除会议室",
				Description: `
该接口用于删除会议室。

**权限说明** ：需要获取 管理会议室信息 权限。
`,
				Request:  "POST https://open.feishu.cn/open-apis/meeting_room/room/delete",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uUzMxYjL1MTM24SNzEjN",
				FuncName: "MeetingRoomDelete",
			},
			{
				Name: "查询会议室ID",
				Description: `
该接口用于根据租户自定义会议室ID查询会议室ID。

**权限说明** ：需要获取 获取会议室信息 权限。
`,
				Request:  "GET https://open.feishu.cn/open-apis/meeting_room/room/batch_get_id?custom_room_ids=test01&custom_room_ids=test02",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYzMxYjL2MTM24iNzEjN",
				FuncName: "MeetingRoomBatchGetById",
				GetParams: []Param{
					{Name: `custom_room_ids`, Type: `string`},
				},
			},
			{
				Name: "获取国家地区列表",
				Description: `
新建建筑时需要标明所处国家/地区，该接口用于获得系统预先提供的可供选择的国家 /地区列表

**权限说明** ：需要获取 获取会议室信息 权限。
`,
				Request:  "GET https://open.feishu.cn/open-apis/meeting_room/country/list",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQTNwYjL0UDM24CN1AjN",
				FuncName: "CountryList",
			},
			{
				Name: "获取城市列表",
				Description: `
新建建筑时需要选择所处国家/地区，该接口用于获得系统预先提供的可供选择的城市列表

**权限说明** ：需要获取 获取会议室信息 权限。
`,
				Request:  "GET https://open.feishu.cn/open-apis/meeting_room/district/list?country_id=1814991",
				See:      "https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uUTNwYjL1UDM24SN1AjN",
				FuncName: "DistrictList",
				GetParams: []Param{
					{Name: `country_id`, Type: `string`},
				},
			},
		},
	},
}
