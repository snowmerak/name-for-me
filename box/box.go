package box

func New[T any](v T) *T {
	return &v
}
