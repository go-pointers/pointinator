package pointinator

func Zero16() *int16 {
	i := int16(0)
	return &i
}

func Int16(i int16) *int16 {
	return &i
}

