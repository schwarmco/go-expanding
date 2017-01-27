package expanding

type Integer struct {
	Min, Max, Step int
}

type Float struct {
	Min, Max, Step float64
}

// expands an Integer Map to Integer Slice
func (i Integer) Expand() []int {
	var j []int
	for k := i.Min; k <= i.Max; k += i.Step {
		j = append(j, k)
	}
	return j
}

// expands a Float Map to Float Slice
func (i Float) Expand() []float64 {
	var j []float64
	for k := i.Min; k <= i.Max; k += i.Step {
		j = append(j, k)
	}
	return j
}
