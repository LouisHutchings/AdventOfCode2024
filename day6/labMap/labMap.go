package labMap

import (
	"errors"
)

const (
	Up = iota
	Right
	Down
	Left
)

type LabMap struct {
	Walls     []map[int]bool
	Guard     guard
	MapBounds []int
}

type guard struct {
	x         int
	y         int
	direction int
}

func (lm *LabMap) GetGuardPosition() ([]int, error) {
	return []int{lm.Guard.x, lm.Guard.y, Up}, nil
}

func (lm *LabMap) MoveGuard() ([]int, error) {
	directionToNewPosition := map[int][]int{
		Up:    {lm.Guard.x, lm.Guard.y - 1},
		Right: {lm.Guard.x + 1, lm.Guard.y},
		Down:  {lm.Guard.x, lm.Guard.y + 1},
		Left:  {lm.Guard.x - 1, lm.Guard.y},
	}
	newPosition := directionToNewPosition[lm.Guard.direction]
	if newPosition[0] > lm.MapBounds[0] || newPosition[1] > lm.MapBounds[1] ||
		newPosition[0] < 0 || newPosition[1] < 0 {
		return newPosition, errors.New("guard moves out of map")
	}
	for lm.isWall(newPosition[0], newPosition[1]) {
		lm.Guard.turnRight()
		newPosition = directionToNewPosition[lm.Guard.direction]
	}
	lm.Guard.updatePosition(newPosition[0], newPosition[1])
	return append(newPosition, lm.Guard.direction), nil
}

func (g *guard) updatePosition(x, y int) {
	g.x = x
	g.y = y
}
func (g *guard) turnRight() {
	g.direction = (g.direction + 1) % 4
}
func (lm *LabMap) isWall(x int, y int) bool {

	return lm.Walls[y][x]
}

func New(wallLocations [][]int, guardLocation []int, rows int) *LabMap {
	newMap := LabMap{}
	newMap.initWalls(wallLocations, rows)
	newMap.initGuard(guardLocation)
	newMap.MapBounds = []int{rows - 1, rows - 1}
	return &newMap
}

func (lm *LabMap) initWalls(wallLocations [][]int, rows int) {
	lm.Walls = make([]map[int]bool, rows)
	for i := 0; i < rows; i++ {
		lm.Walls[i] = make(map[int]bool)
	}
	for _, wallLocation := range wallLocations {
		lm.Walls[wallLocation[0]][wallLocation[1]] = true
	}
}

func (lm *LabMap) initGuard(guardLocation []int) {
	lm.Guard = guard{x: guardLocation[0], y: guardLocation[1], direction: Up}
}
