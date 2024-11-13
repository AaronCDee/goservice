package goservice

type GoServiceOrganizer interface {
	Reduce([]Action) GoServiceContext
	
	With(values map[string]any) GoServiceOrganizer
	
	Call(values map[string]any) GoServiceContext
}

type Organizer struct {
	ctx *Context
}

func (org Organizer) Reduce(actions []Action) GoServiceContext {
	for _, action := range actions {
		action.executed(org.ctx)
		
		if org.ctx.ShouldStopProcessing() {
			if org.ctx.ShouldRollback {
				for _, staleAction := range actions {
					if staleAction.hasExecuted {
						staleAction.rolledBack(org.ctx)
					}
				}
			}
			
			break
		}
	}
	
	if !org.ctx.Failure	&& !org.ctx.Success {
		org.ctx.Succeed("")
	}
	
	return org.ctx
}

func (org Organizer) With(values map[string]any) GoServiceOrganizer {
	ctx := &Context{Values: values}
	
	org.ctx = ctx
	
	return org
}


func (org Organizer) Call(values map[string]any) GoServiceContext {
	return org.ctx
}



// var animalConstructors = []func() Animal{
// 	func() Animal { return Dog{Name: "Rex"} },
// 	func() Animal { return Cat{Name: "Whiskers"} },
// }
