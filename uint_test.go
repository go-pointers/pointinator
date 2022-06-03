package pointinator

import (
	"reflect"
	"testing"
)

func TestUInt(t *testing.T) {
	type args struct {
		i uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"ok", args{42}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt(tt.args.i); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUZero(t *testing.T) {
	tests := []struct {
		name string
		want uint
	}{
		{"ok", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UZero(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Zero() = %v, want %v", got, tt.want)
			}
		})
	}
}
