package theme

import (
	"image/color"

	"github.com/spearson78/tinyamifont"
	"github.com/spearson78/tinyamifont/fonts/newwebfont/voyager"
)

type Theme struct {
	Text           color.RGBA
	Background     color.RGBA
	Border         color.RGBA
	DefaultColor   color.RGBA
	SecondaryColor color.RGBA
	SuccessColor   color.RGBA
	ErrorColor     color.RGBA
	DisabedColor   color.RGBA
	Shadow1        color.RGBA
	Shadow2        color.RGBA
	Font           *tinyamifont.Font
	SecondaryFont  *tinyamifont.Font
}

func DefaultTheme() Theme {

	font := tinyamifont.MustLoadFont(voyager.Regular18pt)
	secondaryFont := tinyamifont.MustLoadFont(voyager.Regular9pt)

	return Theme{
		Text:           color.RGBA{0, 0, 0, 255},
		Background:     color.RGBA{250, 250, 250, 255},
		Border:         color.RGBA{117, 117, 117, 255},
		Shadow1:        color.RGBA{177, 177, 177, 255},
		Shadow2:        color.RGBA{205, 205, 205, 255},
		DefaultColor:   color.RGBA{25, 118, 210, 255},
		SecondaryColor: color.RGBA{156, 39, 176, 255},
		SuccessColor:   color.RGBA{46, 125, 50, 255},
		ErrorColor:     color.RGBA{212, 55, 80, 255},
		DisabedColor:   color.RGBA{224, 224, 224, 255},
		Font:           &font,
		SecondaryFont:  &secondaryFont,
	}
}
