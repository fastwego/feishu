// Package sheets 云文档/sheets
package sheets

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"

	"github.com/fastwego/feishu"
)

const (
	apiMetaInfo             = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/metainfo"
	apiUpdateProperties     = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/properties"
	apiBatchUpdate          = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/sheets_batch_update"
	apiValuesPrepend        = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_prepend"
	apiValuesAppend         = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_append"
	apiInsertDimensionRange = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/insert_dimension_range"
	apiCreateDimensionRange = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/dimension_range"
	apiUpdateDimensionRange = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/dimension_range"
	apiDeleteDimensionRange = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/dimension_range"
	apiUpdateStyle          = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/style"
	apiBatchUpdateStyle     = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/styles_batch_update"
	apiProtectedDimension   = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/protected_dimension"
	apiMergeCells           = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/merge_cells"
	apiUnmergeCells         = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/unmerge_cells"
	apiRange                = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values/:range"
	apiValues               = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values"
	apiValuesBatchGet       = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_batch_get"
	apiValuesBatchUpdate    = "/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_batch_update"
)

/*
获取表格元数据


该接口用于根据 spreadsheetToken 获取表格元数据。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uETMzUjLxEzM14SMxMTN

GET https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/metainfo
*/
func MetaInfo(ctx *feishu.App, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiMetaInfo
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+api, header)
}

/*
更新表格属性

该接口用于根据 spreadsheetToken 更新表格属性，如更新表格标题。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ucTMzUjL3EzM14yNxMTN

PUT https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/properties
*/
func UpdateProperties(ctx *feishu.App, payload []byte, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiUpdateProperties
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPut(feishu.FeishuServerUrl+api, bytes.NewReader(payload), header)
}

/*
操作子表/更新子表属性



该接口用于根据 spreadsheetToken 操作表格，如增加sheet，复制sheet、删除sheet。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYTMzUjL2EzM14iNxMTN

POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/sheets_batch_update
*/
func BatchUpdate(ctx *feishu.App, payload []byte, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiBatchUpdate
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+api, bytes.NewReader(payload), header)
}

/*
插入数据



该接口用于根据 spreadsheetToken 和 range 向范围之前增加相应数据的行和相应的数据，相当于数组的插入操作；单次写入不超过5000行，100列，每个格子大小为0.5M。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uIjMzUjLyIzM14iMyMTN

POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_prepend
*/
func ValuesPrepend(ctx *feishu.App, payload []byte, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiValuesPrepend
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+api, bytes.NewReader(payload), header)
}

/*
追加数据



该接口用于根据 spreadsheetToken 和 range 遇到空行则进行覆盖追加或新增行追加数据。 空行：默认该行第一个格子是空，则认为是空行；单次写入不超过5000行，100列，每个格子大小为0.5M。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMjMzUjLzIzM14yMyMTN

POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_append
*/
func ValuesAppend(ctx *feishu.App, payload []byte, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiValuesAppend
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+api, bytes.NewReader(payload), header)
}

/*
插入行列



该接口用于根据 spreadsheetToken 和维度信息 插入空行/列 。如 startIndex=3， endIndex=7，则从第 4 行开始开始插入行列，一直到第 7 行，共插入 4 行；单次操作不超过5000行或列。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQjMzUjL0IzM14CNyMTN

POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/insert_dimension_range
*/
func InsertDimensionRange(ctx *feishu.App, payload []byte, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiInsertDimensionRange
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+api, bytes.NewReader(payload), header)
}

/*
增加行列



该接口用于根据 spreadsheetToken 和长度，在末尾增加空行/列；单次操作不超过5000行或列。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uUjMzUjL1IzM14SNyMTN

POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/dimension_range
*/
func CreateDimensionRange(ctx *feishu.App, payload []byte, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiCreateDimensionRange
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+api, bytes.NewReader(payload), header)
}

/*
更新行列

该接口用于根据 spreadsheetToken 和维度信息更新隐藏行列、单元格大小；单次操作不超过5000行或列。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYjMzUjL2IzM14iNyMTN

PUT https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/dimension_range
*/
func UpdateDimensionRange(ctx *feishu.App, payload []byte, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiUpdateDimensionRange
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPut(feishu.FeishuServerUrl+api, bytes.NewReader(payload), header)
}

