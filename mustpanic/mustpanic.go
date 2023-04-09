package must

// Ok error must be not nil
func Ok(err error) {
	if err != nil {
		panic(err)
	}
}

// Ignore ignores the error code
func Ignore(_ error) {
}

// OkSkipOne error must be not nil and skips one additional return value
func OkSkipOne[T any](_ T, err error) {
	if err != nil {
		panic(err)
	}
}

// OkSkipTwo error must be not nil and skips two additional return values
func OkSkipTwo[T any, S any](_ T, _ S, err error) {
	if err != nil {
		panic(err)
	}
}

// OkOne error must be not nil and returns one value
func OkOne[T any](arg T, err error) T {
	if err != nil {
		panic(err)
	}
	return arg
}

// OkTwo error must be not nil and returns two values
func OkTwo[T any, S any](arg1 T, arg2 S, err error) (T, S) {
	if err != nil {
		panic(err)
	}
	return arg1, arg2
}
