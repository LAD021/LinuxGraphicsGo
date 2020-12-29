package drawork

type POINT3D struct {
	x, y, z float64
}

func CreatePoint3D(x float64, y float64, z float64) *POINT3D {
	return &POINT3D{
		x: x,
		y: y,
		z: z,
	}
}

func CopyPoint3D(_p *POINT3D) *POINT3D {
	return &POINT3D{
		x: _p.x,
		y: _p.y,
		z: _p.z,
	}
}
