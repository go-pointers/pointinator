//go:build go1.18

package pointinator

func Any[T any](x T) *T {
	return &x
}
