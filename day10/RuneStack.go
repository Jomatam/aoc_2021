package day10

type RuneStack []rune

func NewRuneStack() RuneStack {
	return make(RuneStack, 0)
}

func (stack RuneStack) Push(r rune) RuneStack {
	return append(stack, r)
}

func (stack RuneStack) Pop() (RuneStack, rune) {
	length := len(stack)
	return stack[:length-1], stack[length-1]
}
