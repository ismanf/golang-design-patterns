package interpreter

import (
	"strings"
)

type Expression interface {
	interpret(string) bool
}

type MainInterpreter struct {
	state string
}
func (i *MainInterpreter) interpret(context string)  bool{
	if strings.Contains(i.state, context) {
		return true
	}
	return false
}

type OrExpressionInterpreter struct {
	expOne, expTwo Expression
}
func (o *OrExpressionInterpreter) interpret(context string) bool {
	return o.expOne.interpret(context) || o.expTwo.interpret(context)
}

type AndExpressionInterpreter struct {
	expOne, expTwo Expression
}
func (o *AndExpressionInterpreter) interpret(context string) bool {
	return o.expOne.interpret(context) && o.expTwo.interpret(context)
}
