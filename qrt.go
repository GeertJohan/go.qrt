package qrt

import (
	"code.google.com/p/rsc/qr"
	"github.com/foize/go.sgr"
)

// blocks defined in a slice with 16 positions (4 bits)
// _____
// |1|2|
// |4|8|
// ⎺⎺⎺⎺⎺
//
var blocks = []rune{
	' ', '▘', '▝', '▀',
	'▖', '▌', '▞', '▛',
	'▗', '▚', '▐', '▜',
	'▄', '▙', '▟', '█',
}

// A Level denotes a QR error correction level.
// From least to most tolerant of errors, they are L, M, Q, H.
type Level int

const (
	L Level = iota // 20% redundant
	M              // 38% redundant
	Q              // 55% redundant
	H              // 65% redundant
)

// Config hold all paramters that can be set with this library
type Config struct {
	// UseSGR, when true, Generate() uses Select Graphic Rendition (github.com/foize/go.sgr) to force correct colors in the terminal (black/white)
	// You should disable this when writing the result to a file
	UseSGR bool

	// Level indicates the required redudancy level for the generated QR's
	// You can keep the redundancy value low because screens are pretty bright and clean.... Or well, they should be.
	Level Level
}

// DefaultConfig is used by the Generate function
// You can change it's values, or create a Config instance of your own.
var DefaultConfig = &Config{
	UseSGR: true,
	Level:  M,
}

// Generate a text string to a QR code, which you can write to a terminal or file.
// Generate is a shorthand for DefaultConfig.Generate(text)
func Generate(text string) (string, error) {
	return DefaultConfig.Generate(text)
}

// Generate a text string to a QR code, which you can write to a terminal or file.
func (c *Config) Generate(text string) (string, error) {
	code, err := qr.Encode(text, qr.Level(c.Level))
	if err != nil {
		return "", err
	}

	// calculate size in blocks
	// two bits per block
	// add one block for remaining singlebit (if existing)
	// add two blocks for border left and right
	size := code.Size/2 + (code.Size % 2) + 2

	// rune slice
	//++ TODO: precalculate size
	qrRunes := make([]rune, 0)

	// upper border
	c.addWhiteRow(&qrRunes, size)

	// content
	for y := 0; y < code.Size; y += 2 {
		if c.UseSGR {
			qrRunes = append(qrRunes, []rune(sgr.FgWhite+sgr.BgBlack)...)
		}
		qrRunes = append(qrRunes, '█')
		for x := 0; x < code.Size; x += 2 {
			var num int8
			if !code.Black(x, y) {
				num += 1
			}
			if !code.Black(x+1, y) {
				num += 2
			}
			if !code.Black(x, y+1) {
				num += 4
			}
			if !code.Black(x+1, y+1) {
				num += 8
			}
			qrRunes = append(qrRunes, blocks[num])
		}
		qrRunes = append(qrRunes, '█')
		if c.UseSGR {
			qrRunes = append(qrRunes, []rune(sgr.Reset)...)
		}
		qrRunes = append(qrRunes, '\n')
	}

	// add lower border when required (only required when QR size is odd)
	if code.Size%2 == 0 {
		c.addWhiteRow(&qrRunes, size)
	}

	return string(qrRunes), nil
}

func (c *Config) addWhiteRow(qrRunes *[]rune, width int) {
	if c.UseSGR {
		*qrRunes = append(*qrRunes, []rune(sgr.FgWhite+sgr.BgBlack)...)
	}
	for i := 0; i < width; i++ {
		*qrRunes = append(*qrRunes, '█')
	}
	if c.UseSGR {
		*qrRunes = append(*qrRunes, []rune(sgr.Reset)...)
	}
	*qrRunes = append(*qrRunes, '\n')
}
