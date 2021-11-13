package primitive

import (
	"image/color"
)

//go:noinline
func TransitionColor(fromColor color.RGBA, toColor color.RGBA, amount uint8) color.RGBA {

	fh, fs, fv := ToHSV(fromColor.R, fromColor.G, fromColor.B)
	th, ts, tv := ToHSV(toColor.R, toColor.G, toColor.B)

	if ts <= 25 {
		th = fh
	}

	dh := int(fh) - int(th)
	ds := int(fs) - int(ts)
	dv := int(fv) - int(tv)

	ch := uint8(int(fh) - (dh*(255-int(amount)))/255)
	cs := uint8(int(fs) - (ds*(255-int(amount)))/255)
	cv := uint8(int(fv) - (dv*(255-int(amount)))/255)

	sr, sg, sb := ToRGB(ch, cs, cv)

	return color.RGBA{sr, sg, sb, 255}
}

//go:noinline
func Desaturate(c color.RGBA, amount uint8) color.RGBA {
	ch, cs, cv := ToHSV(c.R, c.G, c.B)

	cs = uint8((int(cs) * (255 - int(amount))) / 255)
	icv := int(cv)
	if cs < 25 {
		icv = 127
	}
	icv = icv + int(amount)

	if icv > 250 {
		cv = 250
	} else {
		cv = uint8(icv)
	}

	sr, sg, sb := ToRGB(ch, cs, cv)

	return color.RGBA{sr, sg, sb, 255}
}
