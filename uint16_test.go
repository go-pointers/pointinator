package pointinator

import (
	"reflect"
	"testing"
)

func TestUInt16(t *testing.T) {
	type args struct {
		i uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{"ok", args{42}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt16(tt.args.i); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Int16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUZero16(t *testing.T) {
	tests := []struct {
		name string
		want uint16
	}{
		{"ok", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UZero16(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Zero16() = %v, want %v", got, tt.want)
			}
		})
	}
}
