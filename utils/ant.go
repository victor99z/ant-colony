package utils

import (
	"math/rand"
)

type Ant struct {
	hasItem bool
	posX    int
	posY    int
}

func (ant *Ant) Init() {
	ant.hasItem = false
	ant.posX = rand.Intn(MATRIZ_SIZE)
	ant.posY = rand.Intn(MATRIZ_SIZE)
}

func (ant *Ant) Move(env *[][]int) {
	// todo

	vizinhos := vizinhos(env, ant.posX, ant.posY)
	qtdVizinhos := len(vizinhos)

	if ant.hasItem && (*env)[ant.posX][ant.posY] == 0 {
		ant.drop(&vizinhos, env)
	} else if ant.hasItem && (*env)[ant.posX][ant.posY] == 1 {
		ant.posX = vizinhos[rand.Intn(qtdVizinhos)][0]
		ant.posY = vizinhos[rand.Intn(qtdVizinhos)][1]
	} else if !ant.hasItem && (*env)[ant.posX][ant.posY] == 0 {
		ant.posX = vizinhos[rand.Intn(qtdVizinhos)][0]
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
		(*env)[ant.posX][ant.posY] = 0
		ant.hasItem = true
	} else if rand.Intn(100) <= int(calcProb) || calcProb == 0 {
		(*env)[ant.posX][ant.posY] = 1
		ant.hasItem = false
	} else {
		(*env)[ant.posX][ant.posY] = 0
		ant.hasItem = true
	}

	ant.posX = (*v)[rand.Intn(qtdVizinhos)][0]
	ant.posY = (*v)[rand.Intn(qtdVizinhos)][1]

}

func (ant *Ant) drop(v *[][]int, env *[][]int) {
	// todo

	qtdVizinhos := len(*v)
	numVizinhosComItem := 0

	for _, v := range *v {
		if (*env)[v[0]][v[1]] == 1 {
			numVizinhosComItem++
		}
	}

	// so retorna 0 ou 1
	calcProb := (float32(numVizinhosComItem) / float32(qtdVizinhos)) * 100

	if calcProb == 100 {
		(*env)[ant.posX][ant.posY] = 1
		ant.hasItem = false
	} else if rand.Intn(100) <= int(calcProb) || calcProb == 0 {
		(*env)[ant.posX][ant.posY] = 0
		ant.hasItem = true
	} else {
		(*env)[ant.posX][ant.posY] = 1
		ant.hasItem = false
	}

	ant.posX = (*v)[rand.Intn(qtdVizinhos)][0]
	ant.posY = (*v)[rand.Intn(qtdVizinhos)][1]
}
