package util

import "testing"

func TestCoverage(t *testing.T) {
	type args struct {
		condition bool
	}
	testCases := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"no condition",
			args{true},
			true,
		},
	}
	for _, tc := range testCases {
		if err := Coverage(tc.args.condition); (err != nil) != tc.wantErr {
			t.Errorf("Coverage() error = %v, wantErr %v", err, tc.wantErr)
		}
	}
}
