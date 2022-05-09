package g942

import (
	"testing"
)

func Test_diStringMatch(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				s: "IDID",
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				s: "III",
			},
			want: nil,
		},
		{
			name: "test3",
			args: args{
				s: "DDI",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diStringMatch(tt.args.s)
			t.Errorf("diStringMatch() = %v, want %v", got, tt.want)

		})
	}
}
