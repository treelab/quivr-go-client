package quivr_go_client

func Ptr[T any](v T) *T {
	return &v
}
