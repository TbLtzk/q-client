package vm

import "gitlab.com/q-dev/q-client/ethstate"

type Debugger interface {
	BreakHook(step int, op OpCode, mem *Memory, stack *Stack, object *ethstate.StateObject) bool
	StepHook(step int, op OpCode, mem *Memory, stack *Stack, object *ethstate.StateObject) bool
	BreakPoints() []int64
	SetCode(byteCode []byte)
}
