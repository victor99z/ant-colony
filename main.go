package main

import (
	"fmt"
	"sync"

	"github.com/victor99z/ant-colony/utils"
)

func main() {

	ants := make([]utils.Ant, utils.NUMBER_OF_ANTS)

	enviroment := utils.Enviroment{}
	enviroment.Init()

	for ant := range ants {
		ants[ant].Init()
	}

	// utils.PrettyPrint(&enviroment)
	// fmt.Println(enviroment.GetAll(), ants)
	fmt.Println("")

	utils.SaveToFile(&enviroment, "input.csv")

	// for i, v := range ants {
	// 	wg.Add(1)
	// 	go utils.MoveAnt(&v, &enviroment, i, &wg)
	// }

	for i := 0; i < utils.NUMBER_ITERATIONS; i++ {
		var wg sync.WaitGroup

		for i := 0; i < utils.ANT_RANGE; i++ {
			wg.Add(1)
			utils.MoveAnt(&ants[i], &enviroment, i, &wg)
		}

		wg.Wait()
	}

	for i, v := range ants {
		defer fmt.Println("Ant ", i, " has item: ", v.HasItem)
	}

	// defer utils.PrettyPrint(&enviroment)
	defer utils.SaveToFile(&enviroment, "output.csv")
	// fmt.Println(ants)
	// fmt.Print(enviroment)

}
