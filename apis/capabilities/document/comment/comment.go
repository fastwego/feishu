// Package comment 云文档/评论
package comment

import (
	"bytes"
	"net/http"

	"github.com/fastwego/feishu"
)

const (
	apiAddWhole = "/open-apis/comment/add_whole"
)

/*
添加全文评论



该接口用于根据 filetoken 给文档添加全文评论


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ucDN4UjL3QDO14yN0gTN

POST https://open.feishu.cn/open-apis/comment/add_whole
*/
func AddWhole(ctx *feishu.App, payload []byte, accessToken string) (resp []byte, err error) {

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiAddWhole, bytes.NewReader(payload), header)
}
