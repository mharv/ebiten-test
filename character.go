package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Character struct {
	posX   float64
	posY   float64
	width  int
	height int
	sprite *ebiten.Image
	buffer []string
}

func (c *Character) Init(x, y float64, filepath string) {

	img, _, err := ebitenutil.NewImageFromFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	c.posX = x
	c.posY = y
	c.sprite = img
	c.width, c.height = c.sprite.Size()
	c.buffer = []string{}
}

func (c *Character) Update() {

	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		c.buffer = []string{}
	}

	for _, p := range inpututil.PressedKeys() {
		c.buffer = append(c.buffer, p.String())
	}
}

func randomVibration(buffer []string) (float64, float64) {

	multiplier := len(buffer)

	seed := rand.Float64()
	var sign int
	if seed > 0.5 {
		sign = 1
	} else {
		sign = -1
	}

	xNoise := rand.Float64() * float64(multiplier) * 0.1 * float64(sign)
	yNoise := rand.Float64() * float64(multiplier) * 0.1 * float64(sign)

	return xNoise, yNoise
}

func (c *Character) Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions) {

	xNoise, yNoise := randomVibration(c.buffer)

	c.width += int(xNoise)
	c.height += int(yNoise)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.posX+xNoise, c.posY+yNoise)

	red := 0.0
	multiplier := 0.02
	red = float64(len(c.buffer)) * multiplier
	op.ColorM.Scale(red, 0.5, 1, 1)

	screen.DrawImage(c.sprite, op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("POWER LEVEL: %d", len(c.buffer)))
}
