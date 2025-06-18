package game

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func ToLife106Format(g *Grid) string {
	var result strings.Builder
	result.WriteString("#Life 1.06\n")

	array := make([]Point, 0, len(g.cells))

	for cell := range g.cells {
		array = append(array, cell)
	}

	sort.Slice(array, func(i, j int) bool {
		if array[i].X == array[j].X {
			return array[i].Y < array[j].Y
		}
		return array[i].X < array[j].X
	})

	for _, cell := range array {
		result.WriteString(fmt.Sprintf("%d %d\n", cell.X, cell.Y))
	}

	return result.String()
}

func FromLife106Format(scanner *bufio.Scanner) (*Grid, error) {
	grid := NewGrid()

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#Life 1.06") {
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}

		x, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid x coordinate: %s", parts[0])
		}

		y, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid y coordinate: %s", parts[1])
		}

		grid.Set(x, y)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}
