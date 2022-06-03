package pointinator

import (
	"reflect"
	"testing"
)

func TestAny(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"ok struct", args{struct{}{}}, struct{}{}},
		{"ok int", args{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Any(tt.args.x); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}
