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

// Package document 云文档
package document

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"

	"github.com/fastwego/feishu"
)

const (
	apiCreateFolder                     = "/open-apis/drive/explorer/v2/folder/:folderToken"
	apiRootFolderMeta                   = "/open-apis/drive/explorer/v2/root_folder/meta"
	apiFolderChildren                   = "/open-apis/drive/explorer/v2/folder/:folderToken/children"
	apiCreateFile                       = "/open-apis/drive/explorer/v2/file/:folderToken"
	apiCopyFile                         = "/open-apis/drive/explorer/v2/file/copy/files/:fileToken"
	apiDeleteFile                       = "/open-apis/drive/explorer/v2/file/spreadsheets/:spreadsheetToken"
	apiPermissionMemberCreate           = "/open-apis/drive/permission/member/create"
	apiPermissionMemberTransfer         = "/open-apis/drive/permission/member/transfer"
	apiPermissionPublicUpdate           = "/open-apis/drive/permission/public/update"
	apiPermissionMemberList             = "/open-apis/drive/permission/member/list"
	apiPermissionMemberDelete           = "/open-apis/drive/permission/member/delete"
	apiPermissionMemberUpdate           = "/open-apis/drive/permission/member/update"
	apiPermissionMemberPermitted        = "/open-apis/drive/permission/member/permitted"
	apiRawContent                       = "/open-apis/doc/v2/:docToken/raw_content"
	apiSheetMeta                        = "/open-apis/doc/v2/:docToken/sheet_meta"
	apiDocMeta                          = "/open-apis/doc/v2/meta/:docToken"
	apiSpreadSheetsMetainfo             = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/metainfo"
	apiSheetsBatchUpdate                = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/sheets_batch_update"
	apiSpreadSheetsValuesPrepend        = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_prepend"
	apiSpreadSheetsValuesAppend         = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_append"
	apiSpreadSheetsInsertDimensionRange = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/insert_dimension_range"
	apiSpreadSheetsDimensionRange       = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/dimension_range"
	apiSpreadSheetsProtectedDimension   = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/protected_dimension"
	apiSpreadSheetsMergeCells           = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/merge_cells"
	apiSpreadSheetsUnmergeCells         = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/unmerge_cells"
	apiSpreadSheetsRange                = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values/:range"
	apiSpreadSheetsValuesBatchGet       = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_batch_get"
	apiSpreadSheetsValuesBatchUpdate    = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_batch_update"
	apiDocsMeta                         = "/open-apis/suite/docs-api/meta"
	apiSearchObject                     = "/open-apis/suite/docs-api/search/object"
	apiCommentAddWhole                  = "/open-apis/comment/add_whole"
)

/*
新建文件夹



该接口用于根据 folderToken 在该 folder 下创建文件夹


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukTNzUjL5UzM14SO1MTN

POST https://open.feishu.cn/open-apis/drive/explorer/v2/folder/:folderToken
*/
func CreateFolder(ctx *feishu.App, payload []byte, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiCreateFolder
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(api, bytes.NewReader(payload), header)
}

/*
获取文件夹元信息


# 获取root folder（我的空间） meta该接口用于获取 我的空间 的元信息


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uAjNzUjLwYzM14CM2MTN

GET https://open.feishu.cn/open-apis/drive/explorer/v2/root_folder/meta
*/
func RootFolderMeta(ctx *feishu.App, header http.Header) (resp []byte, err error) {

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(apiRootFolderMeta, header)
}

/*
获取文件夹下的文档清单



该接口用于根据 folderToken 获取该文件夹的文档清单，如 doc、sheet、folder


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uEjNzUjLxYzM14SM2MTN

GET https://open.feishu.cn/open-apis/drive/explorer/v2/folder/:folderToken/children
*/
func FolderChildren(ctx *feishu.App, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiFolderChildren
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(api, header)
}

