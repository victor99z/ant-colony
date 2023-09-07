package utils

import (
	"fmt"
	"math/rand"
	"sync"
)

type Enviroment struct {
	mu        sync.RWMutex
	Map_items [][]int
	mutex_ant sync.RWMutex
	Map_ants  [][]int
}

func (env *Enviroment) Init() {
	env.Map_items = GenerateEnviroment()
	env.Map_ants = make([][]int, MATRIZ_SIZE)

	for i := range env.Map_ants {
		env.Map_ants[i] = make([]int, MATRIZ_SIZE)
	}

	fmt.Println(env.Map_ants)

}

func (env *Enviroment) moveAnt(ant *Ant) {
	env.mutex_ant.Lock()
	defer env.mutex_ant.Unlock()

	direcao_x := rand.Intn(3) - 1
	direcao_y := rand.Intn(3) - 1

	offsetDirecao := 0

	for (direcao_x == 0 && direcao_y == 0) ||

		((*ant).PosX+direcao_x >= MATRIZ_SIZE || (*ant).PosX+direcao_x < 0) ||
		((*ant).PosY+direcao_y >= MATRIZ_SIZE || (*ant).PosY+direcao_y < 0) ||

		(env.Map_ants[(*ant).PosX+direcao_x][(*ant).PosY+direcao_y] == 1) {

		offsetDirecao = offsetDirecao + 1
		if offsetDirecao > 10 {
			direcao_x = 0
			direcao_y = 0
			break
		}
		direcao_x = rand.Intn(3) - 1
		direcao_y = rand.Intn(3) - 1

	}

	if !(direcao_x == 0 && direcao_y == 0) {
		env.Map_ants[(*ant).PosX][(*ant).PosY] = 0
		(*ant).PosX = (*ant).PosX + direcao_x
		(*ant).PosY = (*ant).PosY + direcao_y
		env.Map_ants[(*ant).PosX][(*ant).PosY] = 1
	}
}

func (env *Enviroment) GetCellValue(x, y int) int {
	env.mu.RLock()
	defer env.mu.RUnlock()

	return env.Map_items[x][y]
}

func (env *Enviroment) SetCellValue(x, y, value int) {
	env.mu.Lock()
	defer env.mu.Unlock()

	env.Map_items[x][y] = value
}

func (env *Enviroment) setCellIncre(x, y int) {
	env.mu.Lock()
	defer env.mu.Unlock()

	env.Map_items[x][y] = env.Map_items[x][y] + 1
}

func (env *Enviroment) setCellDec(x, y int) {
	env.mu.Lock()
	defer env.mu.Unlock()

	env.Map_items[x][y] = env.Map_items[x][y] - 1
}

func (env *Enviroment) GetSize() int {
	return MATRIZ_SIZE
}

func (env *Enviroment) GetSizeCol() int {
	return MATRIZ_SIZE
}

func (env *Enviroment) GetAll() [][]int {
	return env.Map_items
}

func GenerateEnviroment() [][]int {

	if NUMBER_OF_ITEMS > MATRIZ_SIZE*MATRIZ_SIZE {
		panic("Too many items for this enviroment")
	}

	env := make([][]int, MATRIZ_SIZE)

	for i := range env {
		env[i] = make([]int, MATRIZ_SIZE)
	}

	for i := 0; i < NUMBER_OF_ITEMS; i++ {
		x := rand.Intn(MATRIZ_SIZE)
		y := rand.Intn(MATRIZ_SIZE)
		if env[x][y] == 0 {
			env[x][y] = 1
		} else {
			i--
		}
	}
	return env
}
