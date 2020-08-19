// Package file 云文档/file
package file

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"

	"github.com/fastwego/feishu"
)

const (
	apiCreate    = "/open-apis/drive/explorer/v2/file/:folderToken"
	apiCopy      = "/open-apis/drive/explorer/v2/file/copy/files/:fileToken"
	apiDeleteDoc = "/open-apis/drive/explorer/v2/file/docs/:docToken"
)

/*
新建文档


该接口用于根据 folderToken 创建 Doc或 Sheet 。

若没有特定的文件夹用于承载创建的文档，可以先调用「获取文件夹元信息」文档中的「获取 root folder (我的空间) meta」接口，获得我的空间的 token，然后再使用此接口。创建的文档将会在「我的空间」的「归我所有」列表里。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQTNzUjL0UzM14CN1MTN

POST https://open.feishu.cn/open-apis/drive/explorer/v2/file/:folderToken
*/
func Create(ctx *feishu.App, payload []byte, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiCreate
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+api, bytes.NewReader(payload), header)
}

/*
复制文档



该接口用于根据文件 token 复制 Doc 或 Sheet  到目标文件夹中。

若没有特定的文件夹用于承载创建的文档，可以先调用「获取文件夹元信息」文档中的「获取 root folder (我的空间) meta」接口，获得我的空间的 token，然后再使用此接口。复制的文档将会在「我的空间」的「归我所有」列表里。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYTNzUjL2UzM14iN1MTN

POST https://open.feishu.cn/open-apis/drive/explorer/v2/file/copy/files/:fileToken
*/
func Copy(ctx *feishu.App, payload []byte, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiCopy
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+api, bytes.NewReader(payload), header)
}

/*
删除 Doc

该接口用于根据 docToken 删除对应的 Docs 文档。

See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uATM2UjLwEjN14CMxYTN

DELETE https://open.feishu.cn/open-apis/drive/explorer/v2/file/docs/:docToken
*/
func DeleteDoc(ctx *feishu.App, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiDeleteDoc
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPDelete(feishu.FeishuServerUrl+api, header)
}
