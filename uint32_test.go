package pointinator

import (
	"reflect"
	"testing"
)

func TestUInt32(t *testing.T) {
	type args struct {
		i uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"ok", args{42}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt32(tt.args.i); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUZero32(t *testing.T) {
	tests := []struct {
		name string
		want uint32
	}{
		{"ok", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UZero32(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Zero32() = %v, want %v", got, tt.want)
			}
		})
	}
}
