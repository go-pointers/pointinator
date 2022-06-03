package pointinator

import (
	"reflect"
	"testing"
)

func TestInt(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ok", args{42}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.i); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZero(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"ok", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zero(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Zero() = %v, want %v", got, tt.want)
			}
		})
	}
}
