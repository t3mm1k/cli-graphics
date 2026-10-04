package engine

import "slices"

func HitTest(comp Component, x, y int) Component {
	if comp == nil {
		return nil
	}

	w, h := comp.GetSize()
	if x < 0 || y < 0 || x >= w || y >= h {
		return nil
	}

	if container, ok := comp.(Container); ok {
		children := container.Children()

		for _, child := range slices.Backward(children) {

			if child == nil {
				continue
			}

			cx, cy := child.GetCoords()
			cw, ch := child.GetSize()

			if x >= cx && x < cx+cw && y >= cy && y < cy+ch {
				if hit := HitTest(child, x-cx, y-cy); hit != nil {
					return hit
				}
			}
		}
	}

	return comp
}
