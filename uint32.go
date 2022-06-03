package pointinator

func UZero32() *uint32 {
	i := uint32(0)
	return &i
}

func UInt32(i uint32) *uint32 {
	return &i
}
