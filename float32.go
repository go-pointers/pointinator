package pointinator

func FZero32() *float32 {
	f := float32(0)
	return &f
}

func Float32(f float32) *float32 {
	return &f
}
