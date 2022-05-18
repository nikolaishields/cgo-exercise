package main

import "testing"

func Test_doAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Add 2 numbers", args{2, 3}, 5, false},
		{"Add 2 numbers", args{0, 0}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := doAdd(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("doAdd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("doAdd() got = %v, want %v", got, tt.want)
			}
		})
	}
}
