package functional

type Item struct {
	Data    string
	Version int
}

type Collection = []Item

func Filter(c Collection, filter func(o Item) bool) Collection {
	ret := make(Collection, 0)

	for _, o := range c {
		if filter(o) {
			ret = append(ret, o)
		}
	}
	return ret
}

func Map(c Collection, mapping func(o Item) Item) Collection {
	ret := make(Collection, len(c))

	for pos, o := range c {
		ret[pos] = mapping(o)
	}
	return ret
}
