package expanding_test

import (
	"fmt"

	"github.com/schwarmco/go-expanding"
)

func ExampleInteger_Expand() {

	a := expanding.Integer{
		Min:  0,
		Max:  10,
		Step: 2,
	}

	fmt.Println(a.Expand())

	// Output:
	// [0 2 4 6 8 10]
}

func ExampleFloat_Expand() {

	a := expanding.Float{
		Min:  0,
		Max:  1,
		Step: 0.2,
	}

	num := len(a.Expand())

	fmt.Println(num)

	// Output:
	// 6
}
