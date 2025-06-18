package game

type Point struct {
	X, Y int64
}

type Grid struct {
	cells map[Point]bool
}

func NewGrid() *Grid {
	return &Grid{
		cells: make(map[Point]bool),
	}
}

func (g *Grid) Set(x, y int64) {
	g.cells[Point{x, y}] = true
}

func (g *Grid) IsAlive(x, y int64) bool {
	return g.cells[Point{x, y}]
}

func (g *Grid) CountNeighbors(x, y int64) int {
	count := 0
	for dy := int64(-1); dy <= 1; dy++ {
		for dx := int64(-1); dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if g.IsAlive(x+dx, y+dy) {
				count++
			}
		}
	}
	return count
}

func (g *Grid) GetAllCellsToCheck() map[Point]bool {
	cellsToCheck := make(map[Point]bool)

	for cell := range g.cells {
		for dy := int64(-1); dy <= 1; dy++ {
			for dx := int64(-1); dx <= 1; dx++ {
				p := Point{cell.X + dx, cell.Y + dy}
				cellsToCheck[p] = true
			}
		}
	}

	return cellsToCheck
}

func (g *Grid) NextGeneration() *Grid {
	newGrid := NewGrid()
	cellsToCheck := g.GetAllCellsToCheck()

	for cell := range cellsToCheck {
		if g.ShouldBeAlive(cell) {
			newGrid.Set(cell.X, cell.Y)
		}
	}

	return newGrid
}

func (g *Grid) ShouldBeAlive(cell Point) bool {
	neighbors := g.CountNeighbors(cell.X, cell.Y)
	alive := g.IsAlive(cell.X, cell.Y)

	if alive && (neighbors == 2 || neighbors == 3) {
		return true
	} else if !alive && neighbors == 3 {
		return true
	}

	return false
}
