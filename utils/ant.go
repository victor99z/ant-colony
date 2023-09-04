package utils

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
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

func (ant *Ant) MoveGo(env *[][]int, idx int, wg *sync.WaitGroup, c chan Position) {

	defer wg.Done()

	// defer func() { c <- Position{ant.PosX, ant.PosY} }()

	for i := 0; i < NUMBER_ITERATIONS; i++ {

		c <- Position{ant.PosX, ant.PosY}

		if pos := <-c; pos.posX != ant.PosX && pos.posY != ant.PosY {
			ant.move(env)
			fmt.Print("Passed1\n")
			time.Sleep(time.Second)
		}
	}

}

func (ant *Ant) move(env *[][]int) {
	// todo

	vizinhos := vizinhos(env, ant.PosX, ant.PosY)
	qtdVizinhos := len(vizinhos)

	if ant.HasItem && (*env)[ant.PosX][ant.PosY] == 0 {
		ant.drop(&vizinhos, env)
	} else if ant.HasItem && (*env)[ant.PosX][ant.PosY] == 1 {
		ant.PosX = vizinhos[rand.Intn(qtdVizinhos)][0]
		ant.PosY = vizinhos[rand.Intn(qtdVizinhos)][1]
	} else if !ant.HasItem && (*env)[ant.PosX][ant.PosY] == 0 {
		ant.PosX = vizinhos[rand.Intn(qtdVizinhos)][0]
		ant.PosY = vizinhos[rand.Intn(qtdVizinhos)][1]
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
		(*env)[ant.PosX][ant.PosY] = 0
		ant.HasItem = true
	} else if rand.Intn(100) <= int(calcProb) || calcProb == 0 {
		(*env)[ant.PosX][ant.PosY] = 1
		ant.HasItem = false
	} else {
		(*env)[ant.PosX][ant.PosY] = 0
		ant.HasItem = true
	}

	ant.PosX = (*v)[rand.Intn(qtdVizinhos)][0]
	ant.PosY = (*v)[rand.Intn(qtdVizinhos)][1]

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
		(*env)[ant.PosX][ant.PosY] = 1
		ant.HasItem = false
	} else if rand.Intn(100) <= int(calcProb) || calcProb == 0 {
		(*env)[ant.PosX][ant.PosY] = 0
		ant.HasItem = true
	} else {
		(*env)[ant.PosX][ant.PosY] = 1
		ant.HasItem = false
	}

	ant.PosX = (*v)[rand.Intn(qtdVizinhos)][0]
	ant.PosY = (*v)[rand.Intn(qtdVizinhos)][1]
}
