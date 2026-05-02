package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// 背景を黒で塗る
	screen.Fill(color.Black)

	// 画面中央に長さ200pxの赤い横線を描く
	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}

	x1 := screenWidth/2 - 100
	y := screenHeight / 2

	for x := x1; x < x1+200; x++ {
		screen.Set(x, y, red)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebiten Sample")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
