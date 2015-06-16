package helpers

// UnquoteIfNeeded unwraps a string that is quoted between a given sign.
// If the source string is not wrapped in this char, it will be returned as is.
func UnquoteIfNeeded(source string, quoteChar rune) string {
	// strings.IndexRune()
	s := []rune(source)
	if s[0] == quoteChar && s[len(source)-1] == quoteChar {
		return string(s[1 : len(s)-1])
	}
	return source
}
