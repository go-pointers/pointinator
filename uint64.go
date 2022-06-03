package pointinator

func UZero64() *uint64 {
	i := uint64(0)
	return &i
}

func UInt64(i uint64) *uint64 {
	return &i
}
