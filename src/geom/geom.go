package geom

import (
	"math"
)

type Point struct {
	X float64
	Y float64
}

type Rect struct {
	Left   float64
	Top    float64
	Right  float64
	Bottom float64
}

type Polyline struct {
	Points []Point
}

type Glyph struct {
	Polylines []Polyline
}

type AspectRatio struct {
	Width  float64
	Height float64
}

type Viewport struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func (r Rect) Width() float64 {
	return r.Right - r.Left
}

func (r Rect) Height() float64 {
	return r.Bottom - r.Top
}

func (r Rect) ToLinePoints() []Point {
	return []Point{
		{X: r.Left, Y: r.Top},
		{X: r.Right, Y: r.Top},
		{X: r.Right, Y: r.Bottom},
		{X: r.Left, Y: r.Bottom},
		{X: r.Left, Y: r.Top},
	}
}

func (p Polyline) Start() Point {
	return p.Points[0]
}

func (p Polyline) LinePoints() []Point {
	return p.Points[1:]
}

func BoundingCenter(glyph Glyph) Point {
	var sumX, sumY float64
	var length int
	for _, polyline := range glyph.Polylines {
		for _, p := range polyline.Points {
			sumX += p.X
			sumY += p.Y
			length++
		}
	}

	return Point{X: sumX / float64(length), Y: sumY / float64(length)}
}

func BoundingBox(glyph Glyph) (box Rect) {
	if len(glyph.Polylines) == 0 {
		return Rect{0, 0, 0, 0}
	}
	minX, minY := glyph.Polylines[0].Points[0].X, glyph.Polylines[0].Points[0].Y
	maxX, maxY := glyph.Polylines[0].Points[0].X, glyph.Polylines[0].Points[0].Y
	for _, polyline := range glyph.Polylines {
		for _, p := range polyline.Points {
			if p.X < minX {
				minX = p.X
			}
			if p.Y < minY {
				minY = p.Y
			}
			if p.X > maxX {
				maxX = p.X
			}
			if p.Y > maxY {
				maxY = p.Y
			}
		}
	}
	return Rect{Left: minX, Top: minY, Right: maxX, Bottom: maxY}
}

func RotatePoint(p Point, angle float64) Point {
	rad := angle * (3.141592653589793 / 180.0)
	cos := math.Cos(rad)
	sin := math.Sin(rad)
	return Point{
		X: p.X*cos - p.Y*sin,
		Y: p.X*sin + p.Y*cos,
	}
}

func RotatePoints(center Point, points []Point, angle float64) []Point {
	rotated := make([]Point, len(points))
	for i, p := range points {
		// 中心を原点に移動
		translated := Point{X: p.X - center.X, Y: p.Y - center.Y}
		// 回転
		rotated[i] = RotatePoint(translated, angle)
		// 元の位置に戻す
		rotated[i].X += center.X
		rotated[i].Y += center.Y
	}
	return rotated
}

func ScalePoint(p Point, scaleX, scaleY float64) Point {
	return Point{
		X: p.X * scaleX,
		Y: p.Y * scaleY,
	}
}

func ScalePoints(center Point, points []Point, scaleX, scaleY float64) []Point {
	scaled := make([]Point, len(points))
	for i, p := range points {
		// 中心を原点に移動
		translated := Point{X: p.X - center.X, Y: p.Y - center.Y}
		// 拡大縮小
		scaled[i] = ScalePoint(translated, scaleX, scaleY)
		// 元の位置に戻す
		scaled[i].X += center.X
		scaled[i].Y += center.Y
	}
	return scaled
}

func GetFitScale(glyph Glyph, targetWidth, targetHeight float64) (scale float64) {
	box := BoundingBox(glyph)
	width := box.Width()
	height := box.Height()

	if width == 0 || height == 0 {
		return 1
	}

	scaleX := targetWidth / width
	scaleY := targetHeight / height

	if scaleX < scaleY {
		return scaleX
	}
	return scaleY
}

/*
指定した幅と高さに収まるように、Point配列をスケーリングする
元の比率を保つため、小さい方の倍率を返却する
*/
func GetFitScaleWithMargin(glyph Glyph, targetWidth, targetHeight, margin float64) (scale float64) {
	box := BoundingBox(glyph)
	width := box.Width()
	height := box.Height()

	if width == 0 || height == 0 {
		return 1
	}

	scaleX := (targetWidth - margin*2) / width
	scaleY := (targetHeight - margin*2) / height

	if scaleX < scaleY {
		return scaleX
	}
	return scaleY
}

/*
座標を移動させる
*/
func TranslatePoint(p Point, dx, dy float64) Point {
	return Point{
		X: p.X + dx,
		Y: p.Y + dy,
	}
}

func TranslatePoints(points []Point, dx, dy float64) []Point {
	translated := make([]Point, len(points))
	for i, p := range points {
		translated[i] = TranslatePoint(p, dx, dy)
	}
	return translated
}

func TranslateGlyph(glyph Glyph, dx, dy float64) Glyph {
	translatedPolylines := make([]Polyline, len(glyph.Polylines))
	for i, polyline := range glyph.Polylines {
		translatedPolylines[i] = Polyline{Points: TranslatePoints(polyline.Points, dx, dy)}
	}
	return Glyph{Polylines: translatedPolylines}
}
