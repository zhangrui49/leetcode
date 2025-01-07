package g2274

import "testing"

func Test_maxConsecutive(t *testing.T) {
	type args struct {
		bottom  int
		top     int
		special []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{"case 1", args{2, 9, []int{4, 6}}, 3},
		{"case 2", args{6, 8, []int{7, 6, 8}}, 0},
		{"case 3", args{28, 50, []int{35, 48}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxConsecutive(tt.args.bottom, tt.args.top, tt.args.special); got != tt.want {
				t.Errorf("maxConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
