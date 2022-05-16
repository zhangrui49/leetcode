package ginterview0406

import (
	"reflect"
	"testing"
)

func Test_inorderSuccessor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
	}

	p := &TreeNode{Val: 1}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "test1",
			args: args{
				root: &TreeNode{
					Val:  2,
					Left: p,
					Right: &TreeNode{
						Val: 3,
					},
				},
				p: p,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderSuccessor(tt.args.root, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderSuccessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
