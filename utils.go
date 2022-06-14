package main

func mustA[T any](a T, err error) T {
	if err != nil {
		panic(err)
	}
	return a
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
