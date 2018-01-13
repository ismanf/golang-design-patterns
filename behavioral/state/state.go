package state

import (
	"fmt"
)

type State interface {
	executeState(c *Context)
}

type Context struct {
	StepIndex int
	StepName  string
	Current   State
}

type StartState struct{}

func (s *StartState) executeState(c *Context) {
	c.StepIndex = 1
	c.StepName = "start"
	c.Current = &StartState{}
}

type InprogressState struct{}

func (s *InprogressState) executeState(c *Context) {
	c.StepIndex = 2
	c.StepName = "inprogress"
	c.Current = &InprogressState{}
}

type StopState struct{}

func (s *StopState) executeState(c *Context) {
	c.StepIndex = 3
	c.StepName = "stop"
	c.Current = &StopState{}
}

func StateDemo() {
	context := &Context{}

	startState := &StartState{}
	startState.executeState(context)
	fmt.Println("state: ", context)

	inprState := &InprogressState{}
	inprState.executeState(context)
	fmt.Println("state: ", context)

	stopState := &StopState{}
	stopState.executeState(context)
	fmt.Println("state: ", context)
}
