package graphicgo

import (
	"errors"
)

const (
	Slim = iota
	Middle
	Bold
)

func abs(x int64) (abs int64) {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func DrawDot(x int64, y int64, color [4]byte, width int) (err error) {
	if width == Slim {
		dot(x, y, color)
	} else if width == Middle {
		dot(x-1, y, color)
		dot(x, y, color)
		dot(x+1, y, color)
		dot(x, y-1, color)
		dot(x, y+1, color)
	} else if width == Bold {
		dot(x, y-2, color)
		dot(x-1, y-1, color)
		dot(x, y-1, color)
		dot(x+1, y-1, color)
		dot(x-2, y, color)
		dot(x-1, y, color)
		dot(x, y, color)
		dot(x+1, y, color)
		dot(x+2, y, color)
		dot(x-1, y+1, color)
		dot(x, y+1, color)
		dot(x+1, y+1, color)
		dot(x+2, y, color)
	} else {
		return errors.New("width type not found")
	}
	return nil
}

func DrawLine(x1 int64, y1 int64, x2 int64, y2 int64, color [4]byte, width int) {
	var dx int64 = abs(x2 - x1)
	var dy int64 = abs(y2 - y1)
	greater_than_45 := false

	if dx < dy {
		greater_than_45 = true
		x1, y1 = y1, x1
		x2, y2 = y2, x2
		dx, dy = dy, dx
	}

	var ix, iy int64 = 1, 1
	if x2-x1 < 0 {
		ix = -1
	}
	if y2-y1 < 0 {
		iy = -1
	}
	cx := x1
	cy := y1
	n2dy := dy * 2
	n2DyDx := (dy - dx) * 2
	d := dy*2 - dx

	for cx != x2 {
		if d < 0 {
			d += n2dy
		} else {
			cy += iy
			d += n2DyDx
		}
		if greater_than_45 {
			DrawDot(cy, cx, color, width)
		} else {
			DrawDot(cx, cy, color, width)
		}
		cx += ix
	}
}

func drawCircle8(xc int64, yc int64, x int64, y int64, color [4]byte, width int) {
	DrawDot(xc+x, yc-y, color, width)
	DrawDot(xc-x, yc+y, color, width)
	DrawDot(xc-x, yc-y, color, width)
	DrawDot(xc+x, yc+y, color, width)
	DrawDot(xc+y, yc+x, color, width)
	DrawDot(xc-y, yc+x, color, width)
	DrawDot(xc+y, yc-x, color, width)
	DrawDot(xc-y, yc-x, color, width)
}

func DrawCircle(xc int64, yc int64, r int64, color [4]byte, width int, fill bool) {
	if xc+r < 0 ||
		xc-r >= screenWidth ||
		yc+r < 0 ||
		yc-r >= screenHeight {
		return
	}

	var x int64 = 0
	var y int64 = r
	var d int64 = 3 - 2*r
	var yi int64

	for x <= y {
		if fill {
			for yi = y - 10; yi <= y; yi++ {
				drawCircle8(xc, yc, x, yi, color, width)
			}
		} else {
			drawCircle8(xc, yc, x, y, color, width)
		}
		if d < 0 {
			d = d + 4*x + 6
		} else {
			d = d + 4*(x-y) + 10
			y--
		}
		x++
	}
}
