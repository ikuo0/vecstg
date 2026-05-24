package geom

var asciiPolylines [256]Glyph

// ASCIIコードに対応するPolylineを初期化
// 高さ1.0, 幅0.5の矩形を基本とし、必要に応じて線を追加していく

func initAsciiCharacters() {
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
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.0, Y: 0.0},
					{X: 0.0, Y: 0.0},
				},
			},
		},
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
	// M (0x4D)
	asciiPolylines[0x4D] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 1.000},
					{X: 0.000, Y: 0.000},
					{X: 0.250, Y: 0.500},
					{X: 0.500, Y: 0.000},
					{X: 0.500, Y: 1.000},
				},
			},
		},
	}
	// N (0x4E)
	asciiPolylines[0x4E] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 1.000},
					{X: 0.000, Y: 0.000},
					{X: 0.500, Y: 1.000},
					{X: 0.500, Y: 0.000},
				},
			},
		},
	}
	// O (0x4F)
	asciiPolylines[0x4F] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.125},
					{X: 0.125, Y: 0.000},
					{X: 0.375, Y: 0.000},
					{X: 0.500, Y: 0.125},
					{X: 0.500, Y: 0.875},
					{X: 0.375, Y: 1.000},
					{X: 0.125, Y: 1.000},
					{X: 0.000, Y: 0.875},
					{X: 0.000, Y: 0.125},
				},
			},
		},
	}
	// P (0x50)
	asciiPolylines[0x50] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 1.000},
					{X: 0.000, Y: 0.000},
					{X: 0.500, Y: 0.250},
					{X: 0.000, Y: 0.500},
				},
			},
		},
	}
	// Q (0x51)
	asciiPolylines[0x51] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.125},
					{X: 0.125, Y: 0.000},
					{X: 0.375, Y: 0.000},
					{X: 0.500, Y: 0.125},
					{X: 0.500, Y: 0.625},
					{X: 0.375, Y: 0.750},
					{X: 0.125, Y: 0.750},
					{X: 0.000, Y: 0.625},
					{X: 0.000, Y: 0.125},
				},
			},
			{
				Points: []Point{
					{X: 0.125, Y: 0.625},
					{X: 0.250, Y: 0.875},
					{X: 0.500, Y: 0.875},
				},
			},
		},
	}
	// R (0x52)
	asciiPolylines[0x52] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 1.000},
					{X: 0.000, Y: 0.000},
					{X: 0.500, Y: 0.250},
					{X: 0.000, Y: 0.500},
					{X: 0.500, Y: 1.000},
				},
			},
		},
	}
	// S (0x53)
	asciiPolylines[0x53] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.500, Y: 0.125},
					{X: 0.375, Y: 0.000},
					{X: 0.000, Y: 0.250},
					{X: 0.500, Y: 0.625},
					{X: 0.125, Y: 1.000},
					{X: 0.000, Y: 0.875},
				},
			},
		},
	}
	// T (0x54)
	asciiPolylines[0x54] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.500, Y: 0.000},
				},
			},
			{
				Points: []Point{
					{X: 0.250, Y: 0.000},
					{X: 0.250, Y: 1.000},
				},
			},
		},
	}
	// U (0x55)
	asciiPolylines[0x55] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.000, Y: 0.875},
					{X: 0.125, Y: 1.000},
					{X: 0.375, Y: 1.000},
					{X: 0.500, Y: 0.875},
					{X: 0.500, Y: 0.000},
				},
			},
		},
	}
	// V (0x56)
	asciiPolylines[0x56] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.250, Y: 1.000},
					{X: 0.500, Y: 0.000},
				},
			},
		},
	}
	// W (0x57)
	asciiPolylines[0x57] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.125, Y: 1.000},
					{X: 0.250, Y: 0.000},
					{X: 0.375, Y: 1.000},
					{X: 0.500, Y: 0.000},
				},
			},
		},
	}
	// X (0x58)
	asciiPolylines[0x58] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.375, Y: 1.000},
					{X: 0.500, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.000, Y: 1.000},
					{X: 0.375, Y: 0.000},
					{X: 0.500, Y: 0.000},
				},
			},
		},
	}
	// Y (0x59)
	asciiPolylines[0x59] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.250, Y: 0.375},
					{X: 0.250, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.250, Y: 0.375},
					{X: 0.500, Y: 0.000},
				},
			},
		},
	}
	// Z (0x5A)
	asciiPolylines[0x5A] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.500, Y: 0.000},
					{X: 0.000, Y: 1.000},
					{X: 0.500, Y: 1.000},
				},
			},
		},
	}
	// [ (0x5B)
	asciiPolylines[0x5B] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.500, Y: 0.000},
					{X: 0.250, Y: 0.000},
					{X: 0.250, Y: 1.000},
					{X: 0.500, Y: 1.000},
				},
			},
		},
	}
	// \ (0x5C)
	asciiPolylines[0x5C] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.500, Y: 1.000},
				},
			},
		},
	}
	// ] (0x5D)
	asciiPolylines[0x5D] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.250, Y: 0.000},
					{X: 0.250, Y: 1.000},
					{X: 0.000, Y: 1.000},
				},
			},
		},
	}
	// ^ (0x5E)
	asciiPolylines[0x5E] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.125, Y: 0.125},
					{X: 0.250, Y: 0.000},
					{X: 0.375, Y: 0.125},
				},
			},
		},
	}
	// _ (0x5F)
	asciiPolylines[0x5F] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 1.000},
					{X: 0.500, Y: 1.000},
				},
			},
		},
	}
	// ` (0x60)
	asciiPolylines[0x60] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.125, Y: 0.000},
					{X: 0.375, Y: 0.125},
				},
			},
		},
	}
	// a-z (0x61-0x7A)
	// a (0x61)
	asciiPolylines[0x61] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.125, Y: 0.375},
					{X: 0.375, Y: 0.375},
					{X: 0.500, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.375, Y: 0.500},
					{X: 0.125, Y: 0.750},
					{X: 0.375, Y: 0.875},
				},
			},
		},
	}
	// b (0x62)
	asciiPolylines[0x62] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.125, Y: 0.250},
					{X: 0.125, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.125, Y: 0.625},
					{X: 0.375, Y: 0.750},
					{X: 0.125, Y: 1.000},
				},
			},
		},
	}
	// c (0x63)
	asciiPolylines[0x63] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.500, Y: 0.500},
					{X: 0.250, Y: 0.500},
					{X: 0.125, Y: 0.625},
					{X: 0.125, Y: 0.875},
					{X: 0.250, Y: 1.000},
					{X: 0.500, Y: 1.000},
				},
			},
		},
	}
	// d (0x64)
	asciiPolylines[0x64] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.375, Y: 0.250},
					{X: 0.375, Y: 1.000},
					{X: 0.375, Y: 0.625},
					{X: 0.000, Y: 0.750},
					{X: 0.375, Y: 1.000},
				},
			},
		},
	}
	// e (0x65)
	asciiPolylines[0x65] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.750},
					{X: 0.500, Y: 0.625},
					{X: 0.000, Y: 0.500},
					{X: 0.000, Y: 1.000},
					{X: 0.500, Y: 1.000},
				},
			},
		},
	}
	// f (0x66)
	asciiPolylines[0x66] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.500, Y: 0.375},
					{X: 0.250, Y: 0.375},
					{X: 0.250, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.125, Y: 0.625},
					{X: 0.375, Y: 0.625},
				},
			},
		},
	}
	// g (0x67)
	asciiPolylines[0x67] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.375, Y: 0.625},
					{X: 0.125, Y: 0.625},
					{X: 0.125, Y: 0.375},
					{X: 0.375, Y: 0.375},
					{X: 0.375, Y: 1.000},
					{X: 0.125, Y: 0.875},
					{X: 0.500, Y: 0.875},
				},
			},
		},
	}
	// h (0x68)
	asciiPolylines[0x68] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.125, Y: 0.375},
					{X: 0.125, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.125, Y: 0.750},
					{X: 0.375, Y: 0.750},
					{X: 0.375, Y: 1.000},
				},
			},
		},
	}
	// i (0x69)
	asciiPolylines[0x69] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.250, Y: 0.375},
					{X: 0.250, Y: 0.500},
				},
			},
			{
				Points: []Point{
					{X: 0.250, Y: 0.625},
					{X: 0.250, Y: 1.000},
				},
			},
		},
	}
	// j (0x6A)
	asciiPolylines[0x6A] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.250, Y: 0.375},
					{X: 0.250, Y: 0.500},
				},
			},
			{
				Points: []Point{
					{X: 0.250, Y: 0.625},
					{X: 0.250, Y: 1.000},
					{X: 0.125, Y: 0.875},
				},
			},
		},
	}
	// k (0x6B)
	asciiPolylines[0x6B] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.125, Y: 0.375},
					{X: 0.125, Y: 1.000},
					{X: 0.125, Y: 0.875},
					{X: 0.375, Y: 0.625},
				},
			},
			{
				Points: []Point{
					{X: 0.375, Y: 1.000},
					{X: 0.250, Y: 0.875},
				},
			},
		},
	}
	// l (0x6C)
	asciiPolylines[0x6C] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.125, Y: 0.375},
					{X: 0.250, Y: 0.375},
					{X: 0.250, Y: 1.000},
				},
			},
		},
	}
	// m (0x6D)
	asciiPolylines[0x6D] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.500},
					{X: 0.125, Y: 0.500},
					{X: 0.125, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.125, Y: 0.500},
					{X: 0.375, Y: 0.500},
					{X: 0.375, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.250, Y: 0.500},
					{X: 0.250, Y: 1.000},
				},
			},
		},
	}
	// n (0x6E)
	asciiPolylines[0x6E] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.500},
					{X: 0.375, Y: 0.500},
					{X: 0.375, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.125, Y: 0.500},
					{X: 0.125, Y: 1.000},
				},
			},
		},
	}
	// o (0x6F)
	asciiPolylines[0x6F] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.625},
					{X: 0.125, Y: 0.500},
					{X: 0.375, Y: 0.500},
					{X: 0.500, Y: 0.625},
					{X: 0.500, Y: 0.875},
					{X: 0.375, Y: 1.000},
					{X: 0.125, Y: 1.000},
					{X: 0.000, Y: 0.875},
					{X: 0.000, Y: 0.625},
				},
			},
		},
	}
	// p (0x70)
	asciiPolylines[0x70] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.125, Y: 0.375},
					{X: 0.500, Y: 0.500},
					{X: 0.125, Y: 0.625},
				},
			},
			{
				Points: []Point{
					{X: 0.125, Y: 0.375},
					{X: 0.125, Y: 1.000},
				},
			},
		},
	}
	// q (0x71)
	asciiPolylines[0x71] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.375, Y: 0.375},
					{X: 0.000, Y: 0.500},
					{X: 0.375, Y: 0.625},
				},
			},
			{
				Points: []Point{
					{X: 0.375, Y: 0.375},
					{X: 0.375, Y: 1.000},
				},
			},
		},
	}
	// r (0x72)
	asciiPolylines[0x72] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.125, Y: 0.500},
					{X: 0.125, Y: 1.000},
					{X: 0.125, Y: 0.625},
					{X: 0.500, Y: 0.500},
				},
			},
		},
	}
	// s (0x73)
	asciiPolylines[0x73] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.500, Y: 0.375},
					{X: 0.125, Y: 0.625},
					{X: 0.500, Y: 0.750},
					{X: 0.125, Y: 1.000},
				},
			},
		},
	}
	// t (0x74)
	asciiPolylines[0x74] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.125, Y: 0.500},
					{X: 0.375, Y: 0.500},
				},
			},
			{
				Points: []Point{
					{X: 0.250, Y: 0.375},
					{X: 0.250, Y: 1.000},
					{X: 0.375, Y: 0.875},
				},
			},
		},
	}
	// u (0x75)
	asciiPolylines[0x75] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.500},
					{X: 0.000, Y: 0.875},
					{X: 0.125, Y: 1.000},
					{X: 0.250, Y: 1.000},
					{X: 0.375, Y: 0.875},
					{X: 0.375, Y: 0.500},
				},
			},
		},
	}
	// v (0x76)
	asciiPolylines[0x76] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.500},
					{X: 0.250, Y: 1.000},
					{X: 0.500, Y: 0.500},
				},
			},
		},
	}
	// w (0x77)
	asciiPolylines[0x77] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.500},
					{X: 0.125, Y: 1.000},
					{X: 0.250, Y: 0.500},
					{X: 0.375, Y: 1.000},
					{X: 0.500, Y: 0.500},
				},
			},
		},
	}
	// x (0x78)
	asciiPolylines[0x78] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.500},
					{X: 0.500, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.000, Y: 1.000},
					{X: 0.500, Y: 0.500},
				},
			},
		},
	}
	// y (0x79)
	asciiPolylines[0x79] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.500, Y: 0.375},
					{X: 0.125, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.125, Y: 0.375},
					{X: 0.250, Y: 0.750},
				},
			},
		},
	}
	// z (0x7A)
	asciiPolylines[0x7A] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.500},
					{X: 0.500, Y: 0.500},
					{X: 0.000, Y: 1.000},
					{X: 0.500, Y: 1.000},
				},
			},
		},
	}
	// { (0x7B)
	asciiPolylines[0x7B] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.500, Y: 0.000},
					{X: 0.250, Y: 0.125},
					{X: 0.250, Y: 0.875},
					{X: 0.500, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.125, Y: 0.500},
					{X: 0.250, Y: 0.500},
				},
			},
		},
	}
	// | (0x7C)
	asciiPolylines[0x7C] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.250, Y: 0.000},
					{X: 0.250, Y: 1.000},
				},
			},
		},
	}
	// } (0x7D)
	asciiPolylines[0x7D] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.000},
					{X: 0.250, Y: 0.125},
					{X: 0.250, Y: 0.875},
					{X: 0.000, Y: 1.000},
				},
			},
			{
				Points: []Point{
					{X: 0.250, Y: 0.500},
					{X: 0.375, Y: 0.500},
				},
			},
		},
	}
	// ~ (0x7E)
	asciiPolylines[0x7E] = Glyph{
		Polylines: []Polyline{
			{
				Points: []Point{
					{X: 0.000, Y: 0.625},
					{X: 0.125, Y: 0.375},
					{X: 0.250, Y: 0.625},
					{X: 0.375, Y: 0.375},
				},
			},
		},
	}
	// Extended ASCII (0x7F-0xFF)
	for i := 0x7F; i <= 0xFF; i++ {
		asciiPolylines[i] = undisplayablePoints
	}
}
