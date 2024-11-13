package goservice

type GoServiceContext interface {
	Fail(message string, errorCode int)
	
	Succeed(message string)
	
	FailAndReturn(message string, errorCode int)
	
	FailWithRollback(message string, errorCode int)
	
	SkipRemaining(message string)
	
	AddValue(key string, value any)
	
	ShouldStopProcessing() bool
}

type Context struct {
	Values    	  		map[string]any
	Success       		bool
	Failure       	    bool
	ShouldSkipRemaining bool
	ShouldRollback		bool
	Message   	  		string
	ErrorCode 	  		int
}

func (ctx *Context) AddValue(key string, value any) { // Maybe return an error?
	ctx.Values[key] = value
}

// Weird decision here, but better to have required args then none at all...
func (ctx *Context) Fail(message string, errorCode int) {
	ctx.Success   = false
	ctx.Failure   = true
	ctx.Message   = message
	ctx.ErrorCode = errorCode
}

func (ctx *Context) FailAndReturn(message string, errorCode int) {
	ctx.Fail(message, errorCode)
	
	// How would you signal a return in an action from this?
	// This might not be possible
}

func (ctx *Context) SkipRemaining(message string) {
	ctx.ShouldSkipRemaining = true
	ctx.Message 			= message
}

func (ctx *Context) FailWithRollback(message string, errorCode int) {
	ctx.ShouldRollback = true

	ctx.Fail(message, errorCode)
}

func (ctx *Context) Succeed(message string) {
	ctx.Success = true
	ctx.Failure = false
	ctx.Message = message
}

func (ctx Context) ShouldStopProcessing() bool {
	return ctx.Failure || ctx.ShouldSkipRemaining
}
