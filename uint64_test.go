package pointinator

import (
	"reflect"
	"testing"
)

func TestUInt64(t *testing.T) {
	type args struct {
		i uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"ok", args{42}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt64(tt.args.i); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUZero64(t *testing.T) {
	tests := []struct {
		name string
		want uint64
	}{
		{"ok", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UZero64(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Zero64() = %v, want %v", got, tt.want)
			}
		})
	}
}
