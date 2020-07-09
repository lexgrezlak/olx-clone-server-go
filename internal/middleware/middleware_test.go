package middleware

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestApplyMiddleware(t *testing.T) {
	h := func(http.ResponseWriter, *http.Request) {}
	type args struct {
		h          http.HandlerFunc
		middleware []Middleware
	}

	testCases := []struct {
		name string
		args args
	}{
		{
			"base case",
			args{
				h:          h,
				middleware: []Middleware{SetID(34)},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ApplyMiddleware(tc.args.h, tc.args.middleware...)
		})
	}

}

func TestLogger(t *testing.T) {
	h := func(http.ResponseWriter, *http.Request) {}
	type args struct {
		l *log.Logger
	}

	testCases := []struct {
		name string
		args args
	}{
		{
			"base case",
			args{log.New(os.Stdout, "", 0)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			Logger(tc.args.l)(h)(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))
		})
	}
}
