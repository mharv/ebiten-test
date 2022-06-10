package main

import (
	"errors"
	"image/color"

	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var gopherImg *ebiten.Image
var character Character

const (
	screenWidth  = 640
	screenHeight = 480
)

func init() {

}

type Game struct {
}

var regularTermination = errors.New("regular termination")

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return regularTermination
	}
	character.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0x10, 0xff})
	character.Draw(screen, &ebiten.DrawImageOptions{})

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Rumble pack")

	offsetX := 25.0
	offsetY := 0.0
	character.Init(offsetX, offsetY, "gopher.png")

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil && err != regularTermination {
		log.Fatal(err)
	}
}
