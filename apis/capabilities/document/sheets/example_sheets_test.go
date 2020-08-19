package sheets_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/capabilities/document/sheets"
)

func ExampleMetaInfo() {
	var ctx *feishu.App

	params := url.Values{}
	accessToken := ""
	resp, err := sheets.MetaInfo(ctx, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleUpdateProperties() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	accessToken := ""
	resp, err := sheets.UpdateProperties(ctx, payload, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleBatchUpdate() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	accessToken := ""
	resp, err := sheets.BatchUpdate(ctx, payload, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleValuesPrepend() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	accessToken := ""
	resp, err := sheets.ValuesPrepend(ctx, payload, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleValuesAppend() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	accessToken := ""
	resp, err := sheets.ValuesAppend(ctx, payload, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleInsertDimensionRange() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	accessToken := ""
	resp, err := sheets.InsertDimensionRange(ctx, payload, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleCreateDimensionRange() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	accessToken := ""
	resp, err := sheets.CreateDimensionRange(ctx, payload, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleUpdateDimensionRange() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	accessToken := ""
	resp, err := sheets.UpdateDimensionRange(ctx, payload, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleDeleteDimensionRange() {
	var ctx *feishu.App

	params := url.Values{}
	accessToken := ""
	resp, err := sheets.DeleteDimensionRange(ctx, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleUpdateStyle() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	accessToken := ""
	resp, err := sheets.UpdateStyle(ctx, payload, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleBatchUpdateStyle() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	accessToken := ""
	resp, err := sheets.BatchUpdateStyle(ctx, payload, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleProtectedDimension() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	accessToken := ""
	resp, err := sheets.ProtectedDimension(ctx, payload, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleMergeCells() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	accessToken := ""
	resp, err := sheets.MergeCells(ctx, payload, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleUnmergeCells() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	accessToken := ""
	resp, err := sheets.UnmergeCells(ctx, payload, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleRange() {
	var ctx *feishu.App

	params := url.Values{}
	accessToken := ""
	resp, err := sheets.Range(ctx, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleValues() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	accessToken := ""
	resp, err := sheets.Values(ctx, payload, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleValuesBatchGet() {
	var ctx *feishu.App

	params := url.Values{}
	accessToken := ""
	resp, err := sheets.ValuesBatchGet(ctx, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleValuesBatchUpdate() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	accessToken := ""
	resp, err := sheets.ValuesBatchUpdate(ctx, payload, params, accessToken)

	fmt.Println(resp, err)
}
