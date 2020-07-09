package middleware

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)


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
			Logger(http.Handler())(h)
		})
	}
}
