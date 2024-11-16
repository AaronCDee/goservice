package goservice

import (
	"fmt"
	"testing"
)

type exampleAction struct { Action }

func (exampleAction) expects() []string {
	return []string{"hello"}
}

func (exampleAction) promises() []string {
	return []string{"goodbye"}
}

func (exampleAction) executed(ctx *Context) {
	fmt.Println("--------Executing the RIGHT action-------")
	hello := ctx.Values["hello"].(string)
	
	ctx.Values["goodbye"] = "goodbye"
	
	fmt.Println(hello)
}

func NewExampleAction() *exampleAction {
	action := &exampleAction{}

	action.Actionable = action

	return action
}

func TestAction_Execute(t *testing.T) {
	fmt.Println("===Testing TextAction_Execute===")

	action  := NewExampleAction()
	outcome := action.Execute(map[string]any{"hello": "hello"})

	values 			:= outcome.Values
	expectedGoodbye := "goodbye"
	goodbye 		:= values["goodbye"].(string)
	
	if (outcome.Failure) {
		t.Errorf("Expected outcome to be a success but was a failure")
	}
	
	if (expectedGoodbye != goodbye) {
		t.Errorf("Expected %s in context.Values, got %s", expectedGoodbye, goodbye)
	}
}
