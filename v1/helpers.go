package v1

// Ptr returns a pointer to v. Useful for optional query params in generated methods.
//
//go:fix inline
func Ptr[T any](v T) *T { return new(v) }
