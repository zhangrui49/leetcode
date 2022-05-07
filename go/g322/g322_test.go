package g322

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "test1",
		// 	args: args{
		// 		amount: 11,
		// 		coins:  []int{1, 2, 5},
		// 	},
		// 	want: 3,
		// },
		// {
		// 	name: "test2",
		// 	args: args{
		// 		amount: 6249,
		// 		coins:  []int{186, 419, 83, 408},
		// 	},
		// 	want: 20,
		// },
		// {
		// 	name: "test3",
		// 	args: args{
		// 		amount: 100,
		// 		coins:  []int{1, 2, 5},
		// 	},
		// 	want: 20,
		// },

		{
			name: "test41",
			args: args{
				amount: 3,
				coins:  []int{2},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coinChangeV3(tt.args.coins, tt.args.amount)
			t.Errorf("coinChange() = %v, want %v", got, tt.want)
		})
	}
}
