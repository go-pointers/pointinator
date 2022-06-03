package pointinator

func EmptyString() *string {
	s := ""
	return &s
}

func String(s string) *string {
	return &s
}
