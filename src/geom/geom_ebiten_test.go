package geom

// ebiten 依存の描画テストコード
// geom.Point を使用して線を描画し PNG 出力を行う
// xvfb-run -a go test ./src/geom -run TestDrawPolyline -v

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const ScreenWidth = 640
const ScreenHeight = 480

type polylineCaptureGame struct {
	glyph    Glyph
	fileName string
	err      error
	done     bool
}

func Execute(t *testing.T, glyph Glyph, fileName string) {
	g := &polylineCaptureGame{
		glyph:    glyph,
		fileName: fileName,
	}

	err := ebiten.RunGameWithOptions(g, &ebiten.RunGameOptions{InitUnfocused: true})
	if err != nil {
		t.Fatalf("RunGameWithOptions failed: %v", err)
	}
	if g.err != nil {
		t.Fatalf("Failed to render/save image: %v", g.err)
	}
}

func TestDrawPolyline(t *testing.T) {
	Execute(
		t,
		Glyph{Polylines: []Polyline{
			{Points: []Point{
				{X: 200, Y: 100},
				{X: 100, Y: 200},
				{X: 300, Y: 200},
				{X: 200, Y: 100},
			}},
		}},
		fmt.Sprintf("%s.png", t.Name()),
	)
}

// xvfb-run -a go test ./src/geom -run TestDrawNumber2 -v
// 数字の「2」を描くテスト
func TestDrawNumber2(t *testing.T) {
	points := []Point{
		{X: 0.0, Y: 0.0},
		{X: 0.5, Y: 0.0},
		{X: 0.5, Y: 0.5},
		{X: 0.0, Y: 0.5},
		{X: 0.0, Y: 1.0},
		{X: 0.5, Y: 1.0},
	}
	numberGlyph := Glyph{Polylines: []Polyline{{Points: points}}}

	// 文字を25行で描くときの高さを計算
	box := BoundingBox(numberGlyph)
	heightPx := ScreenHeight / 25.0

	numberScale := GetFitScaleWithMargin(numberGlyph, ScreenWidth, heightPx, heightPx*0.2) // 20%のマージンを追加
	numberCenter := Point{X: (box.Left + box.Right) / 2, Y: (box.Top + box.Bottom) / 2}
	numberScaled := ScalePoints(numberCenter, points, numberScale, numberScale)
	numberTranslated := TranslatePoints(numberScaled, ScreenWidth/2-numberCenter.X, ScreenHeight/2-numberCenter.Y)
	numberGlyph = Glyph{Polylines: []Polyline{{Points: numberTranslated}}}

	boxGlyph := Glyph{Polylines: []Polyline{{Points: box.ToLinePoints()}}}
	outboxScale := GetFitScale(boxGlyph, ScreenWidth, heightPx)
	outboxPoints := box.ToLinePoints()
	outboxScaled := ScalePoints(numberCenter, outboxPoints, outboxScale, outboxScale)
	outboxTranslated := TranslatePoints(outboxScaled, ScreenWidth/2-numberCenter.X, ScreenHeight/2-numberCenter.Y)
	outboxGlyph := Glyph{Polylines: []Polyline{{Points: outboxTranslated}}}

	glyph := Glyph{Polylines: append(numberGlyph.Polylines, outboxGlyph.Polylines...)}

	Execute(
		t,
		glyph,
		fmt.Sprintf("%s.png", t.Name()),
	)
}

func (g *polylineCaptureGame) Update() error {
	if g.done {
		return ebiten.Termination
	}

	img := ebiten.NewImage(ScreenWidth, ScreenHeight)
	img.Fill(color.Black)

	line := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	for _, polyline := range g.glyph.Polylines {
		drawPolyline(img, polyline, line)
	}

	if err := saveImageAsPNG(img, g.fileName); err != nil {
		g.err = err
	}

	g.done = true
	return nil
}

func (g *polylineCaptureGame) Draw(_ *ebiten.Image) {}

func (g *polylineCaptureGame) Layout(_, _ int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func geomOutDir() (string, error) {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to resolve caller file path")
	}

	return filepath.Join(filepath.Dir(thisFile), "out"), nil
}

func drawPolyline(img *ebiten.Image, polyline Polyline, clr color.Color) {
	if len(polyline.Points) < 2 {
		return
	}

	var path vector.Path
	path.MoveTo(float32(polyline.Points[0].X), float32(polyline.Points[0].Y))
	for _, p := range polyline.Points[1:] {
		path.LineTo(float32(p.X), float32(p.Y))
	}

	strokeOp := &vector.StrokeOptions{Width: 1}
	drawOp := &vector.DrawPathOptions{AntiAlias: false}
	drawOp.ColorScale.ScaleWithColor(clr)
	vector.StrokePath(img, &path, strokeOp, drawOp)
}

func saveImageAsPNG(img *ebiten.Image, fileName string) error {
	outDir, err := geomOutDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return err
	}

	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	pix := make([]byte, 4*w*h)
	img.ReadPixels(pix)

	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	copy(rgba.Pix, pix)

	outPath := filepath.Join(outDir, fileName)
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, rgba)
}
