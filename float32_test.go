package pointinator

import (
	"reflect"
	"testing"
)

func TestFZero32(t *testing.T) {
	tests := []struct {
		name string
		want float32
	}{
		{"ok", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FZero32(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("FZero32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat321(t *testing.T) {
	type args struct {
		f float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"ok", args{42}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32(tt.args.f); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}
