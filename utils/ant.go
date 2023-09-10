package utils

import (
	"math/rand"
	"sync"
)

type Ant struct {
	HasItem bool
	PosX    int
	PosY    int
}

func (ant *Ant) Init() {
	ant.HasItem = false
	ant.PosX = rand.Intn(MATRIZ_SIZE)
	ant.PosY = rand.Intn(MATRIZ_SIZE)
}

func MoveAnt(ant *Ant, ants *[]Ant, env *Enviroment, idx int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < NUMBER_ITERATIONS; i++ {
		move(ant, ants, env, idx)

	}

	// Move ant until all items are dropped in the enviroment
	for ant.HasItem {
		move(ant, ants, env, idx)
	}

}

func move(ant *Ant, ants *[]Ant, env *Enviroment, idx int) {

	vizinhos := neighbors(env, (*ant).PosX, (*ant).PosY)

	(*env).moveAnt(ant)

	pos_atual := (*env).GetCellValue((*ant).PosX, (*ant).PosY)
	if (*ant).HasItem && pos_atual == 0 {
		drop(ant, vizinhos, env)
	} else if !(*ant).HasItem && pos_atual == 1 {
		pick(ant, vizinhos, env)
	}
}

// Get offset of all directions from the range
func generateAllDirections() [][]int {
	directions := [][]int{}

	for x := ANT_RANGE * -1; x <= ANT_RANGE; x++ {
		for y := ANT_RANGE * -1; y <= ANT_RANGE; y++ {
			if x != 0 || y != 0 {
				directions = append(directions, []int{x, y})
			}
		}
	}
	return directions
}

// Get all neighbors of a cell using the directions
func neighbors(env *Enviroment, x, y int) [][]int {

	neighbors := [][]int{}
	rows, cols := env.GetSize(), env.GetSizeCol()

	directions := generateAllDirections()

	for _, dir := range directions {
		r, c := x+dir[0], y+dir[1]

		if r >= 0 && r < rows && c >= 0 && c < cols {
			neighbors = append(neighbors, []int{r, c})
		}
	}
	return neighbors
}

// Logic to pick a item from the enviroment
func pick(ant *Ant, v [][]int, env *Enviroment) {

	qtdVizinhos := len(v)
	numVizinhosComItem := 0

	(*env).mu.RLock()
	for i := 0; i < qtdVizinhos; i++ {

		if (*env).Map_items[v[i][0]][v[i][1]] == 1 {
			numVizinhosComItem++
		}
	}
	(*env).mu.RUnlock()

	calcProb := float32(1.0 - (float32(numVizinhosComItem) / float32(qtdVizinhos)))

	calcProb = (calcProb * calcProb)

	if calcProb > 0.9999 {
		(*env).setCellDec(ant.PosX, ant.PosY)
		(*ant).HasItem = true

	} else if rand.Float32() < calcProb {
		(*env).setCellDec(ant.PosX, ant.PosY)
		(*ant).HasItem = true
	}

}

// Logic to drop a item to the enviroment
// Drop has a higher probability to happen than pick
func drop(ant *Ant, v [][]int, env *Enviroment) {

	qtdVizinhos := len(v)
	numVizinhosComItem := 0

	(*env).mu.RLock()
	for i := 0; i < qtdVizinhos; i++ {
		if (*env).Map_items[v[i][0]][v[i][1]] == 1 {
			numVizinhosComItem++
		}
	}
	(*env).mu.RUnlock()

	calcProb := (float32(numVizinhosComItem) / float32(qtdVizinhos))
	calcProb = (calcProb * calcProb)

	if calcProb > 0.9999 {
		(*env).setCellIncre(ant.PosX, ant.PosY)
		(*ant).HasItem = false

	} else if rand.Float32() < calcProb {
		(*env).setCellIncre(ant.PosX, ant.PosY)
		(*ant).HasItem = false

	}

}
