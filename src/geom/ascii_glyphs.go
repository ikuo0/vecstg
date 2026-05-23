package geom

var asciiPolylines [256]Glyph

// ASCIIコードに対応するPolylineを初期化
// 高さ1.0, 幅0.5の矩形を基本とし、必要に応じて線を追加していく

func init() {
	// undisplayable characters (0x00-0x1F) が指定されたらバッテンを表示する
	undisplayablePoints := Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0, Y: 0},
					{X: 0.5, Y: 1.0},
				},
			},
			{
				Points: []Point{
					{X: 0.5, Y: 0},
					{X: 0.0, Y: 1.0},
				},
			},
		},
	}
	// 0x00 NULL
	asciiPolylines[0x00] = undisplayablePoints
	// 0x01 SOH
	asciiPolylines[0x01] = undisplayablePoints

}
