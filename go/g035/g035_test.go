package g035

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			want: 4,
		},
		{
			name: "test3",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 0,
		},
		{
			name: "test4",
			args: args{
				nums:   []int{1, 3},
				target: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchInsert(tt.args.nums, tt.args.target)
			t.Errorf("searchInsert() = %v, want %v", got, tt.want)

		})
	}
}
