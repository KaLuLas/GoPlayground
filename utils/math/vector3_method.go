package math

func NewVector3(x float64, y float64, z float64) *Vector3 {
	return &Vector3{x, y, z}
}

func (receiver Vector3) Abs() float64 {
	return receiver.x*receiver.x + receiver.y*receiver.y + receiver.z*receiver.z
}
