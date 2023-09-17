package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/victor99z/ant-colony/tools"
	"github.com/victor99z/ant-colony/utils"
)

func main() {

	items := utils.ParseData("./test_cases/case_4_groups.csv")

	// Inicializa a lista de formigas
	ants := make([]utils.Ant, utils.NUMBER_OF_ANTS)
	var wg sync.WaitGroup

	enviroment := utils.Enviroment{}
	enviroment.Init(&items)

	fmt.Println(len(items))

	for ant := range ants {
		ants[ant].Init()
	}

	if utils.DEBUG {
		go tools.Debug(&enviroment.Map_ants, &enviroment.Map_items, items)
		time.Sleep(time.Second * 3)
	}

	for i := 0; i < utils.NUMBER_OF_ANTS; i++ {
		wg.Add(1)
		go utils.MoveAnt(&ants[i], &ants, &enviroment, &wg)
	}

	wg.Wait()

	// for i, v := range ants {
	// 	defer fmt.Println("Ant ", i, " has item: ", v.HasItem)
	// }

	// Loop apenas para uso no debug evitando que a engine feche a janela apos a execução, assim podemos analisar o resultado final
	if utils.DEBUG {
		defer func() {
			select {}
		}()
	}

}
