package moe

import (
	"net/http"
	"testing"
)

func TestResponse_Send(t *testing.T) {
	type fields struct {
		Status       int
		body         []byte
		Header       http.Header
		LastModified string
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &Response{
				Status:       tt.fields.Status,
				body:         tt.fields.body,
				Header:       tt.fields.Header,
				LastModified: tt.fields.LastModified,
			}
			if err := res.Send(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Response.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
