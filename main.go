package main

import (
	"fmt"

	"github.com/victor99z/ant-colony/utils"
)

func main() {

	ants := make([]utils.Ant, utils.NUMBER_OF_ANTS)

	enviroment := utils.GenerateEnviroment()

	for ant := range ants {
		ants[ant].Init()
	}

	//utils.PrettyPrint(&enviroment)
	fmt.Println(ants)
	fmt.Println("")

	utils.SaveToFile(&enviroment, "input.csv")

	for i := 0; i < utils.NUMBER_ITERATIONS; i++ {
		for ant := 0; ant < len(ants); ant++ {
			ants[ant].Move(&enviroment)
		}
		//fmt.Println(ants)
	}

	for ant := 0; ant < len(ants); ant++ {
		for ants[ant].HasItem {
			for ant := 0; ant < len(ants); ant++ {
				ants[ant].Move(&enviroment)
			}
		}
	}

	for i, v := range ants {
		fmt.Println("Ant ", i, " has item: ", v.HasItem)
	}

	//utils.PrettyPrint(&enviroment)
	utils.SaveToFile(&enviroment, "output.csv")
	// fmt.Println(ants)
	// fmt.Print(enviroment)
}
