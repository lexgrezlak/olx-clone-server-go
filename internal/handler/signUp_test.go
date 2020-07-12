package handler

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"net/http/httptest"
	"olx-clone-server/internal/service"
	"olx-clone-server/internal/util"
	"testing"
)

func TestSignUp(t *testing.T) {
	api, _, mock, err := service.NewTestAPI()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("valid input", func(t *testing.T) {
		input := service.SignUpInput{
			FirstName: "Johny",
			LastName:  "Smithy",
			Email:     "john123@icloud.com",
			Password:  "Kjqj3i432n",
		}
		var lastInsertID, affected int64
		mock.ExpectExec("INSERT .*").WillReturnResult(sqlmock.NewResult(lastInsertID, affected))

		body := util.EncodeJSONBody(&input)
		req := httptest.NewRequest("Post", "/auth/sign-up", body)
		res := httptest.NewRecorder()
		h := SignUp(api)
		h(res, req)

		want := http.StatusCreated
		got := res.Code
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want, +got): \n%s", diff)
		}
	})

}
