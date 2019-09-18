package util

type Comparable interface {
	Equals(other Comparable) bool
	Less(other Comparable) bool
}

func Min(elements ...Comparable) Comparable {
	var min Comparable
	for _, element := range elements {
		if min == nil || element.Less(min) {
			min = element
		}
	}
	return min
}
