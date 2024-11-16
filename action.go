package goservice

type Actionable interface {
	Execute(values map[string]any) *Context

	executed(ctx *Context)

	rolledBack(ctx *Context)

	validateExpectationsInCtx(ctx *Context) error

	validatePromisesInCtx(ctx *Context) error
	
	expects() []string
	
	promises() []string
}

type Action struct {
	Actionable

	hasExecuted bool
}

func (act Action) validatePromisesInCtx(ctx *Context) error {
	return validateInCtx(ctx, act.Actionable.promises(), "promises")
}

func (act Action) validateExpectationsInCtx(ctx *Context) error {
	return validateInCtx(ctx, act.Actionable.expects(), "expectations")
}


func (act *Action) Execute(values map[string]any) *Context {
	ctx := &Context{Values: values}
	
	err := act.Actionable.validateExpectationsInCtx(ctx)
	if err != nil {
		panic(err)
	}

	act.Actionable.executed(ctx)

	if ctx.ShouldRollback {
		act.Actionable.rolledBack(ctx)
	}
	
	err = act.Actionable.validatePromisesInCtx(ctx)
	if err != nil {
		panic(err)
	}
	
	act.hasExecuted = true

	if(!ctx.Failure && !ctx.Success) {
		ctx.Succeed("")
	}
	
	return ctx
}

func (Action) executed(ctx *Context) {}

func (Action) rolledBack(ctx *Context) {}

func (Action) expects() []string {
	return []string{}
}

func (Action) promises() []string {
	return []string{}
}
