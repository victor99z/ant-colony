package main

import (
	"fmt"
	"sync"

	"github.com/victor99z/ant-colony/utils"
)

func main() {

	ants := make([]utils.Ant, utils.NUMBER_OF_ANTS)
	var wg sync.WaitGroup

	enviroment := utils.GenerateEnviroment()

	for ant := range ants {
		ants[ant].Init()
	}

	//utils.PrettyPrint(&enviroment)
	fmt.Println(ants)
	fmt.Println("")

	utils.SaveToFile(&enviroment, "input.csv")

	for i, v := range ants {
		wg.Add(1)
		go v.MoveGo(&enviroment, i, &wg)
	}

	wg.Wait()

	for i, v := range ants {
		defer fmt.Println("Ant ", i, " has item: ", v.HasItem)
	}

	//utils.PrettyPrint(&enviroment)
	utils.SaveToFile(&enviroment, "output.csv")
	// fmt.Println(ants)
	// fmt.Print(enviroment)

}
