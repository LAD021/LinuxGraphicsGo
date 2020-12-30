package drawork

import (
	"LinuxGraphicsGo/graphicgo"
	"math"
)

func InitCube() {
	End = 0
	for num := 0; num < Cubenum; num++ {
		MyCube[num].SideLength = Square_root_two * (float64(num) + 1) * sidegap

		MyCube[num].PT[0] = POINT3D{0, 0, 0}
		MyCube[num].PT[1] = POINT3D{
			MyCube[num].SideLength / Square_root_two,
			MyCube[num].SideLength / Square_root_two,
			0,
		}
		MyCube[num].PT[2] = POINT3D{
			0,
			MyCube[num].SideLength / Square_root_two,
			0,
		}
		MyCube[num].PT[3] = POINT3D{
			MyCube[num].SideLength / Square_root_two * (-1),
			MyCube[num].SideLength / Square_root_two,
			0,
		}
		for num1 := 4; num1 < 9; num1++ {
			MyCube[num].PP[(num1 - 4)] = MyCube[num].PT[(num1 - 4)]
			MyCube[num].PT[num1] = POINT3D{
				MyCube[num].PP[(num1 - 4)].x,
				MyCube[num].PP[(num1 - 4)].y,
				MyCube[num].PP[(num1-4)].z + MyCube[num].SideLength,
			}
			MyCube[num].PP[num1] = MyCube[num].PT[num1]
		}
		MyCube[num].Taking = 1
		MyCube[num].ShopTime = num * TurnRate
		MyCube[num].HadDone = 0
	}
	Changing()
}

func Rotation() {
	for num := 0; num < Cubenum; num++ {
		if MyCube[num].Taking == 1 || MyCube[num].ShopTime == 0 {
			for num1 := 0; num1 < 4; num1++ {
				MyCube[num].PP[num1].x =
					(MyCube[num].PT[num1].x-0)*math.Cos(-float64(MyCube[num].HadDone)*PI/FPS) -
						MyCube[num].PT[num-1].y - MyCube[num].SideLength/Square_root_two*math.Sin(-float64(MyCube[num].HadDone)*PI/FPS+0)
				MyCube[num].PP[num1].y =
					(MyCube[num].PT[num1].x-0)*math.Sin(-float64(MyCube[num].HadDone)*PI/FPS) +
						MyCube[num].PT[num-1].y - MyCube[num].SideLength/Square_root_two*math.Cos(-float64(MyCube[num].HadDone)*PI/FPS+0) +
						MyCube[num].SideLength/Square_root_two
			}

			for num2 := 4; num2 < 8; num2++ {
				MyCube[num].PP[num2].x = MyCube[num].PP[num2-4].x
				MyCube[num].PP[num2].y = MyCube[num].PP[num2-4].y
				MyCube[num].PP[num2].z = MyCube[num].PP[num2-4].z + MyCube[num].SideLength
			}

			MyCube[num].HadDone++
			if MyCube[num].HadDone == FPS+1 {
				MyCube[num].Taking = 0
			} else {
				MyCube[num].ShopTime--
			}
		}
	}
}

func Changing() {
	for num := 0; num < Cubenum; num++ {
		for num1 := 0; num1 < 8; num1++ {
			MyCube[num].Pp[num1] = Projection(&MyCube[num].PP[num1])
		}
	}
}

func DrawCube() {
	var temp int
	for num := 0; num < Cubenum; num++ {
		for num1 := 0; num1 < 4; num1++ {
			if (num1 + 1) == 4 {
				temp = 0
			} else {
				temp = num1 + 1
			}
			graphicgo.DrawLine(
				int64(MyCube[num].Pp[num1].x),
				int64(MyCube[num].Pp[num1].y),
				int64(MyCube[num].Pp[temp].x),
				int64(MyCube[num].Pp[temp].y),
				LineColor,
				graphicgo.Middle,
			)

			temp = num1 + 4
			graphicgo.DrawLine(
				int64(MyCube[num].Pp[num1].x),
				int64(MyCube[num].Pp[num1].y),
				int64(MyCube[num].Pp[temp].x),
				int64(MyCube[num].Pp[temp].y),
				LineColor,
				graphicgo.Middle,
			)

			if (num1 + 5) == 8 {
				temp = 4
			} else {
				temp = num1 + 5
			}
			graphicgo.DrawLine(
				int64(MyCube[num].Pp[num1+4].x),
				int64(MyCube[num].Pp[num1+4].y),
				int64(MyCube[num].Pp[temp].x),
				int64(MyCube[num].Pp[temp].y),
				LineColor,
				graphicgo.Middle,
			)
		}
	}
}

func IfEnd() {
	if MyCube[Cubenum-1].Taking == 0 {
		End = 1
	}
}

func Projection(p3 *POINT3D) POINT3D {
	var p2 POINT3D
	p2.x = p3.x + Center_x
	p2.y = (p3.y*Square_root_two - p3.z*2) / (Square_root_six + Center_y)
	p2.z = 0
	return p2
}
