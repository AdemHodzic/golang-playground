package end

import "strings"

func Truncate(str string, len int) string {
	
	var builder strings.Builder
	
	builder.Write([]byte(str[:len]))
	builder.Write([]byte("..."))
	return builder.String()
}