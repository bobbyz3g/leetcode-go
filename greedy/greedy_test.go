package greedy

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				gas:  []int{1, 2, 3, 4, 5},
				cost: []int{3, 4, 5, 1, 2},
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				gas:  []int{2, 3, 4},
				cost: []int{3, 4, 3},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanCompleteCircuit(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("CanCompleteCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaximumSwap(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{num: 2376},
			want: 7326,
		},
		{
			name: "case2",
			args: args{num: 9953},
			want: 9953,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximumSwap(tt.args.num); got != tt.want {
				t.Errorf("MaximumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
