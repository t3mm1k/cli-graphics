package utils

func GetActuallySize(hasBorder bool, w, h int) (W, H int) {
	if hasBorder {
		return w + 2, h + 2
	}
	return w, h
}
