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

//var mutex sync.Mutex

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

	for ant.HasItem {
		move(ant, ants, env, idx)
	}

}

func move(ant *Ant, ants *[]Ant, env *Enviroment, idx int) {

	// for i := 0; i < len(*ants); i++ {
	// 	if ant.Id != (*ants)[i].Id && ant.PosX == (*ants)[i].PosX && ant.PosY == (*ants)[i].PosY {
	// 		fmt.Print("Ant ", idx, " has collided with ant ", i, "\n")

	// 		ant.PosX = rand.Intn(MATRIZ_SIZE)
	// 		ant.PosY = rand.Intn(MATRIZ_SIZE)

	// 		return
	// 	}
	// }

	vizinhos := neighbors(env, (*ant).PosX, (*ant).PosY)
	//qtdVizinhos := len(vizinhos)

	//randomFactor := rand.Intn(qtdVizinhos)
	//vizinhosRandomFactor := vizinhos[randomFactor]

	// if localCellValue != 0 && localCellValue != 1 {
	// 	fmt.Println(localCellValue)
	// }

	(*env).moveAnt(ant)

	pos_atual := (*env).GetCellValue((*ant).PosX, (*ant).PosY)
	if (*ant).HasItem && pos_atual == 0 {
		drop(ant, vizinhos, env)
	} else if !(*ant).HasItem && pos_atual == 1 {
		pick(ant, vizinhos, env)
	}

	/* env.mutex_ant.Lock()
	env.antMap[ant.PosX][ant.PosY] = 0
	ant.PosX = vizinhosRandomFactor[0]
	ant.PosY = vizinhosRandomFactor[1]
	env.antMap[ant.PosX][ant.PosY] = 1
	env.mutex_ant.Unlock() */

	// PrettyPrint(&env.antMap)
	// fmt.Println("-----------------------")
	// PrettyPrint(&env.Map_items)
	// fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%")

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
		//env.SetCellValue(ant.PosX, ant.PosY, 0)
		(*env).setCellDec(ant.PosX, ant.PosY)
		(*ant).HasItem = true

		// fmt.Println(" pick item")
	} else if rand.Float32() < calcProb {
		//env.SetCellValue(ant.PosX, ant.PosY, 0)
		(*env).setCellDec(ant.PosX, ant.PosY)
		(*ant).HasItem = true

		// fmt.Print(ant.Id[0])
		// fmt.Println(" pick item")
	}

}

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

	//calcProb = (calcProb * calcProb * calcProb) + 0
	//calcProb = calcProb * 100

	if calcProb > 0.9999 {
		//env.SetCellValue(ant.PosX, ant.PosY, 0)
		(*env).setCellIncre(ant.PosX, ant.PosY)
		(*ant).HasItem = false

		// fmt.Println(" dropped item")
	} else if rand.Float32() < calcProb {
		//env.SetCellValue(ant.PosX, ant.PosY, 0)
		(*env).setCellIncre(ant.PosX, ant.PosY)
		(*ant).HasItem = false

		// fmt.Println(" dropped item")
	}

}
