package pointinator

import (
	"reflect"
	"testing"
)

func TestInt8(t *testing.T) {
	type args struct {
		i int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{"ok", args{42}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8(tt.args.i); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Int8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZero8(t *testing.T) {
	tests := []struct {
		name string
		want int8
	}{
		{"ok", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zero8(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Zero8() = %v, want %v", got, tt.want)
			}
		})
	}
}
