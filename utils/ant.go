package utils

import "math/rand"


type Ant struct {
	hasItem bool
	position[2] int
}

func (ant *Ant) Move(){
	// todo
	ant.position[0] = rand.Intn(MATRIZ_SIZE)
	ant.position[1] = rand.Intn(MATRIZ_SIZE)
}

func (ant *Ant) Pick(){
	// todo
	ant.hasItem = true
}

func (ant *Ant) Drop(){
	// todo
	ant.hasItem = false
}