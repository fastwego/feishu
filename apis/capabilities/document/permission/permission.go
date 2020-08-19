// Package permission 云文档/permission
package permission

import (
	"bytes"
	"net/http"

	"github.com/fastwego/feishu"
)

const (
	apiMemberCreate    = "/open-apis/drive/permission/member/create"
	apiMemberTransfer  = "/open-apis/drive/permission/member/transfer"
	apiPublicUpdate    = "/open-apis/drive/permission/public/update"
	apiMemberList      = "/open-apis/drive/permission/member/list"
	apiMemberDelete    = "/open-apis/drive/permission/member/delete"
	apiMemberUpdate    = "/open-apis/drive/permission/member/update"
	apiMemberPermitted = "/open-apis/drive/permission/member/permitted"
)

/*
增加权限



该接口用于根据 filetoken 给用户增加文档的权限


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uMzNzUjLzczM14yM3MTN

POST https://open.feishu.cn/open-apis/drive/permission/member/create
*/
func MemberCreate(ctx *feishu.App, payload []byte, accessToken string) (resp []byte, err error) {

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiMemberCreate, bytes.NewReader(payload), header)
}

/*
转移拥有者



该接口用于根据文档信息和用户信息转移文档的所有者。


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uQzNzUjL0czM14CN3MTN

POST https://open.feishu.cn/open-apis/drive/permission/member/transfer
*/
func MemberTransfer(ctx *feishu.App, payload []byte, accessToken string) (resp []byte, err error) {

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiMemberTransfer, bytes.NewReader(payload), header)
}

/*
更新文档公共设置



该接口用于根据 filetoken 更新文档的公共设置


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ukTM3UjL5EzN14SOxcTN

POST https://open.feishu.cn/open-apis/drive/permission/public/update
*/
func PublicUpdate(ctx *feishu.App, payload []byte, accessToken string) (resp []byte, err error) {

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiPublicUpdate, bytes.NewReader(payload), header)
}

/*
获取协作者列表



该接口用于根据 filetoken 查询协作者，目前包括人('user')和群('chat')

**你能获取到协作者列表的前提是你对该文档有权限**


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uATN3UjLwUzN14CM1cTN

POST https://open.feishu.cn/open-apis/drive/permission/member/list
*/
func MemberList(ctx *feishu.App, payload []byte, accessToken string) (resp []byte, err error) {

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiMemberList, bytes.NewReader(payload), header)
}

/*
移除协作者权限



该接口用于根据 filetoken 移除文档协作者的权限


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYTN3UjL2UzN14iN1cTN

POST https://open.feishu.cn/open-apis/drive/permission/member/delete
*/
func MemberDelete(ctx *feishu.App, payload []byte, accessToken string) (resp []byte, err error) {

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiMemberDelete, bytes.NewReader(payload), header)
}

/*
更新协作者权限



该接口用于根据 filetoken 更新文档协作者的权限


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ucTN3UjL3UzN14yN1cTN

POST https://open.feishu.cn/open-apis/drive/permission/member/update
*/
func MemberUpdate(ctx *feishu.App, payload []byte, accessToken string) (resp []byte, err error) {

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiMemberUpdate, bytes.NewReader(payload), header)
}

/*
判断协作者是否有某权限



该接口用于根据 filetoken 判断当前登录用户是否具有某权限


See: https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uYzN3UjL2czN14iN3cTN

POST https://open.feishu.cn/open-apis/drive/permission/member/permitted
*/
func MemberPermitted(ctx *feishu.App, payload []byte, accessToken string) (resp []byte, err error) {

	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", "application/json")

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl+apiMemberPermitted, bytes.NewReader(payload), header)
}
