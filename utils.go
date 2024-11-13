package goservice

import (
	"errors"
	"strings"
	"fmt"
)

func includes(slice []string, element string) bool {
    for _, item := range slice {
        if item == element {
            return true
        }
    }
    return false
}

func validateInCtx(ctx *Context, values []string, keyType string) error {
	keysInCtx 	  := make([]string, 0, len(ctx.Values))
	missingValues := make([]string, 0)
	
	for key := range ctx.Values {
        keysInCtx = append(keysInCtx, key)
    }
    
    for _, value := range values {
    	if !includes(keysInCtx, value) {
     		missingValues = append(missingValues, value)
     	}
    }
    
    if len(missingValues) > 0 {
    	errorMsg := fmt.Sprintf("Missing %s in context: %s", keyType, strings.Join(missingValues, ", "))
    
    	return errors.New(errorMsg)
    }
    
    return nil
}