package utils

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type Ant struct {
	HasItem bool
	PosX    int
	PosY    int
	Item    int // Item that the ant is carrying 1...N
}

func (ant *Ant) Init() {
	ant.HasItem = false
	ant.PosX = rand.Intn(MATRIZ_SIZE)
	ant.PosY = rand.Intn(MATRIZ_SIZE)
}

func MoveAnt(ant *Ant, ants *[]Ant, env *Enviroment, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < NUMBER_ITERATIONS; i++ {
		move(ant, ants, env)
		time.Sleep(time.Microsecond)
	}

	// Move ant until all items are dropped in the enviroment
	for ant.HasItem {
		move(ant, ants, env)
		time.Sleep(time.Microsecond)
	}

}

func move(ant *Ant, ants *[]Ant, env *Enviroment) {

	vizinhos := neighbors(env, (*ant).PosX, (*ant).PosY)

	(*env).moveAnt(ant)

	pos_atual := (*env).GetCellValue((*ant).PosX, (*ant).PosY)
	if (*ant).HasItem && pos_atual == 0 {
		drop(ant, vizinhos, env)
	} else if !(*ant).HasItem && pos_atual > 0 {
		pick(ant, vizinhos, env)
	}
}

// Get offset of all directions from the range
func generateAllDirections() [][]int {
	directions := [][]int{}

	for x := ANT_RANGE * -1; x <= ANT_RANGE; x++ {
		for y := ANT_RANGE * -1; y <= ANT_RANGE; y++ {
			if x != 0 || y != 0 {
				directions = append(directions, []int{x, y})
			}
		}
	}
	return directions
}

// Get all neighbors of a cell using the directions
func neighbors(env *Enviroment, x, y int) [][]int {

	neighbors := [][]int{}
	rows, cols := env.GetSize(), env.GetSizeCol()

	directions := generateAllDirections()

	for _, dir := range directions {
		r, c := x+dir[0], y+dir[1]

		if r >= 0 && r < rows && c >= 0 && c < cols {
			neighbors = append(neighbors, []int{r, c})
		}
	}
	return neighbors
}

func calcSimilaridade(v [][]int, env *Enviroment, items *[]Data, x, y int, vizinho_ret *int32) float64 {
	var similaridade float64
	var qtdDadosVizinhos int32
	var itemAtual Data
	similaridade = 0.0
	(*env).mu.RLock()
	if (*env).Map_items[x][y]-1 > 0 {
		itemAtual = (*items)[(*env).Map_items[x][y]-1]
	}

	for i := 0; i < len(v); i++ {
		valueFromCell := (*env).Map_items[v[i][0]][v[i][1]]
		if valueFromCell > 0 {
			qtdDadosVizinhos++
			itemInfo := (*items)[valueFromCell-1]
			quad := (math.Sqrt(math.Pow(itemAtual.PosX-itemInfo.PosX, 2) + math.Pow(itemAtual.PosY-itemInfo.PosY, 2)))
			similaridade += 1 - (quad / ALPHA)
		}
	}
	(*env).mu.RUnlock()
	*vizinho_ret = qtdDadosVizinhos
	if qtdDadosVizinhos <= 0 {
		return 0
	}
	similaridade = similaridade / math.Pow(float64(qtdDadosVizinhos), 2)

	if similaridade < 0.0 {
		//fmt.Println("Similaridade invalida: ", similaridade)
		return 0
	}
	/* if similaridade > 1.0 {
		//fmt.Println("Similaridade invalida: ", similaridade)
		return 1
	} */
	return similaridade

}

// Logic to pick a item from the enviroment
func pick(ant *Ant, v [][]int, env *Enviroment) {
	var vizinho_ret int32 = 0
	similaridade := calcSimilaridade(v, env, env.Items, ant.PosX, ant.PosY, &vizinho_ret)
	k1 := 0.01
	k1floor := k1 / (1.0 + k1)
	k1ceil := 1 - k1floor

	p_pick := k1 / (k1 + similaridade)

	if p_pick > k1floor && p_pick <= 1.0 {
		//    0.2  ...   1.0
		if (k1floor + rand.Float64()*k1ceil) >= p_pick {
			//pega
			(*ant).Item = env.GetCellValue(ant.PosX, ant.PosY)
			(*env).SetCellValue(ant.PosX, ant.PosY, 0)
			(*ant).HasItem = true
			fmt.Println("pick ", vizinho_ret, similaridade, p_pick)
		}
	} /* else {
		fmt.Println("Pick inválido ", similaridade, p_pick)
	} */
	return

}

// Logic to drop a item to the enviroment
// Drop has a higher probability to happen than pick
func drop(ant *Ant, v [][]int, env *Enviroment) {

	var vizinho_ret int32 = 0
	similaridade := calcSimilaridade(v, env, env.Items, ant.PosX, ant.PosY, &vizinho_ret)

	k2 := 0.01
	k2ceil := 1.0 / (k2 + 1.0)

	p_drop := (similaridade / (k2 + similaridade))

	if p_drop <= k2ceil && p_drop >= 0.0 {
		//    0.0  ... 0.917
		if (rand.Float64() * k2ceil) <= p_drop {
			//dropa
			(*env).SetCellValue(ant.PosX, ant.PosY, ant.Item)
			(*ant).HasItem = false
			(*ant).Item = 0
			fmt.Println("drop ", vizinho_ret, similaridade, p_drop)
		}
	} /* else {
		fmt.Println("Drop inválido ", similaridade, p_drop)
	} */

}
