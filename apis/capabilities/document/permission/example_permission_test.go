package permission_test

import (
	"fmt"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/capabilities/document/permission"
)

func ExampleMemberCreate() {
	var ctx *feishu.App

	payload := []byte("{}")
	accessToken := ""
	resp, err := permission.MemberCreate(ctx, payload, accessToken)

	fmt.Println(resp, err)
}

func ExampleMemberTransfer() {
	var ctx *feishu.App

	payload := []byte("{}")
	accessToken := ""
	resp, err := permission.MemberTransfer(ctx, payload, accessToken)

	fmt.Println(resp, err)
}

func ExamplePublicUpdate() {
	var ctx *feishu.App

	payload := []byte("{}")
	accessToken := ""
	resp, err := permission.PublicUpdate(ctx, payload, accessToken)

	fmt.Println(resp, err)
}

func ExampleMemberList() {
	var ctx *feishu.App

	payload := []byte("{}")
	accessToken := ""
	resp, err := permission.MemberList(ctx, payload, accessToken)

	fmt.Println(resp, err)
}

func ExampleMemberDelete() {
	var ctx *feishu.App

	payload := []byte("{}")
	accessToken := ""
	resp, err := permission.MemberDelete(ctx, payload, accessToken)

	fmt.Println(resp, err)
}

func ExampleMemberUpdate() {
	var ctx *feishu.App

	payload := []byte("{}")
	accessToken := ""
	resp, err := permission.MemberUpdate(ctx, payload, accessToken)

	fmt.Println(resp, err)
}

func ExampleMemberPermitted() {
	var ctx *feishu.App

	payload := []byte("{}")
	accessToken := ""
	resp, err := permission.MemberPermitted(ctx, payload, accessToken)

	fmt.Println(resp, err)
}
