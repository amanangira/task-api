package router

import "fmt"

func AddQueryParams(template string, args ...interface{}) string {
	var parsedArgs []interface{}
	for i := range args {
		parsedArgs = append(parsedArgs, fmt.Sprintf("{%s}", args[i]))
	}

	return fmt.Sprintf(template, parsedArgs...)
}
