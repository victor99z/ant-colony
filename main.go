package main

import (
	"fmt"

	"github.com/victor99z/ant-colony/utils"
)




func prettyPrint(environment *[][]int){
	for i := range *environment {
		for j := range (*environment)[i] {
			fmt.Print((*environment)[i][j])
		}
		fmt.Println()
	}
}

func main(){

	ants := make([]utils.Ant, utils.NUMBER_OF_ANTS)

	enviroment := utils.GenerateEnviroment(utils.MATRIZ_SIZE)


	prettyPrint(&enviroment)

	fmt.Println(ants)
	// fmt.Print(enviroment)
}