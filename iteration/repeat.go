package iteration

func Repeat(a string) string {

	var repeated string
	for range 5 {
		repeated += a
	}
	return repeated
}