/*
新建文档


该接口用于根据 folderToken 创建 Doc或 Sheet 。

若没有特定的文件夹用于承载创建的文档，可以先调用「获取文件夹元信息」文档中的「获取 root folder (我的空间) meta」接口，获得我的空间的 token，然后再使用此接口。创建的文档将会在「我的空间」的「归我所有」列表里。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQTNzUjL0UzM14CN1MTN

POST https://open.feishu.cn/open-apis/drive/explorer/v2/file/:folderToken
*/
func CreateFile(ctx *feishu.App, payload []byte, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiCreateFile
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(api, bytes.NewReader(payload), header)
}

/*
复制文档



该接口用于根据文件 token 复制 Doc 或 Sheet  到目标文件夹中。

若没有特定的文件夹用于承载创建的文档，可以先调用「获取文件夹元信息」文档中的「获取 root folder (我的空间) meta」接口，获得我的空间的 token，然后再使用此接口。复制的文档将会在「我的空间」的「归我所有」列表里。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYTNzUjL2UzM14iN1MTN

POST https://open.feishu.cn/open-apis/drive/explorer/v2/file/copy/files/:fileToken
*/
func CopyFile(ctx *feishu.App, payload []byte, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiCopyFile
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(api, bytes.NewReader(payload), header)
}

/*
删除文档

Doc

该接口用于根据 spreadsheetToken 删除对应的 sheet 文档。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uATM2UjLwEjN14CMxYTN

DELETE https://open.feishu.cn/open-apis/drive/explorer/v2/file/spreadsheets/:spreadsheetToken
*/
func DeleteFile(ctx *feishu.App, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiDeleteFile
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPDelete(api, header)
}

/*
增加权限



该接口用于根据 filetoken 给用户增加文档的权限


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMzNzUjLzczM14yM3MTN

POST https://open.feishu.cn/open-apis/drive/permission/member/create
*/
func PermissionMemberCreate(ctx *feishu.App, payload []byte, header http.Header) (resp []byte, err error) {

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiPermissionMemberCreate, bytes.NewReader(payload), header)
}

/*
转移拥有者



该接口用于根据文档信息和用户信息转移文档的所有者。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQzNzUjL0czM14CN3MTN

POST https://open.feishu.cn/open-apis/drive/permission/member/transfer
*/
func PermissionMemberTransfer(ctx *feishu.App, payload []byte, header http.Header) (resp []byte, err error) {

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiPermissionMemberTransfer, bytes.NewReader(payload), header)
}

/*
更新文档公共设置



该接口用于根据 filetoken 更新文档的公共设置


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukTM3UjL5EzN14SOxcTN

POST https://open.feishu.cn/open-apis/drive/permission/public/update
*/
func PermissionPublicUpdate(ctx *feishu.App, payload []byte, header http.Header) (resp []byte, err error) {

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiPermissionPublicUpdate, bytes.NewReader(payload), header)
}

/*
获取协作者列表



该接口用于根据 filetoken 查询协作者，目前包括人('user')和群('chat')

**你能获取到协作者列表的前提是你对该文档有权限**


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uATN3UjLwUzN14CM1cTN

POST https://open.feishu.cn/open-apis/drive/permission/member/list
*/
func PermissionMemberList(ctx *feishu.App, payload []byte, header http.Header) (resp []byte, err error) {

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiPermissionMemberList, bytes.NewReader(payload), header)
}

/*
移除协作者权限



该接口用于根据 filetoken 移除文档协作者的权限


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYTN3UjL2UzN14iN1cTN

POST https://open.feishu.cn/open-apis/drive/permission/member/delete
*/
func PermissionMemberDelete(ctx *feishu.App, payload []byte, header http.Header) (resp []byte, err error) {

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiPermissionMemberDelete, bytes.NewReader(payload), header)
}

/*
更新协作者权限



该接口用于根据 filetoken 更新文档协作者的权限


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ucTN3UjL3UzN14yN1cTN

POST https://open.feishu.cn/open-apis/drive/permission/member/update
*/
func PermissionMemberUpdate(ctx *feishu.App, payload []byte, header http.Header) (resp []byte, err error) {

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiPermissionMemberUpdate, bytes.NewReader(payload), header)
}

