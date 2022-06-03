package pointinator

func Zero() *int {
	i := 0
	return &i
}

func Int(i int) *int {
	return &i
}
