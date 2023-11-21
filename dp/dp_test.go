package dp

import "testing"

func TestUniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{3, 7},
			want: 28,
		},
		{
			name: "case2",
			args: args{3, 2},
			want: 3,
		},
		{
			name: "case3",
			args: args{7, 3},
			want: 28,
		},
		{
			name: "case4",
			args: args{3, 3},
			want: 6,
		},
		{
			name: "case5",
			args: args{1, 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("UniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
