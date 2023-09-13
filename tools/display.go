package tools

import (
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/victor99z/ant-colony/utils"
)

const (
	SCREEN_HEIGHT = 640
	SCREEN_WIDTH  = 480
)

type Game struct {
	Enviroment *[][]int
	AntMap     *[][]int
}

func (g *Game) Update() error {

	if ebiten.IsWindowBeingClosed() {
		os.Exit(3)
	}

	return nil
}

// func agrupar( env *[][]int,int limitador){
// 	if limitador > 10 return;
// 	if (*g.Enviroment)[x][y] == 1
// 		return 1
// 	else return 0

// 	lista_itens = []
// 	int valor = agrupar((*g.Enviroment)[x-1][y-1], &lista_itens)
// 	valor += agrupar((*g.Enviroment)[x-1][y+1])
// 	valor += agrupar((*g.Enviroment)[x+1][y-1])
// 	valor += agrupar((*g.Enviroment)[x+1][y+1])

// 	if valor > 25
// 		for item in lista_itens:
// 			item.cor = sla
// }

func (g *Game) Draw(screen *ebiten.Image) {
	purpleCol := color.RGBA{193, 62, 130, 0.8 * 100} // Item
	greyColor := color.RGBA{128, 128, 128, 1}        // background
	antColor := color.RGBA{0, 0, 0, 1}               // Ant
	antAndItem := color.RGBA{0, 100, 0, 1}

	for x := 0; x < utils.MATRIZ_SIZE; x++ {
		for y := 0; y < utils.MATRIZ_SIZE; y++ {

			// var size2 float64 = (utils.MATRIZ_SIZE / 2.0)
			// var lo float64 = -3.1415 + (6.2831 * (size2 + float64(x)) / float64(utils.MATRIZ_SIZE))
			// var la float64 = -1.5707 + (3.1415 * (size2 + float64(y)) / float64(utils.MATRIZ_SIZE))
			// var x2 int = int(size2) + int(size2*math.Sin(lo)*math.Cos(la))
			// var y2 int = int(size2) + int(size2*math.Sin(lo)*math.Sin(la))
			// var z int = int(size2 * math.Cos(lo))

			// if z > 0 {
			// 	continue
			// }

			if (*g.Enviroment)[x][y] == 1 && (*g.AntMap)[x][y] == 1 {
				screen.Set(x, y, antAndItem)
			} else if (*g.Enviroment)[x][y] == 1 {
				screen.Set(x, y, purpleCol)
			} else if (*g.AntMap)[x][y] == 1 {
				screen.Set(x, y, antColor)
			} else {
				screen.Set(x, y, greyColor)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return utils.MATRIZ_SIZE, utils.MATRIZ_SIZE
}

func SetupDisplay(antMap, envMap *[][]int) {

	game := Game{
		Enviroment: envMap,
		AntMap:     antMap,
	}
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("2D matrix display")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}

}

func Debug(antMap, envMap *[][]int) {
	SetupDisplay(antMap, envMap)
}
