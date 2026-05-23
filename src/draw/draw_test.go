package draw

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

type captureGame struct {
	width  int
	height int
	outPNG string
	done   bool
}

func (g *captureGame) Update() error {
	if g.done {
		return ebiten.Termination
	}

	img := ebiten.NewImage(g.width, g.height)
	img.Fill(color.RGBA{R: 8, G: 16, B: 24, A: 255})

	line := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	y := g.height / 2
	for x := g.width / 4; x < g.width*3/4; x++ {
		img.Set(x, y, line)
	}

	pix := make([]byte, 4*g.width*g.height)
	img.ReadPixels(pix)

	rgba := image.NewRGBA(image.Rect(0, 0, g.width, g.height))
	copy(rgba.Pix, pix)

	f, err := os.Create(g.outPNG)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := png.Encode(f, rgba); err != nil {
		return err
	}

	g.done = true
	return nil
}

func (g *captureGame) Draw(_ *ebiten.Image) {}

func (g *captureGame) Layout(_, _ int) (int, int) {
	return g.width, g.height
}

func TestRunGameOffscreenAndWritePNG(t *testing.T) {
	if runtime.GOOS == "linux" && os.Getenv("DISPLAY") == "" && os.Getenv("WAYLAND_DISPLAY") == "" {
		t.Skip("display server is not available; run with xvfb-run -a go test ./src/draw -run TestRunGameOffscreenAndWritePNG -v")
	}

	outPNG := filepath.Join(t.TempDir(), "frame.png")

	ebiten.SetWindowTitle("offscreen-test")
	ebiten.SetWindowSize(1, 1)
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowPosition(-32000, -32000)

	g := &captureGame{
		width:  64,
		height: 64,
		outPNG: outPNG,
	}

	err := ebiten.RunGameWithOptions(g, &ebiten.RunGameOptions{
		InitUnfocused: true,
	})
	if err != nil {
		t.Fatalf("RunGameWithOptions failed: %v", err)
	}

	f, err := os.Open(outPNG)
	if err != nil {
		t.Fatalf("failed to open generated PNG: %v", err)
	}
	defer f.Close()

	decoded, err := png.Decode(f)
	if err != nil {
		t.Fatalf("failed to decode generated PNG: %v", err)
	}

	r, gch, b, a := decoded.At(32, 32).RGBA()
	if r != 0xffff || gch != 0 || b != 0 || a != 0xffff {
		t.Fatalf("unexpected center pixel: got rgba=(%d,%d,%d,%d)", r, gch, b, a)
	}
}
