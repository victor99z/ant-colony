package tools

import (
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/victor99z/ant-colony/utils"
)

const (
	SCREEN_HEIGHT = 640
	SCREEN_WIDTH  = 480
)

type Game struct {
	mu         sync.RWMutex
	Enviroment *[][]int
	AntMap     *[][]int
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	purpleCol := color.RGBA{193, 62, 130, 0.8 * 100} // Item
	greyColor := color.RGBA{128, 128, 128, 1}
	antColor := color.RGBA{0, 0, 0, 1}
	antAndItem := color.RGBA{0, 1, 0, 1}

	g.mu.RLock()
	for x := 0; x < utils.MATRIZ_SIZE; x++ {
		for y := 0; y < utils.MATRIZ_SIZE; y++ {

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
	g.mu.RUnlock()

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
	ebiten.SetWindowSize(utils.MATRIZ_SIZE*5, utils.MATRIZ_SIZE*5)
	ebiten.SetWindowTitle("2D matrix display")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}

}
