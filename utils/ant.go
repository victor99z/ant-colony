package utils

import (
	"math"
	"math/rand"
	"sync"
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

func (ant *Ant) MoveAnt(ants *[]Ant, env *Enviroment, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < NUMBER_ITERATIONS; i++ {
		ant.move(ants, env)
	}

	// Move ant until all items are dropped in the enviroment
	for ant.HasItem {
		ant.move(ants, env)
	}

}

func (ant *Ant) move(ants *[]Ant, env *Enviroment) {

	vizinhos := neighbors(env, (*ant).PosX, (*ant).PosY)

	(*env).moveAnt(ant)

	pos_atual := (*env).GetCellValue((*ant).PosX, (*ant).PosY)
	if (*ant).HasItem && pos_atual == 0 {
		ant.drop(vizinhos, env)
	} else if !(*ant).HasItem && pos_atual > 0 {
		ant.pick(vizinhos, env)
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

// FIXME: similaridade entre 0-0.499
// OBjetivo é 0.0 - 0.999

func calcSimilaridade(v [][]int, env *Enviroment, items *[]Data, ant *Ant, vizinho_ret *[]int32) float64 {
	var similaridade float64
	var qtdDadosVizinhos int32
	var itemAtual Data
	(*env).mu.RLock()

	if (*env).Map_items[ant.PosX][ant.PosY] > 0 {
		itemAtual = (*items)[(*env).Map_items[ant.PosX][ant.PosY]-1]
	} else if ant.HasItem {
		itemAtual = (*items)[ant.Item-1]
	} else {
		panic("Item atual invalido")
	}

	for i := 0; i < len(v); i++ {
		valueFromCell := (*env).Map_items[v[i][0]][v[i][1]]
		if valueFromCell > 0 {
			qtdDadosVizinhos++
			itemInfo := (*items)[valueFromCell-1]
			quad := (math.Sqrt(math.Pow(itemInfo.PosX-itemAtual.PosX, 2) + math.Pow(itemInfo.PosY-itemAtual.PosY, 2)))
			dist := 1 - (quad / ALPHA)
			if dist > 0 {
				similaridade += dist
			}
		}
	}
	(*env).mu.RUnlock()
	*vizinho_ret = []int32{int32(len(v)), qtdDadosVizinhos}
	if qtdDadosVizinhos <= 0 {
		return 0
	}

	similaridade = similaridade / float64(len(v))
	// similaridade = (similaridade / math.Pow(float64(qtdDadosVizinhos), 2))

	return math.Max(0.0, similaridade)

}

// Logic to pick a item from the enviroment
func (ant *Ant) pick(v [][]int, env *Enviroment) {
	var vizinho_ret []int32
	similaridade := calcSimilaridade(v, env, env.Items, ant, &vizinho_ret)

	// //p_pick := (utils.K1 / (utils.K1 + similaridade))
	p_pick := math.Pow((K1 / (K1 + similaridade)), 3)

	if similaridade >= K1 {
		p_pick = 0
	}

	if p_pick > 0.99 || rand.Float64() < p_pick {
		//pega
		(*ant).Item = env.GetCellValue(ant.PosX, ant.PosY)
		(*env).SetCellValue(ant.PosX, ant.PosY, 0)
		(*ant).HasItem = true
		//fmt.Println("pick ", vizinho_ret, similaridade, p_pick)
	}

	//remover classificação do item que vai tar na mão
}

// Logic to drop a item to the enviroment
// Drop has a higher probability to happen than pick
func (ant *Ant) drop(v [][]int, env *Enviroment) {

	var vizinho_ret []int32
	similaridade := calcSimilaridade(v, env, env.Items, ant, &vizinho_ret)

	//p_drop := (similaridade / (K2 + similaridade))
	p_drop := math.Pow((similaridade / (K2 + similaridade)), 3)

	if similaridade >= K2 {
		p_drop = 1
	}

	if p_drop > 0.99 || (rand.Float64()) < p_drop {
		//dropa
		(*env).SetCellValue(ant.PosX, ant.PosY, ant.Item)
		(*ant).HasItem = false
		(*ant).Item = 0
		//fmt.Println("drop ", vizinho_ret, similaridade, p_drop)
	}

	//atualizar classificação do item ao dropar
}
