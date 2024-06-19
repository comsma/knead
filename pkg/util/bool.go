package util

func NullableBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}
