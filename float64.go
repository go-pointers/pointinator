package pointinator

func FZero64() *float64 {
	f := float64(0)
	return &f
}

func Float64(f float64) *float64 {
	return &f
}
