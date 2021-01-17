package utils

import (
	"testing"
)

func TestStringInSlice(t *testing.T) {
	if !StringInSlice("a", []string{"a", "b", "c", "d"}) {
		t.Error("should return true since list has a")
	}

	if StringInSlice("a", []string{"b", "c", "d"}) {
		t.Error("should return false since list doesnt has a")
	}
}

func TestStringToIntWithDefault(t *testing.T) {
	type args struct {
		s    string
		dflt int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "invalid number",
			args: args{
				"a", 0,
			},
			want: 0,
		},
		{
			name: "valid number",
			args: args{
				"1", 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToIntWithDefault(tt.args.s, tt.args.dflt); got != tt.want {
				t.Errorf("StringToIntWithDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
