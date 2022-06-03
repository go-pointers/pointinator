package pointinator

func UZero16() *uint16 {
	i := uint16(0)
	return &i
}

func UInt16(i uint16) *uint16 {
	return &i
}
