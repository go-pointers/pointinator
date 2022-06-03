package pointinator

func Zero32() *int32 {
	i := int32(0)
	return &i
}

func Int32(i int32) *int32 {
	return &i
}
