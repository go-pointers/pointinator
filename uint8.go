package pointinator

func UZero8() *uint8 {
	i := uint8(0)
	return &i
}

func UInt8(i uint8) *uint8 {
	return &i
}
