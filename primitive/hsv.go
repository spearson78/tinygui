package primitive

func max(a, b uint8) uint8 {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b uint8) uint8 {
	if a < b {
		return a
	} else {
		return b
	}
}

//go:noinline
func ToHSV(r, g, b uint8) (h, s, v uint8) {

	max := int(max(max(r, g), b))
	min := int(min(min(r, g), b))

	v = uint8(max)
	if v == 0 {
		h = 0
		s = 0
		return
	}

	diff := max - min

	s = uint8((255 * diff) / max)

	if s == 0 {
		h = 0
		return
	}

	if max == int(r) {
		h = uint8(0 + (43*(int(g)-int(b)))/diff)
	} else if max == int(g) {
		h = uint8(85 + (43*(int(b)-int(r)))/diff)
	} else if max == int(b) {
		h = uint8(171 + (43*(int(r)-int(g)))/diff)
	}

	return
}

//go:noinline
func ToRGB(h, s, v uint8) (r, g, b uint8) {

	if s == 0 {
		return v, v, v
	}

	region := int(h) / 43

	remainder := (int(h) - (region * 43)) * 6

	p := uint8((int(v) * (255 - int(s))) >> 8)
	q := uint8((int(v) * (255 - ((int(s) * remainder) >> 8))) >> 8)
	t := uint8((int(v) * (255 - ((int(s) * (255 - remainder)) >> 8))) >> 8)

	switch region {
	case 0:
		r, g, b = v, t, p
	case 1:
		r, g, b = q, v, p
	case 2:
		r, g, b = p, v, t
	case 3:
		r, g, b = p, q, v
	case 4:
		r, g, b = t, p, v
	default:
		r, g, b = v, p, q
	}

	return
}
