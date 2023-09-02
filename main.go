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

	utils.PrettyPrint(&enviroment)
	fmt.Println("")
	for i := 0; i < utils.NUMBER_ITERATIONS; i++ {

		for i := 0; i < len(ants); i++ {
			ants[i].Move(&enviroment)
		}
	}
	utils.PrettyPrint(&enviroment)
	utils.SaveToFile(&enviroment)
	// fmt.Println(ants)
	// fmt.Print(enviroment)
}
