package iteration

// Repeat returns repeated N-times char
func Repeat(char string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += char
	}
	return repeated
}
