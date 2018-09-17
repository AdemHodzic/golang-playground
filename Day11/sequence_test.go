package day11

import "testing"

func TestSequence(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"#1", args{[]int{1, 3, 2, 1}}, false},
		{"#2", args{[]int{1, 3, 2}}, true},
		{"#3", args{[]int{1, 2, 1, 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sequence(tt.args.arr); got != tt.want {
				t.Errorf("Sequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
