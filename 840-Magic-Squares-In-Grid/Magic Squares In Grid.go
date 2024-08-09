package leetcode

func numMagicSquaresInside(grid [][]int) int {
	if len(grid) < 3 || len(grid[0]) < 3 {
		return 0
	}

	initialize(grid)

	result := 0
	if isMagic() {
		result = 1
	}
	for hasMore() {
		next()
		if isMagic() {
			result++
		}
	}
	return result
}

var square [][]int
var x int
var y int
var endX int
var endY int
var array []int

func initialize(grid [][]int) {
	square = grid
	x = 0
	y = 0
	array = make([]int, 16)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			array[square[i][j]]++
		}
	}

	endX = len(square[0]) - 3
	endY = len(square) - 3
}

func next() {
	if x != endX {
		for i := 0; i < 3; i++ {
			array[square[y+i][x]]--
			array[square[y+i][x+3]]++
		}
		x++
	} else {
		array = make([]int, 16)
		x = 0
		y++
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				array[square[y+i][j]]++
			}
		}
	}
}

func hasMore() bool {
	return x != endX || y != endY
}
func isMagic() bool {
	for i := 1; i < 10; i++ {
		if array[i] != 1 {
			return false
		}
	}

	for i := 0; i < 3; i++ {
		row := square[y+i][x] + square[y+i][x+1] + square[y+i][x+2]
		if row != 15 {
			return false
		}
		col := square[y][x+i] + square[y+1][x+i] + square[y+2][x+i]
		if col != 15 {
			return false
		}
	}

	diagonal := square[y][x] + square[y+1][x+1] + square[y+2][x+2]
	if diagonal != 15 {
		return false
	}

	diagonal = square[y][x+2] + square[y+1][x+1] + square[y+2][x]
	if diagonal != 15 {
		return false
	}

	return true
}
