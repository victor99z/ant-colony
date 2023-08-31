package main

import (
	"fmt"
	"time"

	"github.com/victor99z/ant-colony/utils"
)

type Position struct {
	x int
	y int
}

func main() {

	ants := make([]utils.Ant, utils.NUMBER_OF_ANTS)

	enviroment := utils.GenerateEnviroment()

	for ant := range ants {
		ants[ant].Init()
	}

	for {
		fmt.Println("")
		utils.PrettyPrint(&enviroment)

		for _, v := range ants {
			v.Move(&enviroment)
		}

		time.Sleep(time.Second)
	}

	// fmt.Println(ants)
	// fmt.Print(enviroment)
}
