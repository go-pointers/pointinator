package pointinator

func Zero64() *int64 {
	i := int64(0)
	return &i
}

func Int64(i int64) *int64 {
	return &i
}
