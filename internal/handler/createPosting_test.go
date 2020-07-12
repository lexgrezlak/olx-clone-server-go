package handler

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"log"
	"net/http"
	"net/http/httptest"
	"olx-clone-server/internal/service"
	"testing"
)

func TestCreatePosting(t *testing.T) {
	api, _, mock, err := service.NewTestAPI()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("invalid token", func(t *testing.T) {
		input := service.CreatePostingInput{
			Title:       "Titleeee",
			Price:       939,
			Condition:   "New",
			Description: "Hello world hello world",
			Phone:       9485934833,
			City:        "Zurich",
			Photos:      []string{},
		}

		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(input)
		if err != nil {
			log.Fatal(err)
		}

		var lastInsertID, affected int64
		mock.ExpectExec("INSERT .*").WillReturnResult(sqlmock.NewResult(lastInsertID,affected))
		req := httptest.NewRequest("POST", "/postings", buf)
		res := httptest.NewRecorder()
		req.Header.Set("token", "hello1world")
		h := CreatePosting(api)
		h(res, req)

		type args struct {
			Code int
			LastInsertID int64
			Affected int64
		}

		want := args{
			Code:         http.StatusBadRequest,
			LastInsertID: 0,
			Affected:     0,
		}

		got := args{
			Code:         res.Code,
			LastInsertID: lastInsertID,
			Affected:     affected,
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want, +got): \n%s", diff)
		}
	})

}
