package utils

import (
	"math/rand"
)

// Generate a 100x100 enviroment using 2d array of int and insert random values between 0 and 1

func GenerateEnviroment(size int) [][] int{

	enviroment := make([][]int, size)
	
	for i := range enviroment {
		enviroment[i] = make([]int, size)
		for j := range enviroment[i] {
			enviroment[i][j] = rand.Intn(2)
		}
	}
	return enviroment
}