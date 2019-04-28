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
	var state State

	state = &StartState{}
	state.executeState(context)
	fmt.Println("state: ", context)

	state = &InprogressState{}
	state.executeState(context)
	fmt.Println("state: ", context)

	state = &StopState{}
	state.executeState(context)
	fmt.Println("state: ", context)
}
