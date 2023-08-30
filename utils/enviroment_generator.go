package utils

import (
	"math/rand"
)

func GenerateEnviroment() [][]int {

	if NUMBER_OF_ITEMS > MATRIZ_SIZE*MATRIZ_SIZE {
		panic("Too many items for this enviroment")
	}

	enviroment := make([][]int, MATRIZ_SIZE)

	for i := range enviroment {
		enviroment[i] = make([]int, MATRIZ_SIZE)
	}

	for i := 0; i < NUMBER_OF_ITEMS; i++ {
		x := rand.Intn(MATRIZ_SIZE)
		y := rand.Intn(MATRIZ_SIZE)
		if enviroment[x][y] == 0 {
			enviroment[x][y] = 1
		} else {
			i--
		}
	}

	return enviroment
}
