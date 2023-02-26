package subcmd

// Contains : check that an array contains a value
// Exercise 2.2
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
