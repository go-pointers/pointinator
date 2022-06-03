package pointinator

import (
	"reflect"
	"testing"
)

func TestUInt8(t *testing.T) {
	type args struct {
		i uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{"ok", args{42}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt8(tt.args.i); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Int8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUZero8(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		{"ok", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UZero8(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Zero8() = %v, want %v", got, tt.want)
			}
		})
	}
}
