package util

type Comparable interface {
	Equals(other interface{}) bool
	Less(other interface{}) bool
}

type CmpInt int

func (c CmpInt) Equals(other interface{}) bool {
	return c == other.(CmpInt)
}

func (c CmpInt) Less(other interface{}) bool {
	return c < other.(CmpInt)
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