/*
判断协作者是否有某权限



该接口用于根据 filetoken 判断当前登录用户是否具有某权限


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYzN3UjL2czN14iN3cTN

POST https://open.feishu.cn/open-apis/drive/permission/member/permitted
*/
func PermissionMemberPermitted(ctx *feishu.App, payload []byte, header http.Header) (resp []byte, err error) {

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiPermissionMemberPermitted, bytes.NewReader(payload), header)
}

/*
获取文档文本内容


该接口用于获取文档的纯文本内容，不包含富文本格式信息，主要用于搜索，如导入 es 等。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukzNzUjL5czM14SO3MTN

GET https://open.feishu.cn/open-apis/doc/v2/:docToken/raw_content
*/
func RawContent(ctx *feishu.App, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiRawContent
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(api, header)
}

/*
获取sheet@doc的元数据

sheet@doc 的元数据

该接口用于根据 docToken 获取 sheet@doc 的元数据。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uADOzUjLwgzM14CM4MTN

GET https://open.feishu.cn/open-apis/doc/v2/:docToken/sheet_meta
*/
func SheetMeta(ctx *feishu.App, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiSheetMeta
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(api, header)
}

/*
获取文档元信息


该接口用于根据 docToken 获取元数据。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uczN3UjL3czN14yN3cTN

GET https://open.feishu.cn/open-apis/doc/v2/meta/:docToken
*/
func DocMeta(ctx *feishu.App, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiDocMeta
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(api, header)
}

/*
获取表格元数据


该接口用于根据 spreadsheetToken 获取表格元数据。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uETMzUjLxEzM14SMxMTN

GET https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/metainfo
*/
func SpreadSheetsMetainfo(ctx *feishu.App, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiSpreadSheetsMetainfo
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(api, header)
}

/*
操作子表



该接口用于根据 spreadsheetToken 操作表格，如增加sheet，复制sheet、删除sheet。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYTMzUjL2EzM14iNxMTN

POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/sheets_batch_update
*/
func SheetsBatchUpdate(ctx *feishu.App, payload []byte, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiSheetsBatchUpdate
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(api, bytes.NewReader(payload), header)
}

/*
插入数据



该接口用于根据 spreadsheetToken 和 range 向范围之前增加相应数据的行和相应的数据，相当于数组的插入操作；单次写入不超过5000行，100列，每个格子大小为0.5M。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uIjMzUjLyIzM14iMyMTN

POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_prepend
*/
func SpreadSheetsValuesPrepend(ctx *feishu.App, payload []byte, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiSpreadSheetsValuesPrepend
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(api, bytes.NewReader(payload), header)
}

/*
追加数据



该接口用于根据 spreadsheetToken 和 range 遇到空行则进行覆盖追加或新增行追加数据。 空行：默认该行第一个格子是空，则认为是空行；单次写入不超过5000行，100列，每个格子大小为0.5M。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMjMzUjLzIzM14yMyMTN

POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_append
*/
func SpreadSheetsValuesAppend(ctx *feishu.App, payload []byte, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiSpreadSheetsValuesAppend
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(api, bytes.NewReader(payload), header)
}

/*
插入行列



该接口用于根据 spreadsheetToken 和维度信息 插入空行/列 。如 startIndex=3， endIndex=7，则从第 4 行开始开始插入行列，一直到第 7 行，共插入 4 行；单次操作不超过5000行或列。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQjMzUjL0IzM14CNyMTN

POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/insert_dimension_range
*/
func SpreadSheetsInsertDimensionRange(ctx *feishu.App, payload []byte, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiSpreadSheetsInsertDimensionRange
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(api, bytes.NewReader(payload), header)
}

