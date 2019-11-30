package leetcode

type point []int

func computeNextCoord(x1, x2 int) int {
	var nextX int
	if x1 < x2 {
		nextX = x1 + 1
	} else if x1 == x2 {
		nextX = x1
	} else if x1 > x2 {
		nextX = x1 - 1
	}
	return nextX
}

func computeMinTravelTime(acc int, x, y point) int {
	if x[0] == y[0] && x[1] == y[1] {
		return acc
	}
	nextX := computeNextCoord(x[0], y[0])
	nextY := computeNextCoord(x[1], y[1])
	return computeMinTravelTime(acc+1, point{nextX, nextY}, y)
}

func minTimeToVisitAllPoints(points [][]int) int {
	prevPoint := points[0]
	totalTime := 0
	for i := 1; i < len(points); i++ {
		curPoint := points[i]
		totalTime += computeMinTravelTime(0, prevPoint, curPoint)
		prevPoint = curPoint
	}
	return totalTime
}
