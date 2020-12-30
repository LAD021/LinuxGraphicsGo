package drawork

import "sync"

type Cube struct {
	SideLength float64
	PP         [8]POINT3D
	PT         [8]POINT3D
	Pp         [8]POINT3D
	Taking     int
	ShopTime   int
	HadDone    int
}

var (
	cubeOnce sync.Once
	instance *[Cubenum]Cube = nil
)

func GetCube() *[Cubenum]Cube {
	cubeOnce.Do(func() {
		instance = &[Cubenum]Cube{}
	})
	return instance
}

var MyCube = GetCube()

var End int = 0

// var Temp int
