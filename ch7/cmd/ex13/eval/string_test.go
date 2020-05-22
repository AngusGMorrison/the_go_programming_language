package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		input, wantStr string
		env            Env
		wantEval       string
	}{
		{"sqrt(A / pi)", "sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", "(pow(x, 3) + pow(y, 3))", Env{"x": 12, "y": 1}, "1729"},
		{"5 / 9 * (F - 32)", "((5 / 9) * (F - 32))", Env{"F": -40}, "-40"},
	}
	for _, test := range tests {
		// Test conversion to string
		fmt.Println(test.input)
		expr, err := Parse(test.input)
		if err != nil {
			t.Error(err)
			continue
		}
		gotStr := expr.String()
		fmt.Printf("\t%s\n", gotStr)
		if gotStr != test.wantStr {
			t.Errorf("%s.String() = %q, want %q\n", test.input, gotStr, test.wantStr)
		}

		// Test parsing of converted string
		expr, err = Parse(gotStr)
		if err != nil {
			t.Error(err)
			continue
		}
		gotEval := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s \n", test.env, gotEval)
		if gotEval != test.wantEval {
			t.Errorf("%s.Eval() in %v = %q, want %q\n", gotStr, test.env, gotEval, test.wantEval)
		}
	}
}
