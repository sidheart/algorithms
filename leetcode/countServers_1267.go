package leetcode

func anotherServerOnCross(grid [][]int, row, col int) bool {
	for rowItr, _ := range grid {
		if rowItr != row && grid[rowItr][col] == 1 {
			return true
		}
	}
	for colItr, _ := range grid[0] {
		if colItr != col && grid[row][colItr] == 1 {
			return true
		}
	}
	return false
}

func countServers(grid [][]int) int {
	total := 0
	for row, _ := range grid {
		for col, _ := range grid[0] {
			if grid[row][col] == 1 && anotherServerOnCross(grid, row, col) {
				total++
			}
		}
	}
	return total
}