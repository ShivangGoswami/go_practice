package display

import "../../ch7/eval"

func Example_expr() {
	e, _ := eval.Parse("sqrt(A / pi)")
	Display("e", e)
}
