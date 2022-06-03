package pointinator

import (
	"reflect"
	"testing"
)

func TestEmptyString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"ok", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EmptyString(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("EmptyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ok", args{"The answer to life, the universe and everything"}, "The answer to life, the universe and everything"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.s); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
