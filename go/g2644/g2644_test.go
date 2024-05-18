package g2644

import "testing"

func Test_maxDivScore(t *testing.T) {
	type args struct {
		nums     []int
		divisors []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			want: 3,
			args: args{
				nums:     []int{4, 7, 9, 3, 9},
				divisors: []int{5, 2, 3},
			},
		},
		{
			name: "t2",
			want: 10,
			args: args{
				nums:     []int{12},
				divisors: []int{10, 16},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDivScore(tt.args.nums, tt.args.divisors); got != tt.want {
				t.Errorf("maxDivScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
