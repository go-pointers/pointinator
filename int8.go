package pointinator

func Zero8() *int8 {
	i := int8(0)
	return &i
}

func Int8(i int8) *int8 {
	return &i
}