/*
增加行列



该接口用于根据 spreadsheetToken 和长度，在末尾增加空行/列；单次操作不超过5000行或列。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uUjMzUjL1IzM14SNyMTN

POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/dimension_range
*/
func SpreadSheetsDimensionRange(ctx *feishu.App, payload []byte, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiSpreadSheetsDimensionRange
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(api, bytes.NewReader(payload), header)
}

/*
增加锁定单元格



该接口用于根据 spreadsheetToken 和维度信息增加多个范围的锁定单元格；单次操作不超过5000行或列。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ugDNzUjL4QzM14CO0MTN

POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/protected_dimension
*/
func SpreadSheetsProtectedDimension(ctx *feishu.App, payload []byte, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiSpreadSheetsProtectedDimension
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(api, bytes.NewReader(payload), header)
}

/*
合并单元格



该接口用于根据 spreadsheetToken 和维度信息合并单元格；单次操作不超过5000行，100列。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukDNzUjL5QzM14SO0MTN

POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/merge_cells
*/
func SpreadSheetsMergeCells(ctx *feishu.App, payload []byte, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiSpreadSheetsMergeCells
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(api, bytes.NewReader(payload), header)
}

/*
拆分单元格



该接口用于根据 spreadsheetToken 和维度信息拆分单元格；单次操作不超过5000行，100列。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uATNzUjLwUzM14CM1MTN

POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/unmerge_cells
*/
func SpreadSheetsUnmergeCells(ctx *feishu.App, payload []byte, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiSpreadSheetsUnmergeCells
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(api, bytes.NewReader(payload), header)
}

/*
读取单个范围



该接口用于根据 spreadsheetToken 和 range 读取表格单个范围的值，返回数据限制为10M。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ugTMzUjL4EzM14COxMTN

GET https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values/:range
*/
func SpreadSheetsRange(ctx *feishu.App, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiSpreadSheetsRange
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(api, header)
}

/*
读取多个范围



该接口用于根据 spreadsheetToken 和 ranges 读取表格多个范围的值，返回数据限制为10M。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukTMzUjL5EzM14SOxMTN

GET https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_batch_get
*/
func SpreadSheetsValuesBatchGet(ctx *feishu.App, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiSpreadSheetsValuesBatchGet
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPGet(api, header)
}

/*
向多个范围写入数据



该接口用于根据 spreadsheetToken 和 range 向多个范围写入数据，若范围内有数据，将被更新覆盖；单次写入不超过5000行，100列，每个格子大小为0.5M。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uEjMzUjLxIzM14SMyMTN

POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_batch_update
*/
func SpreadSheetsValuesBatchUpdate(ctx *feishu.App, payload []byte, params url.Values, header http.Header) (resp []byte, err error) {

	api := apiSpreadSheetsValuesBatchUpdate
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(api, bytes.NewReader(payload), header)
}

/*
获取元数据



该接口用于根据 token 获取各类文件的元数据


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMjN3UjLzYzN14yM2cTN

POST https://open.feishu.cn/open-apis/suite/docs-api/meta
*/
func DocsMeta(ctx *feishu.App, payload []byte, header http.Header) (resp []byte, err error) {

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiDocsMeta, bytes.NewReader(payload), header)
}

/*
文档搜索



该接口用于根据搜索条件进行文档搜索


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ugDM4UjL4ADO14COwgTN

POST https://open.feishu.cn/open-apis/suite/docs-api/search/object
*/
func SearchObject(ctx *feishu.App, payload []byte, header http.Header) (resp []byte, err error) {

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiSearchObject, bytes.NewReader(payload), header)
}

/*
添加全文评论



该接口用于根据 filetoken 给文档添加全文评论


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ucDN4UjL3QDO14yN0gTN

POST https://open.feishu.cn/open-apis/comment/add_whole
*/
func CommentAddWhole(ctx *feishu.App, payload []byte, header http.Header) (resp []byte, err error) {

	header.Set("Content-appType", "application/json")

	return ctx.Client.HTTPPost(apiCommentAddWhole, bytes.NewReader(payload), header)
}
