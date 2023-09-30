package tools

import (
	"image/color"
	"log"
	"math/rand"
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
	ItemsMap   []utils.Data
	Colors     []color.RGBA
}

func (g *Game) Update() error {

	if ebiten.IsWindowBeingClosed() {
		os.Exit(1)
	}

	return nil
}

func (g *Game) GenerateColors() {
	colors := []color.RGBA{}

	initialLabel := 0

	for i := 0; i < len(g.ItemsMap); i++ {
		if (g.ItemsMap)[i].Label > initialLabel {
			colors = append(colors, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 100})
			initialLabel = (g.ItemsMap)[i].Label
		}
	}

	g.Colors = append(g.Colors, colors...)
}

func (g *Game) Draw(screen *ebiten.Image) {
	//purpleCol := color.RGBA{193, 62, 130, 0.8 * 100} // Item
	whiteColor := color.RGBA{255, 255, 255, 3} // background
	// antColor := color.RGBA{0, 0, 0, 1}        // Ant

	for i := 0; i < utils.MATRIZ_SIZE; i++ {
		for j := 0; j < utils.MATRIZ_SIZE; j++ {
			// if (*g.AntMap)[i][j] == 1 {
			// 	screen.Set(i, j, antColor)
			if (*g.Enviroment)[i][j] > 0 {
				// get item from the list
				item := (*g.Enviroment)[i][j] - 1
				// get label from the item
				label := g.ItemsMap[item].Label - 1
				// get color that have the same label
				color := g.Colors[label]
				screen.Set(i, j, color)
			} else {
				screen.Set(i, j, whiteColor)
			}
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return utils.MATRIZ_SIZE, utils.MATRIZ_SIZE
}

func SetupDisplay(antMap, envMap *[][]int, items []utils.Data) {

	game := Game{
		Enviroment: envMap,
		AntMap:     antMap,
		ItemsMap:   items,
	}

	game.GenerateColors()

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("2D matrix display")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}

}

func Debug(antMap, envMap *[][]int, items []utils.Data) {
	SetupDisplay(antMap, envMap, items)
}
