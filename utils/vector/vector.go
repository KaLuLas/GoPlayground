package vector

type Element interface{}

type Vector struct {
	data []Element
}

func NewVector(length int) *Vector {
	var vec = new(Vector)
	vec.data = make([]Element, length)
	return vec
}

func (vec *Vector) At(index int) Element {
	if index < 0 || index >= len(vec.data) {
		return nil
	}

	return vec.data[index]
}

func (vec *Vector) Set(index int, e Element) {
	if index < 0 || index >= len(vec.data) {
		return
	}

	vec.data[index] = e
}
