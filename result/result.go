package result

type Result[T any] struct {
	Result T
	Err error
}
