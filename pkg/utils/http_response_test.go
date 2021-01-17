package utils

import (
	"net/http"
	"testing"
)

func TestHTTPJSONResponse_Resp(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		code int
		data interface{}
	}
	tests := []struct {
		name string
		h    *HTTPJSONResponse
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HTTPJSONResponse{}
			h.Resp(tt.args.w, tt.args.code, tt.args.data)
		})
	}
}

func TestHTTPJSONResponse_Err(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		code int
		err  error
	}
	tests := []struct {
		name string
		h    *HTTPJSONResponse
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HTTPJSONResponse{}
			h.Err(tt.args.w, tt.args.code, tt.args.err)
		})
	}
}
