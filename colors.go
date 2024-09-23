package csskit

import (
	"fmt"
	"image/color"
)

const (
	ColorSlate   = "slate"
	ColorGray    = "gray"
	ColorZinc    = "zinc"
	ColorNeutral = "neutral"
	ColorStone   = "stone"
	ColorRed     = "red"
	ColorOrange  = "orange"
	ColorAmber   = "amber"
	ColorYellow  = "yellow"
	ColorLime    = "lime"
	ColorGreen   = "green"
	ColorEmerald = "emerald"
	ColorTeal    = "teal"
	ColorCyan    = "cyan"
	ColorSky     = "sky"
	ColorBlue    = "blue"
	ColorIndigo  = "indigo"
	ColorViolet  = "violet"
	ColorPurple  = "purple"
	ColorFuchsia = "fuchsia"
	ColorPink    = "pink"
	ColorRose    = "rose"
)

var AllColorNames = []string{
	ColorSlate, ColorGray, ColorZinc, ColorNeutral,
	ColorStone, ColorRed, ColorOrange, ColorAmber,
	ColorYellow, ColorLime, ColorGreen, ColorEmerald,
	ColorTeal, ColorCyan, ColorSky, ColorBlue,
	ColorIndigo, ColorViolet, ColorPurple, ColorFuchsia,
	ColorPink, ColorRose,
}

var shadeCount int

func init() {
	shadeCount = len(Shades)
}

var Shades = []int{50, 100, 200, 300, 400, 500, 600, 700, 800, 900, 950}

func getClosestShades(num int) (int, int) {
	if num <= Shades[0] {
		return -1, -1
	}
	if num >= Shades[shadeCount-1] {
		return -1, -1
	}
	for i := 0; i < shadeCount-1; i++ {
		if num >= Shades[i] && num <= Shades[i+1] {
			return Shades[i], Shades[i+1]
		}
	}
	return -1, -1
}

func interpolateNRGBA(a, b color.NRGBA, t float64) color.NRGBA {
	return color.NRGBA{
		R: uint8(float64(a.R)*(1-t) + float64(b.R)*t),
		G: uint8(float64(a.G)*(1-t) + float64(b.G)*t),
		B: uint8(float64(a.B)*(1-t) + float64(b.B)*t),
		A: uint8(float64(a.A)*(1-t) + float64(b.A)*t),
	}
}

func getColor(name string, shade float64) color.NRGBA {
	prevShade, nextShade := getClosestShades(int(shade))
	if prevShade == -1 || nextShade == -1 {
		panic(fmt.Errorf("color shade out of bounds (50-950): %f", shade))
	}
	shadeMap, ok := Colors[name]
	if !ok {
		panic(fmt.Errorf("unknown color name: %s", name))
	}
	prevColor := shadeMap[prevShade]
	nextColor := shadeMap[nextShade]
	mixf := (shade - float64(prevShade)) / float64(nextShade - prevShade)
	return interpolateNRGBA(prevColor, nextColor, mixf)
}

