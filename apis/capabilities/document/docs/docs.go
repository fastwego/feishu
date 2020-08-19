// Package docs 云文档/docs
package docs

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/fastwego/feishu"
)

const (
	apiRawContent = "/open-apis/doc/v2/:docToken/raw_content"
	apiSheetMeta  = "/open-apis/doc/v2/:docToken/sheet_meta"
	apiDocMeta    = "/open-apis/doc/v2/meta/:docToken"
)

/*
获取文档文本内容


该接口用于获取文档的纯文本内容，不包含富文本格式信息，主要用于搜索，如导入 es 等。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukzNzUjL5czM14SO3MTN

GET https://open.feishu.cn/open-apis/doc/v2/:docToken/raw_content
*/
func RawContent(ctx *feishu.App, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiRawContent
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+api, header)
}

/*
获取sheet@doc的元数据

该接口用于根据 docToken 获取 sheet@doc 的元数据。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uADOzUjLwgzM14CM4MTN

GET https://open.feishu.cn/open-apis/doc/v2/:docToken/sheet_meta
*/
func SheetMeta(ctx *feishu.App, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiSheetMeta
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+api, header)
}

/*
获取文档元信息


该接口用于根据 docToken 获取元数据。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uczN3UjL3czN14yN3cTN

GET https://open.feishu.cn/open-apis/doc/v2/meta/:docToken
*/
func DocMeta(ctx *feishu.App, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiDocMeta
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+api, header)
}
