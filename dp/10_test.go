package dp

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple",
			args: args{
				s: "abccdef",
				p: "abc*def",
			},
			want: true,
		},
		{
			name: ".*",
			args: args{
				s: "abccdef",
				p: ".*",
			},
			want: true,
		},
		{
			name: ".*.*",
			args: args{
				s: "abccdef",
				p: ".*.",
			},
			want: true,
		},
		{
			name: "complex",
			args: args{
				s: "abccdsdef",
				p: "abc*d.",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
