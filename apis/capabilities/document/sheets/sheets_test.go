package sheets

import (
	"net/http"
	"net/url"
	"os"
	"reflect"
	"testing"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/test"
)

func TestMain(m *testing.M) {
	test.Setup()
	os.Exit(m.Run())
}

func TestMetaInfo(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiMetaInfo, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("GET")

	type args struct {
		ctx *feishu.App

		params      url.Values
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := MetaInfo(tt.args.ctx, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("MetaInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("MetaInfo() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestUpdateProperties(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiUpdateProperties, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("PUT")

	type args struct {
		ctx     *feishu.App
		payload []byte

		params      url.Values
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := UpdateProperties(tt.args.ctx, tt.args.payload, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateProperties() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UpdateProperties() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestBatchUpdate(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiBatchUpdate, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("POST")

	type args struct {
		ctx     *feishu.App
		payload []byte

		params      url.Values
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := BatchUpdate(tt.args.ctx, tt.args.payload, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("BatchUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("BatchUpdate() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestValuesPrepend(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiValuesPrepend, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("POST")

	type args struct {
		ctx     *feishu.App
		payload []byte

		params      url.Values
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := ValuesPrepend(tt.args.ctx, tt.args.payload, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValuesPrepend() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("ValuesPrepend() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestValuesAppend(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiValuesAppend, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("POST")

	type args struct {
		ctx     *feishu.App
		payload []byte

		params      url.Values
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := ValuesAppend(tt.args.ctx, tt.args.payload, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValuesAppend() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("ValuesAppend() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestInsertDimensionRange(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiInsertDimensionRange, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("POST")

	type args struct {
		ctx     *feishu.App
		payload []byte

		params      url.Values
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := InsertDimensionRange(tt.args.ctx, tt.args.payload, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertDimensionRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("InsertDimensionRange() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestCreateDimensionRange(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiCreateDimensionRange, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("POST")

	type args struct {
		ctx     *feishu.App
		payload []byte

		params      url.Values
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := CreateDimensionRange(tt.args.ctx, tt.args.payload, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDimensionRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CreateDimensionRange() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestUpdateDimensionRange(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiUpdateDimensionRange, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("PUT")

	type args struct {
		ctx     *feishu.App
		payload []byte

		params      url.Values
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := UpdateDimensionRange(tt.args.ctx, tt.args.payload, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateDimensionRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UpdateDimensionRange() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestDeleteDimensionRange(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiDeleteDimensionRange, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("DELETE")

	type args struct {
		ctx *feishu.App

		params      url.Values
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := DeleteDimensionRange(tt.args.ctx, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteDimensionRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("DeleteDimensionRange() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestUpdateStyle(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiUpdateStyle, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("PUT")

	type args struct {
		ctx     *feishu.App
		payload []byte

		params      url.Values
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := UpdateStyle(tt.args.ctx, tt.args.payload, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateStyle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UpdateStyle() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestBatchUpdateStyle(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiBatchUpdateStyle, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("PUT")

	type args struct {
		ctx     *feishu.App
		payload []byte

		params      url.Values
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := BatchUpdateStyle(tt.args.ctx, tt.args.payload, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("BatchUpdateStyle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("BatchUpdateStyle() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestProtectedDimension(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiProtectedDimension, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("POST")

	type args struct {
		ctx     *feishu.App
		payload []byte

		params      url.Values
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := ProtectedDimension(tt.args.ctx, tt.args.payload, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProtectedDimension() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("ProtectedDimension() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestMergeCells(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiMergeCells, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("POST")

	type args struct {
		ctx     *feishu.App
		payload []byte

		params      url.Values
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := MergeCells(tt.args.ctx, tt.args.payload, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("MergeCells() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("MergeCells() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestUnmergeCells(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiUnmergeCells, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("POST")

	type args struct {
		ctx     *feishu.App
		payload []byte

		params      url.Values
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := UnmergeCells(tt.args.ctx, tt.args.payload, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmergeCells() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UnmergeCells() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestRange(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiRange, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("GET")

	type args struct {
		ctx *feishu.App

		params      url.Values
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := Range(tt.args.ctx, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("Range() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Range() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestValues(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiValues, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("PUT")

	type args struct {
		ctx     *feishu.App
		payload []byte

		params      url.Values
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := Values(tt.args.ctx, tt.args.payload, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("Values() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Values() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestValuesBatchGet(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiValuesBatchGet, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("GET")

	type args struct {
		ctx *feishu.App

		params      url.Values
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := ValuesBatchGet(tt.args.ctx, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValuesBatchGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("ValuesBatchGet() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestValuesBatchUpdate(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiValuesBatchUpdate, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("POST")

	type args struct {
		ctx     *feishu.App
		payload []byte

		params      url.Values
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := ValuesBatchUpdate(tt.args.ctx, tt.args.payload, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValuesBatchUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("ValuesBatchUpdate() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