var Colors = map[string]map[int]color.NRGBA{
	ColorSlate: {
		50:  color.NRGBA{R: 0xF8, G: 0xFA, B: 0xFC, A: 255},
		100: color.NRGBA{R: 0xF1, G: 0xF5, B: 0xF9, A: 255},
		200: color.NRGBA{R: 0xE2, G: 0xE8, B: 0xF0, A: 255},
		300: color.NRGBA{R: 0xCB, G: 0xD5, B: 0xE1, A: 255},
		400: color.NRGBA{R: 0x94, G: 0xA3, B: 0xB8, A: 255},
		500: color.NRGBA{R: 0x64, G: 0x74, B: 0x8B, A: 255},
		600: color.NRGBA{R: 0x47, G: 0x55, B: 0x69, A: 255},
		700: color.NRGBA{R: 0x33, G: 0x41, B: 0x55, A: 255},
		800: color.NRGBA{R: 0x1E, G: 0x29, B: 0x3B, A: 255},
		900: color.NRGBA{R: 0x0F, G: 0x17, B: 0x2A, A: 255},
		950: color.NRGBA{R: 0x02, G: 0x06, B: 0x17, A: 255},
	},
	ColorGray: {
		50:  color.NRGBA{R: 0xF9, G: 0xFA, B: 0xFB, A: 255},
		100: color.NRGBA{R: 0xF3, G: 0xF4, B: 0xF6, A: 255},
		200: color.NRGBA{R: 0xE5, G: 0xE7, B: 0xEB, A: 255},
		300: color.NRGBA{R: 0xD1, G: 0xD5, B: 0xDB, A: 255},
		400: color.NRGBA{R: 0x9C, G: 0xA3, B: 0xAF, A: 255},
		500: color.NRGBA{R: 0x6B, G: 0x72, B: 0x80, A: 255},
		600: color.NRGBA{R: 0x4B, G: 0x55, B: 0x63, A: 255},
		700: color.NRGBA{R: 0x37, G: 0x41, B: 0x51, A: 255},
		800: color.NRGBA{R: 0x1F, G: 0x29, B: 0x37, A: 255},
		900: color.NRGBA{R: 0x11, G: 0x18, B: 0x27, A: 255},
		950: color.NRGBA{R: 0x03, G: 0x07, B: 0x12, A: 255},
	},
	ColorZinc: {
		50:  color.NRGBA{R: 0xFA, G: 0xFA, B: 0xFA, A: 255},
		100: color.NRGBA{R: 0xF4, G: 0xF4, B: 0xF5, A: 255},
		200: color.NRGBA{R: 0xE4, G: 0xE4, B: 0xE7, A: 255},
		300: color.NRGBA{R: 0xD4, G: 0xD4, B: 0xD8, A: 255},
		400: color.NRGBA{R: 0xA1, G: 0xA1, B: 0xAA, A: 255},
		500: color.NRGBA{R: 0x71, G: 0x71, B: 0x7A, A: 255},
		600: color.NRGBA{R: 0x52, G: 0x52, B: 0x5B, A: 255},
		700: color.NRGBA{R: 0x3F, G: 0x3F, B: 0x46, A: 255},
		800: color.NRGBA{R: 0x27, G: 0x27, B: 0x2A, A: 255},
		900: color.NRGBA{R: 0x18, G: 0x18, B: 0x1B, A: 255},
		950: color.NRGBA{R: 0x09, G: 0x09, B: 0x0B, A: 255},
	},
	ColorNeutral: {
		50:  color.NRGBA{R: 0xFA, G: 0xFA, B: 0xFA, A: 255},
		100: color.NRGBA{R: 0xF5, G: 0xF5, B: 0xF5, A: 255},
		200: color.NRGBA{R: 0xE5, G: 0xE5, B: 0xE5, A: 255},
		300: color.NRGBA{R: 0xD4, G: 0xD4, B: 0xD4, A: 255},
		400: color.NRGBA{R: 0xA3, G: 0xA3, B: 0xA3, A: 255},
		500: color.NRGBA{R: 0x73, G: 0x73, B: 0x73, A: 255},
		600: color.NRGBA{R: 0x52, G: 0x52, B: 0x52, A: 255},
		700: color.NRGBA{R: 0x40, G: 0x40, B: 0x40, A: 255},
		800: color.NRGBA{R: 0x26, G: 0x26, B: 0x26, A: 255},
		900: color.NRGBA{R: 0x17, G: 0x17, B: 0x17, A: 255},
		950: color.NRGBA{R: 0x0A, G: 0x0A, B: 0x0A, A: 255},
	},
	ColorStone: {
		50:  color.NRGBA{R: 0xFA, G: 0xFA, B: 0xF9, A: 255},
		100: color.NRGBA{R: 0xF5, G: 0xF5, B: 0xF4, A: 255},
		200: color.NRGBA{R: 0xE7, G: 0xE5, B: 0xE4, A: 255},
		300: color.NRGBA{R: 0xD6, G: 0xD3, B: 0xD1, A: 255},
		400: color.NRGBA{R: 0xA8, G: 0xA2, B: 0x9E, A: 255},
		500: color.NRGBA{R: 0x78, G: 0x71, B: 0x6C, A: 255},
		600: color.NRGBA{R: 0x57, G: 0x53, B: 0x4E, A: 255},
		700: color.NRGBA{R: 0x44, G: 0x40, B: 0x3C, A: 255},
		800: color.NRGBA{R: 0x29, G: 0x25, B: 0x24, A: 255},
		900: color.NRGBA{R: 0x1C, G: 0x19, B: 0x17, A: 255},
		950: color.NRGBA{R: 0x0C, G: 0x0A, B: 0x09, A: 255},
	},
	ColorRed: {
		50:  color.NRGBA{R: 0xFE, G: 0xF2, B: 0xF2, A: 255},
		100: color.NRGBA{R: 0xFE, G: 0xE2, B: 0xE2, A: 255},
		200: color.NRGBA{R: 0xFE, G: 0xCA, B: 0xCA, A: 255},
		300: color.NRGBA{R: 0xFC, G: 0xA5, B: 0xA5, A: 255},
		400: color.NRGBA{R: 0xF8, G: 0x71, B: 0x71, A: 255},
		500: color.NRGBA{R: 0xEF, G: 0x44, B: 0x44, A: 255},
		600: color.NRGBA{R: 0xDC, G: 0x26, B: 0x26, A: 255},
		700: color.NRGBA{R: 0xB9, G: 0x1C, B: 0x1C, A: 255},
		800: color.NRGBA{R: 0x99, G: 0x1B, B: 0x1B, A: 255},
		900: color.NRGBA{R: 0x7F, G: 0x1D, B: 0x1D, A: 255},
		950: color.NRGBA{R: 0x45, G: 0x0A, B: 0x0A, A: 255},
	},
	ColorOrange: {
		50:  color.NRGBA{R: 0xFF, G: 0xF7, B: 0xED, A: 255},
		100: color.NRGBA{R: 0xFF, G: 0xED, B: 0xD5, A: 255},
		200: color.NRGBA{R: 0xFE, G: 0xD7, B: 0xAA, A: 255},
		300: color.NRGBA{R: 0xFD, G: 0xBA, B: 0x74, A: 255},
		400: color.NRGBA{R: 0xFB, G: 0x92, B: 0x3C, A: 255},
		500: color.NRGBA{R: 0xF9, G: 0x73, B: 0x16, A: 255},
		600: color.NRGBA{R: 0xEA, G: 0x58, B: 0x0C, A: 255},
		700: color.NRGBA{R: 0xC2, G: 0x41, B: 0x0C, A: 255},
		800: color.NRGBA{R: 0x9A, G: 0x34, B: 0x12, A: 255},
		900: color.NRGBA{R: 0x7C, G: 0x2D, B: 0x12, A: 255},
		950: color.NRGBA{R: 0x43, G: 0x14, B: 0x07, A: 255},
	},
	ColorAmber: {
		50:  color.NRGBA{R: 0xFF, G: 0xFB, B: 0xEB, A: 255},
		100: color.NRGBA{R: 0xFE, G: 0xF3, B: 0xC7, A: 255},
		200: color.NRGBA{R: 0xFD, G: 0xE6, B: 0x8A, A: 255},
		300: color.NRGBA{R: 0xFC, G: 0xD3, B: 0x4D, A: 255},
		400: color.NRGBA{R: 0xFB, G: 0xBF, B: 0x24, A: 255},
		500: color.NRGBA{R: 0xF5, G: 0x9E, B: 0x0B, A: 255},
		600: color.NRGBA{R: 0xD9, G: 0x77, B: 0x06, A: 255},
		700: color.NRGBA{R: 0xB4, G: 0x53, B: 0x09, A: 255},
		800: color.NRGBA{R: 0x92, G: 0x40, B: 0x0E, A: 255},
		900: color.NRGBA{R: 0x78, G: 0x35, B: 0x0F, A: 255},
		950: color.NRGBA{R: 0x45, G: 0x1A, B: 0x03, A: 255},
	},
	ColorYellow: {
		50:  color.NRGBA{R: 0xFE, G: 0xFC, B: 0xE8, A: 255},
		100: color.NRGBA{R: 0xFE, G: 0xF9, B: 0xC3, A: 255},
		200: color.NRGBA{R: 0xFE, G: 0xF0, B: 0x8A, A: 255},
		300: color.NRGBA{R: 0xFD, G: 0xE0, B: 0x47, A: 255},
		400: color.NRGBA{R: 0xFA, G: 0xCC, B: 0x15, A: 255},
		500: color.NRGBA{R: 0xEA, G: 0xB3, B: 0x08, A: 255},
		600: color.NRGBA{R: 0xCA, G: 0x8A, B: 0x04, A: 255},
		700: color.NRGBA{R: 0xA1, G: 0x62, B: 0x07, A: 255},
		800: color.NRGBA{R: 0x85, G: 0x4D, B: 0x0E, A: 255},
		900: color.NRGBA{R: 0x71, G: 0x3F, B: 0x12, A: 255},
		950: color.NRGBA{R: 0x42, G: 0x20, B: 0x06, A: 255},
	},
	ColorLime: {
		50:  color.NRGBA{R: 0xF7, G: 0xFE, B: 0xE7, A: 255},
		100: color.NRGBA{R: 0xEC, G: 0xFC, B: 0xCB, A: 255},
		200: color.NRGBA{R: 0xD9, G: 0xF9, B: 0x9D, A: 255},
		300: color.NRGBA{R: 0xBE, G: 0xF2, B: 0x64, A: 255},
		400: color.NRGBA{R: 0xA3, G: 0xE6, B: 0x35, A: 255},
		500: color.NRGBA{R: 0x84, G: 0xCC, B: 0x16, A: 255},
		600: color.NRGBA{R: 0x65, G: 0xA3, B: 0x0D, A: 255},
		700: color.NRGBA{R: 0x4D, G: 0x7C, B: 0x0F, A: 255},
		800: color.NRGBA{R: 0x3F, G: 0x62, B: 0x12, A: 255},
		900: color.NRGBA{R: 0x36, G: 0x53, B: 0x14, A: 255},
		950: color.NRGBA{R: 0x1A, G: 0x2E, B: 0x05, A: 255},
	},
	ColorGreen: {
		50:  color.NRGBA{R: 0xF0, G: 0xFD, B: 0xF4, A: 255},
		100: color.NRGBA{R: 0xDC, G: 0xFC, B: 0xE7, A: 255},
		200: color.NRGBA{R: 0xBB, G: 0xF7, B: 0xD0, A: 255},
		300: color.NRGBA{R: 0x86, G: 0xEF, B: 0xAC, A: 255},
		400: color.NRGBA{R: 0x4A, G: 0xDE, B: 0x80, A: 255},
		500: color.NRGBA{R: 0x22, G: 0xC5, B: 0x5E, A: 255},
		600: color.NRGBA{R: 0x16, G: 0xA3, B: 0x4A, A: 255},
		700: color.NRGBA{R: 0x15, G: 0x80, B: 0x3D, A: 255},
		800: color.NRGBA{R: 0x16, G: 0x65, B: 0x34, A: 255},
		900: color.NRGBA{R: 0x14, G: 0x53, B: 0x2D, A: 255},
		950: color.NRGBA{R: 0x05, G: 0x2E, B: 0x16, A: 255},
	},
	ColorEmerald: {
		50:  color.NRGBA{R: 0xEC, G: 0xFD, B: 0xF5, A: 255},
		100: color.NRGBA{R: 0xD1, G: 0xFA, B: 0xE5, A: 255},
		200: color.NRGBA{R: 0xA7, G: 0xF3, B: 0xD0, A: 255},
		300: color.NRGBA{R: 0x6E, G: 0xE7, B: 0xB7, A: 255},
		400: color.NRGBA{R: 0x34, G: 0xD3, B: 0x99, A: 255},
		500: color.NRGBA{R: 0x10, G: 0xB9, B: 0x81, A: 255},
		600: color.NRGBA{R: 0x05, G: 0x96, B: 0x69, A: 255},
		700: color.NRGBA{R: 0x04, G: 0x78, B: 0x57, A: 255},
		800: color.NRGBA{R: 0x06, G: 0x5F, B: 0x46, A: 255},
		900: color.NRGBA{R: 0x06, G: 0x4E, B: 0x3B, A: 255},
		950: color.NRGBA{R: 0x02, G: 0x2C, B: 0x22, A: 255},
	},
	ColorTeal: {
		50:  color.NRGBA{R: 0xF0, G: 0xFD, B: 0xFA, A: 255},
		100: color.NRGBA{R: 0xCC, G: 0xFB, B: 0xF1, A: 255},
		200: color.NRGBA{R: 0x99, G: 0xF6, B: 0xE4, A: 255},
		300: color.NRGBA{R: 0x5E, G: 0xEA, B: 0xD4, A: 255},
		400: color.NRGBA{R: 0x2D, G: 0xD4, B: 0xBF, A: 255},
		500: color.NRGBA{R: 0x14, G: 0xB8, B: 0xA6, A: 255},
		600: color.NRGBA{R: 0x0D, G: 0x94, B: 0x88, A: 255},
		700: color.NRGBA{R: 0x0F, G: 0x76, B: 0x6E, A: 255},
		800: color.NRGBA{R: 0x11, G: 0x5E, B: 0x59, A: 255},
		900: color.NRGBA{R: 0x13, G: 0x4E, B: 0x4A, A: 255},
		950: color.NRGBA{R: 0x04, G: 0x2F, B: 0x2E, A: 255},
	},
	ColorCyan: {
		50:  color.NRGBA{R: 0xEC, G: 0xFE, B: 0xFF, A: 255},
		100: color.NRGBA{R: 0xCF, G: 0xFA, B: 0xFE, A: 255},
		200: color.NRGBA{R: 0xA5, G: 0xF3, B: 0xFC, A: 255},
		300: color.NRGBA{R: 0x67, G: 0xE8, B: 0xF9, A: 255},
		400: color.NRGBA{R: 0x22, G: 0xD3, B: 0xEE, A: 255},
		500: color.NRGBA{R: 0x06, G: 0xB6, B: 0xD4, A: 255},
		600: color.NRGBA{R: 0x08, G: 0x91, B: 0xB2, A: 255},
		700: color.NRGBA{R: 0x0E, G: 0x74, B: 0x90, A: 255},
		800: color.NRGBA{R: 0x15, G: 0x5E, B: 0x75, A: 255},
		900: color.NRGBA{R: 0x16, G: 0x4E, B: 0x63, A: 255},
		950: color.NRGBA{R: 0x08, G: 0x33, B: 0x44, A: 255},
	},
	ColorSky: {
		50:  color.NRGBA{R: 0xF0, G: 0xF9, B: 0xFF, A: 255},
		100: color.NRGBA{R: 0xE0, G: 0xF2, B: 0xFE, A: 255},
		200: color.NRGBA{R: 0xBA, G: 0xE6, B: 0xFD, A: 255},
		300: color.NRGBA{R: 0x7D, G: 0xD3, B: 0xFC, A: 255},
		400: color.NRGBA{R: 0x38, G: 0xBD, B: 0xF8, A: 255},
		500: color.NRGBA{R: 0x0E, G: 0xA5, B: 0xE9, A: 255},
		600: color.NRGBA{R: 0x02, G: 0x84, B: 0xC7, A: 255},
		700: color.NRGBA{R: 0x03, G: 0x69, B: 0xA1, A: 255},
		800: color.NRGBA{R: 0x07, G: 0x59, B: 0x85, A: 255},
		900: color.NRGBA{R: 0x0C, G: 0x4A, B: 0x6E, A: 255},
		950: color.NRGBA{R: 0x08, G: 0x2F, B: 0x49, A: 255},
	},
	ColorBlue: {
		50:  color.NRGBA{R: 0xEF, G: 0xF6, B: 0xFF, A: 255},
		100: color.NRGBA{R: 0xDB, G: 0xEA, B: 0xFE, A: 255},
		200: color.NRGBA{R: 0xBF, G: 0xDB, B: 0xFE, A: 255},
		300: color.NRGBA{R: 0x93, G: 0xC5, B: 0xFD, A: 255},
		400: color.NRGBA{R: 0x60, G: 0xA5, B: 0xFA, A: 255},
		500: color.NRGBA{R: 0x3B, G: 0x82, B: 0xF6, A: 255},
		600: color.NRGBA{R: 0x25, G: 0x63, B: 0xEB, A: 255},
		700: color.NRGBA{R: 0x1D, G: 0x4E, B: 0xD8, A: 255},
		800: color.NRGBA{R: 0x1E, G: 0x40, B: 0xAF, A: 255},
		900: color.NRGBA{R: 0x1E, G: 0x3A, B: 0x8A, A: 255},
		950: color.NRGBA{R: 0x17, G: 0x25, B: 0x54, A: 255},
	},
	ColorIndigo: {
		50:  color.NRGBA{R: 0xEE, G: 0xF2, B: 0xFF, A: 255},
		100: color.NRGBA{R: 0xE0, G: 0xE7, B: 0xFF, A: 255},
		200: color.NRGBA{R: 0xC7, G: 0xD2, B: 0xFE, A: 255},
		300: color.NRGBA{R: 0xA5, G: 0xB4, B: 0xFC, A: 255},
		400: color.NRGBA{R: 0x81, G: 0x8C, B: 0xF8, A: 255},
		500: color.NRGBA{R: 0x63, G: 0x66, B: 0xF1, A: 255},
		600: color.NRGBA{R: 0x4F, G: 0x46, B: 0xE5, A: 255},
		700: color.NRGBA{R: 0x43, G: 0x38, B: 0xCA, A: 255},
		800: color.NRGBA{R: 0x37, G: 0x30, B: 0xA3, A: 255},
		900: color.NRGBA{R: 0x31, G: 0x2E, B: 0x81, A: 255},
		950: color.NRGBA{R: 0x1E, G: 0x1B, B: 0x4B, A: 255},
	},
	ColorViolet: {
		50:  color.NRGBA{R: 0xF5, G: 0xF3, B: 0xFF, A: 255},
		100: color.NRGBA{R: 0xED, G: 0xE9, B: 0xFE, A: 255},
		200: color.NRGBA{R: 0xDD, G: 0xD6, B: 0xFE, A: 255},
		300: color.NRGBA{R: 0xC4, G: 0xB5, B: 0xFD, A: 255},
		400: color.NRGBA{R: 0xA7, G: 0x8B, B: 0xFA, A: 255},
		500: color.NRGBA{R: 0x8B, G: 0x5C, B: 0xF6, A: 255},
		600: color.NRGBA{R: 0x7C, G: 0x3A, B: 0xED, A: 255},
		700: color.NRGBA{R: 0x6D, G: 0x28, B: 0xD9, A: 255},
		800: color.NRGBA{R: 0x5B, G: 0x21, B: 0xB6, A: 255},
		900: color.NRGBA{R: 0x4C, G: 0x1D, B: 0x95, A: 255},
		950: color.NRGBA{R: 0x2E, G: 0x10, B: 0x65, A: 255},
	},
	ColorPurple: {
		50:  color.NRGBA{R: 0xFA, G: 0xF5, B: 0xFF, A: 255},
		100: color.NRGBA{R: 0xF3, G: 0xE8, B: 0xFF, A: 255},
		200: color.NRGBA{R: 0xE9, G: 0xD5, B: 0xFF, A: 255},
		300: color.NRGBA{R: 0xD8, G: 0xB4, B: 0xFE, A: 255},
		400: color.NRGBA{R: 0xC0, G: 0x84, B: 0xFC, A: 255},
		500: color.NRGBA{R: 0xA8, G: 0x55, B: 0xF7, A: 255},
		600: color.NRGBA{R: 0x93, G: 0x33, B: 0xEA, A: 255},
		700: color.NRGBA{R: 0x7E, G: 0x22, B: 0xCE, A: 255},
		800: color.NRGBA{R: 0x6B, G: 0x21, B: 0xA8, A: 255},
		900: color.NRGBA{R: 0x58, G: 0x1C, B: 0x87, A: 255},
		950: color.NRGBA{R: 0x3B, G: 0x07, B: 0x64, A: 255},
	},
	ColorFuchsia: {
		50:  color.NRGBA{R: 0xFD, G: 0xF4, B: 0xFF, A: 255},
		100: color.NRGBA{R: 0xFA, G: 0xE8, B: 0xFF, A: 255},
		200: color.NRGBA{R: 0xF5, G: 0xD0, B: 0xFE, A: 255},
		300: color.NRGBA{R: 0xF0, G: 0xAB, B: 0xFC, A: 255},
		400: color.NRGBA{R: 0xE8, G: 0x79, B: 0xF9, A: 255},
		500: color.NRGBA{R: 0xD9, G: 0x46, B: 0xEF, A: 255},
		600: color.NRGBA{R: 0xC0, G: 0x26, B: 0xD3, A: 255},
		700: color.NRGBA{R: 0xA2, G: 0x1C, B: 0xAF, A: 255},
		800: color.NRGBA{R: 0x86, G: 0x19, B: 0x8F, A: 255},
		900: color.NRGBA{R: 0x70, G: 0x1A, B: 0x75, A: 255},
		950: color.NRGBA{R: 0x4A, G: 0x04, B: 0x4E, A: 255},
	},
	ColorPink: {
		50:  color.NRGBA{R: 0xFD, G: 0xF2, B: 0xF8, A: 255},
		100: color.NRGBA{R: 0xFC, G: 0xE7, B: 0xF3, A: 255},
		200: color.NRGBA{R: 0xFB, G: 0xCF, B: 0xE8, A: 255},
		300: color.NRGBA{R: 0xF9, G: 0xA8, B: 0xD4, A: 255},
		400: color.NRGBA{R: 0xF4, G: 0x72, B: 0xB6, A: 255},
		500: color.NRGBA{R: 0xEC, G: 0x48, B: 0x99, A: 255},
		600: color.NRGBA{R: 0xDB, G: 0x27, B: 0x77, A: 255},
		700: color.NRGBA{R: 0xBE, G: 0x18, B: 0x5D, A: 255},
		800: color.NRGBA{R: 0x9D, G: 0x17, B: 0x4D, A: 255},
		900: color.NRGBA{R: 0x83, G: 0x18, B: 0x43, A: 255},
		950: color.NRGBA{R: 0x50, G: 0x07, B: 0x24, A: 255},
	},
	ColorRose: {
		50:  color.NRGBA{R: 0xFF, G: 0xF1, B: 0xF2, A: 255},
		100: color.NRGBA{R: 0xFF, G: 0xE4, B: 0xE6, A: 255},
		200: color.NRGBA{R: 0xFE, G: 0xCD, B: 0xD3, A: 255},
		300: color.NRGBA{R: 0xFD, G: 0xA4, B: 0xAF, A: 255},
		400: color.NRGBA{R: 0xFB, G: 0x71, B: 0x85, A: 255},
		500: color.NRGBA{R: 0xF4, G: 0x3F, B: 0x5E, A: 255},
		600: color.NRGBA{R: 0xE1, G: 0x1D, B: 0x48, A: 255},
		700: color.NRGBA{R: 0xBE, G: 0x12, B: 0x3C, A: 255},
		800: color.NRGBA{R: 0x9F, G: 0x12, B: 0x39, A: 255},
		900: color.NRGBA{R: 0x88, G: 0x13, B: 0x37, A: 255},
		950: color.NRGBA{R: 0x4C, G: 0x05, B: 0x19, A: 255},
	},
}
