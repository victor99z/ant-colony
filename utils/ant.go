package utils

import "math/rand"

type Ant struct {
	hasItem bool
	posX    int
	posY    int
}

func (ant *Ant) Init() {
	ant.hasItem = false
	ant.posX = rand.Intn(MATRIZ_SIZE)
	ant.posY = rand.Intn(MATRIZ_SIZE)
}

func (ant *Ant) Move(env *[][]int) {
	// todo
	if (*env)[ant.posX][ant.posY] == 1 {
		// Comportamento quando existe alguem na celula
		// dois estados, com item e sem item

		if ant.hasItem {
			// so move
			// ant.Move(env)
		} else {
			ant.pick()
			(*env)[ant.posX][ant.posY] = 0
		}

	} else {
		// Comportamento quando a celula está vazia
		// dois estados, com item e sem item

		if ant.hasItem {
			ant.drop()
			(*env)[ant.posX][ant.posY] = 1
		} else {
			// ant.Move(env)
		}

	}
}

func (ant *Ant) moveEstocastico(env *[][]int) {

	// (*env)[ant.posX+ANT_RANGE <= 10][ant.posY+ANT_RANGE]
	// (*env)[ant.posX+ANT_RANGE][ant.posY]
	// (*env)[ant.posX+ANT_RANGE][ant.posY]
	// (*env)[ant.posX+ANT_RANGE][ant.posY-ANT_RANGE]
	// (*env)[ant.posX-ANT_RANGE][ant.posY+ANT_RANGE]
	// (*env)[ant.posX-ANT_RANGE][ant.posY-ANT_RANGE]
	// (*env)[ant.posX][ant.posY-ANT_RANGE]
	// (*env)[ant.posX][ant.posY+ANT_RANGE]
}

func (ant *Ant) pick() bool {
	// todo
	ant.hasItem = true
	return ant.hasItem
}

func (ant *Ant) drop() bool {
	// todo
	ant.hasItem = false
	return ant.hasItem
}
