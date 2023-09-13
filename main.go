package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/victor99z/ant-colony/tools"
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

	f, err := utils.ReadData("./test_cases/case_4_groups.csv")

	fmt.Println(f, err)

	time.Sleep(time.Minute)
	// Inicializa a lista de formigas
	ants := make([]utils.Ant, utils.NUMBER_OF_ANTS)
	var wg sync.WaitGroup

	enviroment := utils.Enviroment{}
	enviroment.Init()
	fmt.Println("Start - Number of items in enviroment: ", CountItemsEnv(&enviroment))

	for ant := range ants {
		ants[ant].Init()
	}

	utils.SaveToFile(&enviroment, "input.csv")

	if utils.DEBUG {
		go tools.Debug(&enviroment.Map_ants, &enviroment.Map_items)
		time.Sleep(time.Second * 5)
	}

	for i := 0; i < utils.NUMBER_OF_ANTS; i++ {
		wg.Add(1)
		go utils.MoveAnt(&ants[i], &ants, &enviroment, i, &wg)
	}

	wg.Wait()

	for i, v := range ants {
		defer fmt.Println("Ant ", i, " has item: ", v.HasItem)
	}

	defer fmt.Println("Final - Number of items in enviroment: ", CountItemsEnv(&enviroment))

	// Loop apenas para uso no debug evitando que a engine feche a janela apos a execução, assim podemos analisar o resultado final
	if utils.DEBUG {
		defer func() {
			for {
			}
		}()
	}

	//defer utils.PrettyPrint(&enviroment.map)

	// fmt.Println(ants)
	// fmt.Print(enviroment)

}
