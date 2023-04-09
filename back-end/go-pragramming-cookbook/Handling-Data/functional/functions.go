package functional

import "strings"

func LowerCaseData(o Item) Item {
	o.Data = strings.ToLower(o.Data)
	return o
}

func IncreaseVersion(o Item) Item {
	o.Version++
	return o
}

func OldVersionFilter(version int) func(o Item) bool {
	return func(o Item) bool {
		return o.Version >= version
	}
}