/*
删除行列

该接口用于根据 spreadsheetToken 和维度信息删除行/列 。单次删除最大5000行/列。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ucjMzUjL3IzM14yNyMTN

DELETE https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/dimension_range
*/
func DeleteDimensionRange(ctx *feishu.App, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiDeleteDimensionRange
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPDelete(feishu.FeishuServerUrl+api, header)
}

/*
设置单元格样式

该接口用于根据 spreadsheetToken 、range 和样式信息更新单元格样式；单次写入不超过5000行，100列。

See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukjMzUjL5IzM14SOyMTN

PUT https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/style
*/
func UpdateStyle(ctx *feishu.App, payload []byte, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiUpdateStyle
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPut(feishu.FeishuServerUrl+api, bytes.NewReader(payload), header)
}

/*
批量设置单元格样式

该接口用于根据 spreadsheetToken 、range和样式信息 批量更新单元格样式；单次写入不超过5000行，100列。

See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uAzMzUjLwMzM14CMzMTN

PUT https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/styles_batch_update
*/
func BatchUpdateStyle(ctx *feishu.App, payload []byte, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiBatchUpdateStyle
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPut(feishu.FeishuServerUrl+api, bytes.NewReader(payload), header)
}

/*
增加锁定单元格



该接口用于根据 spreadsheetToken 和维度信息增加多个范围的锁定单元格；单次操作不超过5000行或列。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ugDNzUjL4QzM14CO0MTN

POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/protected_dimension
*/
func ProtectedDimension(ctx *feishu.App, payload []byte, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiProtectedDimension
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+api, bytes.NewReader(payload), header)
}

/*
合并单元格



该接口用于根据 spreadsheetToken 和维度信息合并单元格；单次操作不超过5000行，100列。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukDNzUjL5QzM14SO0MTN

POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/merge_cells
*/
func MergeCells(ctx *feishu.App, payload []byte, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiMergeCells
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+api, bytes.NewReader(payload), header)
}

/*
拆分单元格



该接口用于根据 spreadsheetToken 和维度信息拆分单元格；单次操作不超过5000行，100列。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uATNzUjLwUzM14CM1MTN

POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/unmerge_cells
*/
func UnmergeCells(ctx *feishu.App, payload []byte, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiUnmergeCells
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+api, bytes.NewReader(payload), header)
}

/*
读取单个范围



该接口用于根据 spreadsheetToken 和 range 读取表格单个范围的值，返回数据限制为10M。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ugTMzUjL4EzM14COxMTN

GET https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values/:range
*/
func Range(ctx *feishu.App, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiRange
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+api, header)
}

/*
向单个范围写入数据

该接口用于根据 spreadsheetToken 和 range 向单个范围写入数据，若范围内有数据，将被更新覆盖；单次写入不超过5000行，100列，每个格子大小为0.5M。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uAjMzUjLwIzM14CMyMTN

PUT https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values
*/
func Values(ctx *feishu.App, payload []byte, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiValues
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPut(feishu.FeishuServerUrl+api, bytes.NewReader(payload), header)
}

/*
读取多个范围



该接口用于根据 spreadsheetToken 和 ranges 读取表格多个范围的值，返回数据限制为10M。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukTMzUjL5EzM14SOxMTN

GET https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_batch_get
*/
func ValuesBatchGet(ctx *feishu.App, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiValuesBatchGet
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+api+"?"+params.Encode(), header)
}

/*
向多个范围写入数据



该接口用于根据 spreadsheetToken 和 range 向多个范围写入数据，若范围内有数据，将被更新覆盖；单次写入不超过5000行，100列，每个格子大小为0.5M。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uEjMzUjLxIzM14SMyMTN

POST https://open.feishu.cn/open-apis/sheet/v2/spreadsheets/:spreadsheetToken/values_batch_update
*/
func ValuesBatchUpdate(ctx *feishu.App, payload []byte, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiValuesBatchUpdate
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+api, bytes.NewReader(payload), header)
}
