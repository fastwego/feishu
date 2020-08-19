// Package platform 云文档/platform
package platform

import (
	"bytes"
	"net/http"

	"github.com/fastwego/feishu"
)

const (
	apiDocsMeta     = "/open-apis/suite/docs-api/meta"
	apiSearchObject = "/open-apis/suite/docs-api/search/object"
)

/*
获取元数据



该接口用于根据 token 获取各类文件的元数据


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMjN3UjLzYzN14yM2cTN

POST https://open.feishu.cn/open-apis/suite/docs-api/meta
*/
func DocsMeta(ctx *feishu.App, payload []byte, accessToken string) (resp []byte, err error) {

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiDocsMeta, bytes.NewReader(payload), header)
}

/*
文档搜索



该接口用于根据搜索条件进行文档搜索


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ugDM4UjL4ADO14COwgTN

POST https://open.feishu.cn/open-apis/suite/docs-api/search/object
*/
func SearchObject(ctx *feishu.App, payload []byte, accessToken string) (resp []byte, err error) {

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiSearchObject, bytes.NewReader(payload), header)
}
