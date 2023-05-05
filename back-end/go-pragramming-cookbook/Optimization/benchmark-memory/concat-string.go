package benchmark_memory

func concat(values ...string) string {
	ret := ""
	for i, val := range values {
		ret += val
		if i != len(values) {
			ret += " "
		}
	}

	return ret
}
