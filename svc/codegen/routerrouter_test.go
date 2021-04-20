package codegen

import "testing"

func TestGenRouterRouter(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				dir: "/Users/wubin1989/workspace/cloud/comment-svc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenRouterRouter(tt.args.dir)
		})
	}
}
