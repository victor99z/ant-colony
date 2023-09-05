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

var mutex sync.Mutex

func (ant *Ant) Init() {
	ant.HasItem = false
	ant.PosX = rand.Intn(MATRIZ_SIZE)
	ant.PosY = rand.Intn(MATRIZ_SIZE)
}
func MoveAnt(ant *Ant, env *Enviroment, idx int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < NUMBER_ITERATIONS; i++ {
		move(ant, env)
	}

}

func move(ant *Ant, env *Enviroment) {
	// todo

	vizinhos := neighbors(env, ant.PosX, ant.PosY)
	qtdVizinhos := len(vizinhos)

	randomFactor := rand.Intn(qtdVizinhos)

	localCellValue := (*env).GetCellValue(ant.PosX, ant.PosY)

	// if localCellValue != 0 && localCellValue != 1 {
	// 	fmt.Println(localCellValue)
	// }
	if ant.HasItem && localCellValue == 0 {
		drop(ant, &vizinhos, env)
	} else if ant.HasItem && localCellValue == 1 {

		ant.PosX = vizinhos[randomFactor][0]
		ant.PosY = vizinhos[randomFactor][1]

	} else if !ant.HasItem && localCellValue == 0 {

		ant.PosX = vizinhos[randomFactor][0]
		ant.PosY = vizinhos[randomFactor][1]

	} else if !ant.HasItem && localCellValue == 1 {
		pick(ant, &vizinhos, env)
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

func neighbors(env *Enviroment, row, col int) [][]int {

	neighbors := [][]int{}
	rows, cols := env.GetSize(), env.GetSizeCol()

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

func pick(ant *Ant, v *[][]int, env *Enviroment) {
	// todo

	qtdVizinhos := len(*v)
	numVizinhosComItem := 0

	for _, v := range *v {
		if (*env).GetCellValue(v[0], v[1]) == 1 {
			numVizinhosComItem++
		}
	}

	calcProb := (1 - (float32(numVizinhosComItem) / float32(qtdVizinhos))) * 100

	if calcProb == 100 {
		// env.SetCellValue(ant.PosX, ant.PosY, 0)
		env.setCellDec(ant.PosX, ant.PosY)
		ant.HasItem = true

	} else if rand.Intn(100) >= int(calcProb) || calcProb == 0 {
		// env.SetCellValue(ant.PosX, ant.PosY, 1)
		env.setCellIncre(ant.PosX, ant.PosY)
		ant.HasItem = false

	} else {
		// env.SetCellValue(ant.PosX, ant.PosY, 0)
		env.setCellDec(ant.PosX, ant.PosY)
		ant.HasItem = true

	}
	randomFactor := rand.Intn(qtdVizinhos)

	ant.PosX = (*v)[randomFactor][0]
	ant.PosY = (*v)[randomFactor][1]

}

func drop(ant *Ant, v *[][]int, env *Enviroment) {

	qtdVizinhos := len(*v)
	numVizinhosComItem := 0

	for _, v := range *v {
		if (*env).GetCellValue(v[0], v[1]) == 1 {
			numVizinhosComItem++
		}
	}

	calcProb := (float32(numVizinhosComItem) / float32(qtdVizinhos))

	if calcProb == 100 {
		// env.SetCellValue(ant.PosX, ant.PosY, 1)
		env.setCellIncre(ant.PosX, ant.PosY)
		ant.HasItem = false

	} else if rand.Intn(100) <= int(calcProb) || calcProb == 0 {
		// env.SetCellValue(ant.PosX, ant.PosY, 0)
		env.setCellDec(ant.PosX, ant.PosY)
		ant.HasItem = true

	} else {
		// env.SetCellValue(ant.PosX, ant.PosY, 1)
		env.setCellIncre(ant.PosX, ant.PosY)
		ant.HasItem = false

	}

	randomFactor := rand.Intn(qtdVizinhos)

	ant.PosX = (*v)[randomFactor][0]
	ant.PosY = (*v)[randomFactor][1]

}
