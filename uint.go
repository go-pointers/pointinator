package pointinator

func UZero() *uint {
	i := uint(0)
	return &i
}

func UInt(i uint) *uint {
	return &i
}
