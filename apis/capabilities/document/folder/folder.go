// Package folder 云文档/folder
package folder

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"

	"github.com/fastwego/feishu"
)

const (
	apiCreate         = "/open-apis/drive/explorer/v2/folder/:folderToken"
	apiMeta           = "/open-apis/drive/explorer/v2/folder/:folderToken/meta"
	apiRootFolderMeta = "/open-apis/drive/explorer/v2/root_folder/meta"
	apiChildren       = "/open-apis/drive/explorer/v2/folder/:folderToken/children"
)

/*
新建文件夹



该接口用于根据 folderToken 在该 folder 下创建文件夹


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukTNzUjL5UzM14SO1MTN

POST https://open.feishu.cn/open-apis/drive/explorer/v2/folder/:folderToken
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
获取文件夹元信息


该接口用于根据 folderToken 获取该文件夹的元信息

See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uAjNzUjLwYzM14CM2MTN

GET https://open.feishu.cn/open-apis/drive/explorer/v2/folder/:folderToken/meta
*/
func Meta(ctx *feishu.App, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiMeta
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+api, header)
}

/*
获取root folder（我的空间） meta

该接口用于获取 "我的空间" 的元信息

See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uAjNzUjLwYzM14CM2MTN

GET https://open.feishu.cn/open-apis/drive/explorer/v2/root_folder/meta
*/
func RootFolderMeta(ctx *feishu.App, accessToken string) (resp []byte, err error) {

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+apiRootFolderMeta, header)
}

/*
获取文件夹下的文档清单



该接口用于根据 folderToken 获取该文件夹的文档清单，如 doc、sheet、folder


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uEjNzUjLxYzM14SM2MTN

GET https://open.feishu.cn/open-apis/drive/explorer/v2/folder/:folderToken/children
*/
func Children(ctx *feishu.App, params url.Values, accessToken string) (resp []byte, err error) {

	api := apiChildren
	for paramName, paramValues := range params {
		api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPGet(feishu.FeishuServerUrl+api+"?"+params.Encode(), header)
}
