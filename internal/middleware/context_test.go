package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetID(t *testing.T) {
	h := func(w http.ResponseWriter, r *http.Request) {}
	type args struct {
		start int64
	}
	testCases := []struct {
		name string
		args args
	}{
		{
			"base case",
			args{54},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			SetID(tc.args.start)(h)(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))
		})
	}
}

func TestGetID(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	testCases := []struct {
		name string
		args args
		want string
	}{
		{
			"base case",
			args{context.WithValue(context.Background(), ID, "323")},
			"323",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := GetID(tc.args.ctx); got != tc.want {
				t.Errorf("GetID() = %v, want %v", got, tc.want)
			}
		})
	}
}
