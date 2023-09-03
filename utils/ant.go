package utils

import (
	"math/rand"
	"sync"
)

type Ant struct {
	HasItem bool
	PosX    int
	posY    int
}

func (ant *Ant) Init() {
	ant.HasItem = false
	ant.PosX = rand.Intn(MATRIZ_SIZE)
	ant.posY = rand.Intn(MATRIZ_SIZE)
}

func (ant *Ant) MoveGo(env *[][]int, idx int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < NUMBER_ITERATIONS; i++ {
		ant.move(env)
		//fmt.Println("ant ", idx, ant)
	}
}

func (ant *Ant) move(env *[][]int) {
	// todo

	vizinhos := vizinhos(env, ant.PosX, ant.posY)
	qtdVizinhos := len(vizinhos)

	if ant.HasItem && (*env)[ant.PosX][ant.posY] == 0 {
		ant.drop(&vizinhos, env)
	} else if ant.HasItem && (*env)[ant.PosX][ant.posY] == 1 {
		ant.PosX = vizinhos[rand.Intn(qtdVizinhos)][0]
		ant.posY = vizinhos[rand.Intn(qtdVizinhos)][1]
	} else if !ant.HasItem && (*env)[ant.PosX][ant.posY] == 0 {
		ant.PosX = vizinhos[rand.Intn(qtdVizinhos)][0]
		ant.posY = vizinhos[rand.Intn(qtdVizinhos)][1]
	} else {
		ant.pick(&vizinhos, env)
	}

}

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

func vizinhos(env *[][]int, row, col int) [][]int {

	neighbors := [][]int{}
	rows, cols := len(*env), len((*env)[0])

	directions := generateAllDirections()

	for _, dir := range directions {
		r, c := row+dir[0], col+dir[1]

		// Check if the neighbor coordinates are within the matrix borders
		if r >= 0 && r < rows && c >= 0 && c < cols {
			neighbors = append(neighbors, []int{r, c})
		}
	}
	return neighbors
}

func (ant *Ant) pick(v *[][]int, env *[][]int) {
	// todo

	qtdVizinhos := len(*v)
	numVizinhosComItem := 0

	for _, v := range *v {
		if (*env)[v[0]][v[1]] == 1 {
			numVizinhosComItem++
		}
	}

	calcProb := (1 - (float32(numVizinhosComItem) / float32(qtdVizinhos))) * 100

	if calcProb == 100 {
		(*env)[ant.PosX][ant.posY] = 0
		ant.HasItem = true
	} else if rand.Intn(100) <= int(calcProb) || calcProb == 0 {
		(*env)[ant.PosX][ant.posY] = 1
		ant.HasItem = false
	} else {
		(*env)[ant.PosX][ant.posY] = 0
		ant.HasItem = true
	}

	ant.PosX = (*v)[rand.Intn(qtdVizinhos)][0]
	ant.posY = (*v)[rand.Intn(qtdVizinhos)][1]

}

func (ant *Ant) drop(v *[][]int, env *[][]int) {

	qtdVizinhos := len(*v)
	numVizinhosComItem := 0

	for _, v := range *v {
		if (*env)[v[0]][v[1]] == 1 {
			numVizinhosComItem++
		}
	}

	calcProb := (float32(numVizinhosComItem) / float32(qtdVizinhos)) * 100

	if calcProb == 100 {
		(*env)[ant.PosX][ant.posY] = 1
		ant.HasItem = false
	} else if rand.Intn(100) <= int(calcProb) || calcProb == 0 {
		(*env)[ant.PosX][ant.posY] = 0
		ant.HasItem = true
	} else {
		(*env)[ant.PosX][ant.posY] = 1
		ant.HasItem = false
	}

	ant.PosX = (*v)[rand.Intn(qtdVizinhos)][0]
	ant.posY = (*v)[rand.Intn(qtdVizinhos)][1]
}
