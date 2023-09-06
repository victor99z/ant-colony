package utils

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"sync"
)

type Ant struct {
	HasItem bool
	PosX    int
	PosY    int
	Id      [20]byte
}

var mutex sync.Mutex

func (ant *Ant) Init() {

	ant.HasItem = false
	ant.PosX = rand.Intn(MATRIZ_SIZE)
	ant.PosY = rand.Intn(MATRIZ_SIZE)
	ant.Id = sha1.Sum([]byte(fmt.Sprint(ant.PosX + ant.PosY)))

}

func MoveAnt(ant *Ant, ants *[]Ant, env *Enviroment, idx int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < NUMBER_ITERATIONS; i++ {

		move(ant, ants, env, idx)

	}
}

func move(ant *Ant, ants *[]Ant, env *Enviroment, idx int) {

	vizinhos := neighbors(env, ant.PosX, ant.PosY)
	qtdVizinhos := len(vizinhos)

	randomFactor := rand.Intn(qtdVizinhos)
	vizinhosRandomFactor := vizinhos[randomFactor]

	// for i := 0; i < len(*ants); i++ {
	// 	if ant.Id != (*ants)[i].Id && ant.PosX == (*ants)[i].PosX && ant.PosY == (*ants)[i].PosY {
	// 		// fmt.Print("Ant ", idx, " has collided with ant ", i, "\n")

	// 		ant.PosX = vizinhosRandomFactor[0]
	// 		ant.PosY = vizinhosRandomFactor[1]

	// 		return
	// 	}
	// }

	// if localCellValue != 0 && localCellValue != 1 {
	// 	fmt.Println(localCellValue)
	// }
	if ant.HasItem && (*env).GetCellValue(ant.PosX, ant.PosY) == 0 {
		drop(ant, vizinhos, env)
	} else if !ant.HasItem && (*env).GetCellValue(ant.PosX, ant.PosY) == 1 {
		pick(ant, vizinhos, env)
	}

	ant.PosX = vizinhosRandomFactor[0]
	ant.PosY = vizinhosRandomFactor[1]

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

func pick(ant *Ant, v [][]int, env *Enviroment) {
	// todo

	qtdVizinhos := len(v)
	numVizinhosComItem := 0

	// for _, v := range *v {
	// 	if (*env).GetCellValue(v[0], v[1]) == 1 {
	// 		numVizinhosComItem++
	// 	}
	// }

	for i := 0; i < qtdVizinhos; i++ {
		if (*env).GetCellValue(v[i][0], v[i][1]) == 1 {
			numVizinhosComItem++
		}
	}

	calcProb := (1 - (float32(numVizinhosComItem) / float32(qtdVizinhos))) * 100

	if calcProb == 100 {
		//env.SetCellValue(ant.PosX, ant.PosY, 0)
		env.setCellDec(ant.PosX, ant.PosY)
		ant.HasItem = true

	} else if rand.Intn(100) >= int(calcProb) {
		//env.SetCellValue(ant.PosX, ant.PosY, 0)
		env.setCellDec(ant.PosX, ant.PosY)
		ant.HasItem = true
	}

}

func drop(ant *Ant, v [][]int, env *Enviroment) {

	qtdVizinhos := len(v)
	numVizinhosComItem := 0

	for i := 0; i < qtdVizinhos; i++ {
		if (*env).GetCellValue(v[i][0], v[i][1]) == 1 {
			numVizinhosComItem++
		}
	}

	calcProb := (float32(numVizinhosComItem) / float32(qtdVizinhos)) * 100

	if calcProb == 100 {
		//env.SetCellValue(ant.PosX, ant.PosY, 0)
		env.setCellIncre(ant.PosX, ant.PosY)
		ant.HasItem = false

	} else if rand.Intn(100) <= int(calcProb) {
		//env.SetCellValue(ant.PosX, ant.PosY, 0)
		env.setCellIncre(ant.PosX, ant.PosY)
		ant.HasItem = false

	}

}
