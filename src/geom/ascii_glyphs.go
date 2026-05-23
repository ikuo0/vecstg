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
	// 0x02 STX
	asciiPolylines[0x02] = undisplayablePoints
	// 0x03 ETX
	asciiPolylines[0x03] = undisplayablePoints
	// 0x04 EOT
	asciiPolylines[0x04] = undisplayablePoints
	// 0x05 ENQ
	asciiPolylines[0x05] = undisplayablePoints
	// 0x06 ACK
	asciiPolylines[0x06] = undisplayablePoints
	// 0x07 BEL
	asciiPolylines[0x07] = undisplayablePoints
	// 0x08 BS
	asciiPolylines[0x08] = undisplayablePoints
	// 0x09 HT
	asciiPolylines[0x09] = undisplayablePoints
	// 0x0A LF
	asciiPolylines[0x0A] = undisplayablePoints
	// 0x0B VT
	asciiPolylines[0x0B] = undisplayablePoints
	// 0x0C FF
	asciiPolylines[0x0C] = undisplayablePoints
	// 0x0D CR
	asciiPolylines[0x0D] = undisplayablePoints
	// 0x0E SO
	asciiPolylines[0x0E] = undisplayablePoints
	// 0x0F SI
	asciiPolylines[0x0F] = undisplayablePoints
	// 0x10 DLE
	asciiPolylines[0x10] = undisplayablePoints
	// 0x11 DC1
	asciiPolylines[0x11] = undisplayablePoints
	// 0x12 DC2
	asciiPolylines[0x12] = undisplayablePoints
	// 0x13 DC3
	asciiPolylines[0x13] = undisplayablePoints
	// 0x14 DC4
	asciiPolylines[0x14] = undisplayablePoints
	// 0x15 NAK
	asciiPolylines[0x15] = undisplayablePoints
	// 0x16 SYN
	asciiPolylines[0x16] = undisplayablePoints
	// 0x17 ETB
	asciiPolylines[0x17] = undisplayablePoints
	// 0x18 CAN
	asciiPolylines[0x18] = undisplayablePoints
	// 0x19 EM
	asciiPolylines[0x19] = undisplayablePoints
	// 0x1A SUB
	asciiPolylines[0x1A] = undisplayablePoints
	// 0x1B ESC
	asciiPolylines[0x1B] = undisplayablePoints
	// 0x1C FS
	asciiPolylines[0x1C] = undisplayablePoints
	// 0x1D GS
	asciiPolylines[0x1D] = undisplayablePoints
	// 0x1E RS
	asciiPolylines[0x1E] = undisplayablePoints
	// 0x1F US
	asciiPolylines[0x1F] = undisplayablePoints

	// space (0x20)
	asciiPolylines[0x20] = Glyph{
		Polylines: []Polyline{},
	}
	// ! (0x21)
	asciiPolylines[0x21] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.250, Y: 0.000},
					{X: 0.250, Y: 0.625},
				},
			},
			{
				Points: []Point{
					{X: 0.250, Y: 0.750},
					{X: 0.250, Y: 1.000},
				},
			},
		},
	}
	// " (0x22)
	asciiPolylines[0x22] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.125, Y: 0.000},
					{X: 0.125, Y: 0.250},
				},
			},
			{
				Points: []Point{
					{X: 0.375, Y: 0.000},
					{X: 0.375, Y: 0.250},
				},
			},
		},
	}
	// # (0x23)
	asciiPolylines[0x23] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.125, Y: 0.000},
					{X: 0.125, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.375, Y: 0.000},
					{X: 0.375, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.000, Y: 0.250},
					{X: 0.500, Y: 0.250},
				},
			},
			{
				Points: []Point{
					{X: 0.000, Y: 0.625},
					{X: 0.500, Y: 0.625},
				},
			},
		},
	}
	// $ (0x24)
	asciiPolylines[0x24] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.500, Y: 0.000},
					{X: 0.000, Y: 0.375},
					{X: 0.500, Y: 0.625},
					{X: 0.000, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.125, Y: 0.000},
					{X: 0.125, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.375, Y: 0.000},
					{X: 0.375, Y: 1.000},
				},
			},
		},
	}
	// % (0x25)
	asciiPolylines[0x25] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.250},
					{X: 0.000, Y: 0.125},
					{X: 0.125, Y: 0.125},
					{X: 0.125, Y: 0.250},
					{X: 0.000, Y: 0.250},
				},
			},
			{
				Points: []Point{
					{X: 0.500, Y: 0.000},
					{X: 0.000, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.375, Y: 0.875},
					{X: 0.375, Y: 0.750},
					{X: 0.500, Y: 0.750},
					{X: 0.500, Y: 0.875},
					{X: 0.375, Y: 0.875},
				},
			},
		},
	}
	// & (0x26)
	asciiPolylines[0x26] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.500, Y: 0.875},
					{X: 0.000, Y: 0.250},
					{X: 0.125, Y: 0.000},
					{X: 0.250, Y: 0.250},
					{X: 0.000, Y: 0.750},
					{X: 0.125, Y: 1.000},
					{X: 0.500, Y: 0.500},
				},
			},
		},
	}
	// ' (0x27) // シングルクォーテーション
	asciiPolylines[0x27] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.250, Y: 0.000},
					{X: 0.125, Y: 0.250},
				},
			},
		},
	}
	// ( (0x28)
	asciiPolylines[0x28] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.250, Y: 0.000},
					{X: 0.000, Y: 0.250},
					{X: 0.000, Y: 0.750},
					{X: 0.250, Y: 1.000},
				},
			},
		},
	}
	// ) (0x29)
	asciiPolylines[0x29] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.250, Y: 0.000},
					{X: 0.500, Y: 0.250},
					{X: 0.500, Y: 0.750},
					{X: 0.250, Y: 1.000},
				},
			},
		},
	}
	// * (0x2A)
	asciiPolylines[0x2A] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.250, Y: 0.000},
					{X: 0.250, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.000, Y: 0.500},
					{X: 0.500, Y: 0.500},
				},
			},
			{
				Points: []Point{
					{X: 0.000, Y: 0.125},
					{X: 0.500, Y: 0.875},
				},
			},
			{
				Points: []Point{
					{X: 0.000, Y: 0.875},
					{X: 0.500, Y: 0.125},
				},
			},
		},
	}
	// + (0x2B)
	asciiPolylines[0x2B] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.250, Y: 0.000},
					{X: 0.250, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.000, Y: 0.500},
					{X: 0.500, Y: 0.500},
				},
			},
		},
	}
	// , (0x2C)
	asciiPolylines[0x2C] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.250, Y: 0.750},
					{X: 0.125, Y: 1.000},
				},
			},
		},
	}
	// - (0x2D)
	asciiPolylines[0x2D] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.500},
					{X: 0.500, Y: 0.500},
				},
			},
		},
	}
	// . (0x2E)
	asciiPolylines[0x2E] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.125, Y: 0.875},
					{X: 0.125, Y: 1.000},
				},
			},
		},
	}
	// / (0x2F)
	asciiPolylines[0x2F] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.375, Y: 0.000},
					{X: 0.125, Y: 1.000},
				},
			},
		},
	}

	// 0-9
	// 0 (0x30)
	asciiPolylines[0x30] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.250, Y: 0.000},
					{X: 0.000, Y: 0.250},
					{X: 0.000, Y: 0.750},
					{X: 0.250, Y: 1.000},
					{X: 0.500, Y: 0.750},
					{X: 0.500, Y: 0.250},
					{X: 0.250, Y: 0.000},
				},
			},
		},
	}
	// 1 (0x31)
	asciiPolylines[0x31] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.125, Y: 0.125},
					{X: 0.250, Y: 0.000},
					{X: 0.250, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.125, Y: 1.000},
					{X: 0.375, Y: 1.000},
				},
			},
		},
	}
	// 2 (0x32)
	asciiPolylines[0x32] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.500, Y: 0.000},
					{X: 0.500, Y: 0.500},
					{X: 0.000, Y: 0.500},
					{X: 0.000, Y: 1.000},
					{X: 0.500, Y: 1.000},
				},
			},
		},
	}
	// 3 (0x33)
	asciiPolylines[0x33] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.500, Y: 0.000},
					{X: 0.500, Y: 1.000},
					{X: 0.000, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.000, Y: 0.500},
					{X: 0.500, Y: 0.500},
				},
			},
		},
	}
	// 4 (0x34)
	asciiPolylines[0x34] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.000, Y: 0.500},
					{X: 0.500, Y: 0.500},
				},
			},
			{
				Points: []Point{
					{X: 0.500, Y: 0.000},
					{X: 0.500, Y: 1.000},
				},
			},
		},
	}
	// 5 (0x35)
	asciiPolylines[0x35] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.500, Y: 0.000},
					{X: 0.000, Y: 0.000},
					{X: 0.000, Y: 0.500},
					{X: 0.500, Y: 0.500},
					{X: 0.500, Y: 1.000},
					{X: 0.000, Y: 1.000},
				},
			},
		},
	}
	// 6 (0x36)
	asciiPolylines[0x36] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.500, Y: 0.000},
					{X: 0.000, Y: 0.000},
					{X: 0.000, Y: 1.000},
					{X: 0.500, Y: 1.000},
					{X: 0.500, Y: 0.500},
					{X: 0.000, Y: 0.500},
				},
			},
		},
	}
	// 7 (0x37)
	asciiPolylines[0x37] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.250},
					{X: 0.000, Y: 0.000},
					{X: 0.500, Y: 0.000},
					{X: 0.500, Y: 1.000},
				},
			},
		},
	}
	// 8 (0x38)
	asciiPolylines[0x38] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.500, Y: 0.000},
					{X: 0.000, Y: 0.000},
					{X: 0.000, Y: 1.000},
					{X: 0.500, Y: 1.000},
					{X: 0.500, Y: 0.000},
				},
			},
			{
				Points: []Point{
					{X: 0.000, Y: 0.500},
					{X: 0.500, Y: 0.500},
				},
			},
		},
	}
	// 9 (0x39)
	asciiPolylines[0x39] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 1.000},
					{X: 0.500, Y: 1.000},
					{X: 0.500, Y: 0.000},
					{X: 0.000, Y: 0.000},
					{X: 0.000, Y: 0.500},
					{X: 0.500, Y: 0.500},
				},
			},
		},
	}

	// : (0x3A)
	asciiPolylines[0x3A] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.250, Y: 0.125},
					{X: 0.250, Y: 0.250},
				},
			},
			{
				Points: []Point{
					{X: 0.250, Y: 0.625},
					{X: 0.250, Y: 0.750},
				},
			},
		},
	}
	// ; (0x3B)
	asciiPolylines[0x3B] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.250, Y: 0.125},
					{X: 0.250, Y: 0.250},
				},
			},
			{
				Points: []Point{
					{X: 0.250, Y: 0.625},
					{X: 0.125, Y: 0.875},
				},
			},
		},
	}
	// < (0x3C)
	asciiPolylines[0x3C] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.375, Y: 0.125},
					{X: 0.000, Y: 0.500},
					{X: 0.375, Y: 0.875},
				},
			},
		},
	}
	// = (0x3D)
	asciiPolylines[0x3D] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.750},
					{X: 0.500, Y: 0.750},
				},
			},
			{
				Points: []Point{
					{X: 0.000, Y: 0.375},
					{X: 0.500, Y: 0.375},
				},
			},
		},
	}
	// > (0x3E)
	asciiPolylines[0x3E] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.125, Y: 0.125},
					{X: 0.500, Y: 0.500},
					{X: 0.125, Y: 0.875},
				},
			},
		},
	}
	// ? (0x3F)
	asciiPolylines[0x3F] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.500, Y: 0.000},
					{X: 0.500, Y: 0.375},
					{X: 0.250, Y: 0.375},
					{X: 0.250, Y: 0.625},
				},
			},
			{
				Points: []Point{
					{X: 0.250, Y: 0.750},
					{X: 0.250, Y: 0.875},
				},
			},
		},
	}
	// @ (0x40)
	asciiPolylines[0x40] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.375, Y: 0.625},
					{X: 0.375, Y: 0.250},
					{X: 0.125, Y: 0.250},
					{X: 0.125, Y: 0.750},
					{X: 0.500, Y: 0.750},
					{X: 0.500, Y: 0.000},
					{X: 0.000, Y: 0.000},
					{X: 0.000, Y: 0.875},
					{X: 0.125, Y: 1.000},
					{X: 0.375, Y: 1.000},
				},
			},
		},
	}
	// A-Z (0x41-0x5A)
	// A (0x41)
	asciiPolylines[0x41] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 1.000},
					{X: 0.250, Y: 0.000},
					{X: 0.500, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.000, Y: 0.625},
					{X: 0.375, Y: 0.625},
				},
			},
		},
	}
	// B (0x42)
	asciiPolylines[0x42] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.500, Y: 0.250},
					{X: 0.000, Y: 0.500},
					{X: 0.500, Y: 0.750},
					{X: 0.000, Y: 1.000},
					{X: 0.000, Y: 0.000},
				},
			},
		},
	}
	// C (0x43)
	asciiPolylines[0x43] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.500, Y: 0.125},
					{X: 0.250, Y: 0.000},
					{X: 0.000, Y: 0.125},
					{X: 0.000, Y: 0.875},
					{X: 0.250, Y: 1.000},
					{X: 0.500, Y: 0.875},
				},
			},
		},
	}
	// D (0x44)
	asciiPolylines[0x44] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.250, Y: 0.000},
					{X: 0.500, Y: 0.250},
					{X: 0.500, Y: 0.750},
					{X: 0.250, Y: 1.000},
					{X: 0.000, Y: 1.000},
					{X: 0.000, Y: 0.000},
				},
			},
		},
	}
	// E (0x45)
	asciiPolylines[0x45] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.500, Y: 0.000},
					{X: 0.000, Y: 0.000},
					{X: 0.000, Y: 1.000},
					{X: 0.500, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.000, Y: 0.500},
					{X: 0.500, Y: 0.500},
				},
			},
		},
	}
	// F (0x46)
	asciiPolylines[0x46] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 1.000},
					{X: 0.000, Y: 0.000},
					{X: 0.500, Y: 0.000},
				},
			},
			{
				Points: []Point{
					{X: 0.000, Y: 0.375},
					{X: 0.500, Y: 0.375},
				},
			},
		},
	}
	// G (0x47)
	asciiPolylines[0x47] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.500, Y: 0.125},
					{X: 0.375, Y: 0.000},
					{X: 0.125, Y: 0.000},
					{X: 0.000, Y: 0.125},
					{X: 0.000, Y: 0.875},
					{X: 0.125, Y: 1.000},
					{X: 0.500, Y: 1.000},
					{X: 0.500, Y: 0.625},
					{X: 0.250, Y: 0.625},
				},
			},
		},
	}
	// H (0x48)
	asciiPolylines[0x48] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.000, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.000, Y: 0.500},
					{X: 0.500, Y: 0.500},
				},
			},
			{
				Points: []Point{
					{X: 0.500, Y: 0.000},
					{X: 0.500, Y: 1.000},
				},
			},
		},
	}
	// I (0x49)
	asciiPolylines[0x49] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.125, Y: 0.000},
					{X: 0.375, Y: 0.000},
				},
			},
			{
				Points: []Point{
					{X: 0.250, Y: 0.000},
					{X: 0.250, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.125, Y: 1.000},
					{X: 0.375, Y: 1.000},
				},
			},
		},
	}
	// J (0x4A)
	asciiPolylines[0x4A] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.250, Y: 0.000},
					{X: 0.500, Y: 0.000},
					{X: 0.500, Y: 0.750},
					{X: 0.250, Y: 1.000},
					{X: 0.000, Y: 0.750},
				},
			},
		},
	}
	// K (0x4B)
	asciiPolylines[0x4B] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.000, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.500, Y: 0.000},
					{X: 0.000, Y: 0.500},
					{X: 0.500, Y: 1.000},
				},
			},
		},
	}
	// L (0x4C)
	asciiPolylines[0x4C] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.000, Y: 1.000},
					{X: 0.500, Y: 1.000},
				},
			},
		},
	}

}
