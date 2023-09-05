package main

import (
	"fmt"
	"sync"

	"github.com/victor99z/ant-colony/utils"
)

func CountItemsEnv(env *utils.Enviroment) int {
	count := 0
	for i := 0; i < utils.MATRIZ_SIZE; i++ {
		for j := 0; j < utils.MATRIZ_SIZE; j++ {
			if env.GetCellValue(i, j) != 0 {
				count++
			}
		}
	}
	return count
}

func main() {

	ants := make([]utils.Ant, utils.NUMBER_OF_ANTS)

	enviroment := utils.Enviroment{}
	enviroment.Init()

	fmt.Println("Start - Number of items in enviroment: ", CountItemsEnv(&enviroment))

	for ant := range ants {
		ants[ant].Init()
	}

	// utils.PrettyPrint(&enviroment)
	// fmt.Println(enviroment.GetAll(), ants)
	fmt.Println("")

	utils.SaveToFile(&enviroment, "input.csv")

	var wg sync.WaitGroup

	for i := 0; i < utils.NUMBER_OF_ANTS; i++ {
		wg.Add(1)
		go utils.MoveAnt(&ants[i], &enviroment, i, &wg)
	}
	wg.Wait()

	for i, v := range ants {
		defer fmt.Println("Ant ", i, " has item: ", v.HasItem)
	}

	defer fmt.Println("Final - Number of items in enviroment: ", CountItemsEnv(&enviroment))

	//defer utils.PrettyPrint(&enviroment)
	defer utils.SaveToFile(&enviroment, "output.csv")
	// fmt.Println(ants)
	// fmt.Print(enviroment)

}
