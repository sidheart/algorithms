package leetcode

type Cell struct {
	Row,Col int
}

type Island []Cell

func (island Island) isClosed(n, m int) bool {
	for _, cell := range island {
		r, c := cell.Row, cell.Col
		if r == 0 || r == n - 1 || c == 0 || c == m - 1 {
			return false
		}
	}
	return true
}

func constructIsland(startCell Cell, n int, m int, newIsland *Island, discoveredCells map[Cell]bool, grid [][]int) {
	r, c := startCell.Row, startCell.Col
	if r < 0 || r >= n || c < 0 || c >= m || grid[r][c] != 0 {
		return
	}
	// If this cell has already been discovered as part of an island, don't include it
	_, ok := discoveredCells[startCell]
	if ok {
		return
	}
	// Add the current cell to the island, and mark it as discovered
	*newIsland = append(*newIsland, startCell)
	discoveredCells[startCell] = true
	// Construct the top, bottom, left, and right sides of the island
	constructIsland(Cell{r-1, c}, n, m, newIsland, discoveredCells, grid)
	constructIsland(Cell{r+1, c}, n, m, newIsland, discoveredCells, grid)
	constructIsland(Cell{r, c-1}, n, m, newIsland, discoveredCells, grid)
	constructIsland(Cell{r, c+1}, n, m, newIsland, discoveredCells, grid)
}

func closedIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	n, m := len(grid), len(grid[0])
	var allIslands []Island
	discoveredCells := make(map[Cell]bool)
	for r, row := range grid {
		for c, col := range row {
			if col == 0 {
				newIsland := make(Island, 0)
				constructIsland(Cell{r,c}, n, m, &newIsland, discoveredCells, grid)
				if len(newIsland) > 0 {
					allIslands = append(allIslands, newIsland)
				}
			}
		}
	}
	sum := 0
	for _, island := range allIslands {
		if island.isClosed(n,m) {
			sum += 1
		}
	}
	return sum
}
