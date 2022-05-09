package g283

import "testing"

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
		},
		{
			name: "test2",
			args: args{
				nums: []int{1, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroesV2(tt.args.nums)
			t.Error((tt.args.nums))
		})
	}
}
