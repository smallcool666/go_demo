package main

import (
	"fmt"
	"math"
	"testing"
)

/*
	构建一个简单算术表达式的求值器。
	表达式语言由浮点数符号（小数点）；二元操作符+，-，*， 和/；一元操作符-x和+x；调用pow(x,y)，sin(x)，和sqrt(x)的函数；例如x和pi的变量；当然也有括号和标准的优先级运算符。所有
的值都是float64类型。
	这个包只会对外公开Expr，Env，和Var类型。调用方不需要获取其它的表达式类型就可以使用这个求值器。
*/

//使用一个接口Expr来表示Go语言中任意的表达式。
type Expr interface{
	//这个方法会根据给定的environment变量返回表达式的值
	Eval(env Env) float64
}
//Var类型表示对一个变量的引用
type Var string
type literal float64
type unary struct {
	op rune // one of '+', '-'
	x Expr
}
type binary struct {
	op rune // one of '+', '-', '*', '/'
	x, y Expr
}
// A call represents a function call expression, e.g., sin(x).
type call struct {
	fn string // one of "pow", "sin", "sqrt"
	args []Expr
}
//为了计算一个包含变量的表达式，我们需要一个environment变量将变量的名字映射成对应的值
type Env map[Var]float64
//Var类型的这个方法对一个environment变量进行查找，如果这个变量没有在environment中定义过这个方法会返回一个零值
func (v Var) Eval(env Env) float64 {
	return env[v]
}
//literal类型的这个方法简单的返回它真实的值
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}
func (b binary) Eval(env Env) float64 {
	switch b.op {
		case '+':
			return b.x.Eval(env) + b.y.Eval(env)
		case '-':
			return b.x.Eval(env) - b.y.Eval(env)
		case '*':
			return b.x.Eval(env) * b.y.Eval(env)
		case '/':
			return b.x.Eval(env) / b.y.Eval(env)
		}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
		case "pow":
			return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
		case "sin":
			return math.Sin(c.args[0].Eval(env))
		case "sqrt":
			return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}
func main() {
	
}
