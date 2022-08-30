package dynamic

func wReplace(a byte, b byte) int {
	if a == b {
		return 0
	}
	return 1
}

func wInsert(a byte) int {
	return 1
}

func wDelete(a byte) int {
	return 1
}

func localMin(integers ...int) (min int, index int) {
	n := len(integers)
	if n == 0 {
		return min, -1
	}
	min, index = integers[0], 0
	for i := 1; i < n; i++ {
		if integers[i] < min {
			min, index = integers[i], i
		}
	}
	return
}

func EditDistance(s1 string, s2 string) int {
	rows, cols := len(s1), len(s2)
	memoTable := make([][]int, rows)
	for i, _ := range memoTable {
		memoTable[i] = make([]int, cols)
	}
	return editDistance(&s1, &s2, &memoTable, 0, 0)
}

func editDistance(x *string, y *string, memoTable *[][]int, i int, j int) (solution int) {
	m, n := len(*x) - i, len(*y) - j
	if m <= 0 && n <= 0 {
		return
	}
	if m <= 0 {
		solution = wInsert((*y)[j]) + editDistance(x, y, memoTable, i, j+1)
		return
	}
	if n <= 0 {
		solution = wDelete((*x)[i]) + editDistance(x, y, memoTable, i+1, j)
		return
	}
	solution = (*memoTable)[i][j]
	if solution != 0 {
		return
	}
	deleteCost := wDelete((*x)[i]) + editDistance(x, y, memoTable, i+1, j)
	insertCost := wInsert((*y)[j]) + editDistance(x, y, memoTable, i, j+1)
	replaceCost := wReplace((*x)[i], (*y)[j]) + editDistance(x, y, memoTable, i+1, j+1)
	solution, _ = localMin(deleteCost, insertCost, replaceCost)
	(*memoTable)[i][j] = solution
	return
}
