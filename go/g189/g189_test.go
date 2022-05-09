package g189

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotateV3(tt.args.nums, tt.args.k)
		})
	}
}
