package g016

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "test1",
		// 	args: args{
		// 		nums:   []int{-1, 2, 1, -4},
		// 		target: 1,
		// 	},
		// 	want: 2,
		// },
		{
			name: "test2",
			args: args{
				nums:   []int{0, 1, 2},
				target: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSumClosest(tt.args.nums, tt.args.target)
			t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)

		})
	}
}
