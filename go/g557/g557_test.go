package g557

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{
			name: "test2",
			args: args{
				s: "Let's take LeetCode contest",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseWordsV2(tt.args.s)
			t.Errorf("reverseWords() = %v, want %v", got, tt.want)

		})
	}
}
