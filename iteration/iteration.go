package iteration

func Repeat(character string, round int) string {
	var repeated string
	for i := 0; i < round; i++ {
		repeated += character
	}
	return repeated
}
