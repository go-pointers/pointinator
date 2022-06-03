package pointinator

import (
	"reflect"
	"testing"
)

func TestInt64(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"ok", args{42}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64(tt.args.i); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZero64(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"ok", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zero64(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Zero64() = %v, want %v", got, tt.want)
			}
		})
	}
}
