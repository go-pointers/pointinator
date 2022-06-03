package pointinator

import (
	"reflect"
	"testing"
)

func TestInt16(t *testing.T) {
	type args struct {
		i int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{"ok", args{42}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16(tt.args.i); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Int16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZero16(t *testing.T) {
	tests := []struct {
		name string
		want int16
	}{
		{"ok", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zero16(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Zero16() = %v, want %v", got, tt.want)
			}
		})
	}
}
