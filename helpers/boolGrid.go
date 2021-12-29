package helpers

type BoolGrid [][]bool

func NewBoolGrid(width, height int, init bool) BoolGrid {
	grid := make([][]bool, height)
	for y := range grid {
		grid[y] = make([]bool, width)
		for x := range grid[y] {
			grid[y][x] = init
		}
	}
	return grid
}

func (grid *BoolGrid) Set(x, y int, value bool) {
	(*grid)[y][x] = value
}

func (grid *BoolGrid) Get(x, y int) (bool, bool) {
	if x < 0 || y < 0 || x >= len((*grid)[0]) || y >= len(*grid) {
		return false, false
	}
	return (*grid)[y][x], true
}
