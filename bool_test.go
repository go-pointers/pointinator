package pointinator

import (
	"reflect"
	"testing"
)

func TestFalse(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"ok", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := False(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("False() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrue(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"ok", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := True(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("False() = %v, want %v", got, tt.want)
			}
		})
	}
}
