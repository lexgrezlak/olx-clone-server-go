package handler

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"net/http/httptest"
	"olx-clone-server/internal/service"
	"testing"
)

func TestAllPostings(t *testing.T) {
	api, mock, err := service.NewTestAPI()
	if err != nil {
		t.Fatal(err)
	}

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
		t.Run(tc.name, func(t *testing.T) {
			mock.ExpectQuery("SELECT id, title, price, photos FROM posting").WillReturnRows(sqlmock.NewRows([]string{"id", "title", "price", "photos"}))
			req := httptest.NewRequest("GET", "/postings", nil)
			res := httptest.NewRecorder()
			h := AllPostings(api)
			h(res, req)
			got := res.Code
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("mismatch (-want, +got): \n%s", diff)
			}
		})
	}
}
