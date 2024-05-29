package g2981

import "testing"

func Test_maximumLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "case 1",
		// 	args: args{s: "aaaa"},
		// 	want: 2,
		// },
		// {
		// 	name: "case 2",
		// 	args: args{s: "abcdef"},
		// 	want: -1,
		// },
		// {
		// 	name: "case 3",
		// 	args: args{s: "abcaba"},
		// 	want: 1,
		// },
		// {
		// 	name: "case 4",
		// 	args: args{s: "fafff"},
		// 	want: 1,
		// },
		{
			name: "case 5",
			args: args{s: "ceeeeeeeeeeeebmmmfffeeeeeeeeeeeewww"},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumLength(tt.args.s); got != tt.want {
				t.Errorf("maximumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
