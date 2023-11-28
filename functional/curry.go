package functional

func Add(a int, b int, c int) int {
	return a + b + c
}

func AddCurry(a int) func(b int) func(c int) int {
	return func(b int) func(c int) int {
		return func(c int) int {
			return a + b + c
		}
	}
}
