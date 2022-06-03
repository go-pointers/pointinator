package pointinator

import (
	"reflect"
	"testing"
)

func TestFZero64(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
		{"ok", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FZero64(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("FZero64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat641(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"ok", args{42}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64(tt.args.f); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}
