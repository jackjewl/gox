package apply

import (
	"gox/stack"
)

// decimal  convert to other base num,example ,decimal convert to binary,hex
func BaseConvert(s stack.Stack[rune], num int, base int) {
	digit := []rune{'0', '1', '2', '3', '4', '5', '6',
		'7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}
	for num > 0 {
		s.Push(digit[num%base])
		num /= base
	}
}

func ParenthesisMatch(expression []rune) bool {
	var s stack.Stack[rune] = stack.NewLinkedStack[rune]()
	for i := 0; i < len(expression); i++ {
		switch expression[i] {
		case '(':
			s.Push(expression[i])
		case '[':
			s.Push(expression[i])
		case '{':
			s.Push(expression[i])
		case ')':
			{
				if s.Empty() || s.Pop() != '(' {
					return false
				}
			}
		case ']':
			{
				if s.Empty() || s.Pop() != '[' {
					return false
				}
			}
		case '}':
			{
				if s.Empty() || s.Pop() != '{' {
					return false
				}
			}
		}
	}
	return s.Empty()
}

/*
//表达式求值
func Evaluate(expression []rune) float64 {

	//计算的同时生成逆波兰式
	rpn := vector.NewVector[rune]()
	numStack := stack.LinkedStack[rune]{}
	optrStack := stack.LinkedStack[rune]{}

	optrStack.Push(rune('#'))
	for !optrStack.Empty() {
		//如果是操作数，如入操作数栈，追加到rpn

		//如果是操作符，看操作符栈的元素和当前输入操作符的优先级
		//当前优先级大，操作符进栈
		//优先级相等，脱括号，取下一个自负
		//当前操作符优先级小,当前操作符追加到rpn，取出操作数，将操作结果压入操作数栈
	}

return numStack.top()

}
*/

/*
func RPNEvaluate(expression []rune)float64{
	for expression尚未扫描完毕{
		扫描下一个元素
		如果是操作数，入操作数栈
		如果是操作符，在操作数栈中取出需要的操作数，将运算结果入栈
	}
	返回操作数栈顶元素，即为结果
}

*/

type QueenCoordinate struct {
	X, Y int
}

func Conflict(a, b QueenCoordinate) bool {
	if a.X == b.X || a.Y == b.Y || a.X-a.Y == b.X-b.Y || a.X+a.Y == b.X+b.Y {
		return true
	}
	return false
}

func PlaceQueen(queenNum int) {
	solution := stack.SequenceStack[QueenCoordinate]{}
	q := QueenCoordinate{
		X: 0,
		Y: 0,
	}

	for {

		if solution.Size() >= queenNum || q.Y >= queenNum {
			q = solution.Pop()
			q.Y++
		} else {
			for q.Y < queenNum && solution.Elems.Find(q, Conflict) < 0 {
				q.Y++
			}
			if q.Y < queenNum {
				solution.Push(q)
				if solution.Size() >= queenNum {
					//解的数量+1
				}
				q.X++
				q.Y = 0
			}
		}
	}

}
