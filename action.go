package goservice

type GoServiceAction interface {
	Execute(values map[string]any) (GoServiceContext, error) // makes a context and calls executed with said context.

	rollback(ctx GoServiceContext)

	executed(ctx GoServiceContext)

	rolledBack(ctx GoServiceContext)

	validateExpectationsInCtx(ctx GoServiceContext)

	validatePromisesInCtx(ctx GoServiceContext)
}

type Action struct {
	hasExecuted bool
	Expects 	[]string
	Promises 	[]string
}

func (act Action) validatePromisesInCtx(ctx *Context) error {
	return validateInCtx(ctx, act.Promises, "promises")
}

func (act Action) validateExpectationsInCtx(ctx *Context) error {
	return validateInCtx(ctx, act.Expects, "expectations")
}


func (act *Action) Execute(values map[string]any) (GoServiceContext, error) {
	ctx := &Context{Values: values}
	
	err := act.validateExpectationsInCtx(ctx)
	if err != nil {
		return nil, err
	}
	
	act.executed(ctx)
	
	if ctx.ShouldRollback {
		act.rolledBack(ctx)
	}
	
	err = act.validatePromisesInCtx(ctx)
	if err != nil {
		return nil, err
	}
	
	act.hasExecuted = true
	
	return ctx, nil
}

func (Action) executed(ctx *Context) {}

func (Action) rolledBack(ctx *Context) {}
