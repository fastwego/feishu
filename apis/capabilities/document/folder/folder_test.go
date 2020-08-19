package folder

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

func TestCreate(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiCreate, func(w http.ResponseWriter, r *http.Request) {
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
			gotResp, err := Create(tt.args.ctx, tt.args.payload, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Create() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestMeta(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiMeta, func(w http.ResponseWriter, r *http.Request) {
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
			gotResp, err := Meta(tt.args.ctx, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("Meta() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Meta() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestRootFolderMeta(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiRootFolderMeta, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("GET")

	type args struct {
		ctx *feishu.App

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
			gotResp, err := RootFolderMeta(tt.args.ctx, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("RootFolderMeta() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("RootFolderMeta() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestChildren(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiChildren, func(w http.ResponseWriter, r *http.Request) {
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
			gotResp, err := Children(tt.args.ctx, tt.args.params, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("Children() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Children() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
