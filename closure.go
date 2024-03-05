func Closure() func() int {
	return func() int {
		return 1
	}
}