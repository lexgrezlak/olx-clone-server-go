package handler

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"net/http/httptest"
	"olx-clone-server/internal/service"
	"olx-clone-server/internal/util"
	"testing"
)

func TestSignUp(t *testing.T) {
	testDb := service.NewTestDatabase(t)

	testCases := []struct {
		name string
		input service.SignUpInput
		want int
	}{
		{
			name: "valid input",
			input: service.SignUpInput{
				FirstName: "johnrwewerw",
				LastName:  "adskasd",
				Email:     "asdsaad@das.sad",
				Password:  "askdaskdask",
			},
			want: 201,
		},		{
			name: "not enough input",
			input: service.SignUpInput{
				LastName:  "adskasd",
				Email:     "asdsaad@das.sad",
				Password:  "askdaskdask",
			},
			want: 400,
		},
	}



	for _, tc := range testCases {
		body := util.EncodeJSONBody(&tc.input)
		req := httptest.NewRequest("Post", "/auth/sign-up", body)
		res := httptest.NewRecorder()
		h := SignUp(testDb)
		h(res, req)

		t.Run(tc.name, func(t *testing.T) {
			got := res.Code
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("mismatch (-want, +got): \n%s", diff)
			}
		})
	}
}