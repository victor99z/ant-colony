package utils

import (
	"math/rand"
	"sync"
)

type Enviroment struct {
	mu        sync.Mutex
	map_items [][]int
}

func (env *Enviroment) Init() {
	env.map_items = GenerateEnviroment()
}

func (env *Enviroment) GetCellValue(x, y int) int {
	env.mu.Lock()
	defer env.mu.Unlock()

	return env.map_items[x][y]
}

func (env *Enviroment) SetCellValue(x, y, value int) {
	env.mu.Lock()
	defer env.mu.Unlock()

	env.map_items[x][y] = value
}

func (env *Enviroment) setCellIncre(x, y int) {
	env.mu.Lock()
	defer env.mu.Unlock()

	env.map_items[x][y] = env.map_items[x][y] + 1
}

func (env *Enviroment) setCellDec(x, y int) {
	env.mu.Lock()
	defer env.mu.Unlock()

	env.map_items[x][y] = env.map_items[x][y] - 1
}

func (env *Enviroment) GetSize() int {
	return MATRIZ_SIZE
}

func (env *Enviroment) GetSizeCol() int {
	return MATRIZ_SIZE
}

func (env *Enviroment) GetAll() [][]int {
	return env.map_items
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
