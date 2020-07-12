package handler

import (
	"github.com/google/go-cmp/cmp"
	"net/http/httptest"
	"olx-clone-server/internal/database"
	"testing"
)

func TestAllPostings(t *testing.T) {
	testDb := database.NewTestDatabase(t)

	testCases := []struct {
		name string
		want int
	}{
		{
			name: "valid request",
			want: 200,
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest("GET", "/postings", nil)
		res := httptest.NewRecorder()
		h := AllPostings(testDb)
		h(res, req)

		t.Run(tc.name, func(t *testing.T) {
			got := res.Code
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("mismatch (-want, +got): \n%s", diff)
			}
		})
	}
}
