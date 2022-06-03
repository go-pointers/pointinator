package pointinator

import (
	"reflect"
	"testing"
)

func TestInt32(t *testing.T) {
	type args struct {
		i int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"ok", args{42}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32(tt.args.i); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZero32(t *testing.T) {
	tests := []struct {
		name string
		want int32
	}{
		{"ok", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zero32(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Zero32() = %v, want %v", got, tt.want)
			}
		})
	}
}
