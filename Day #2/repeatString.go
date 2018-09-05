package end

import (
	"strings")

func RepeatString(str string, times int)  string {
	var result strings.Builder
	for i := 0; i < times; i++ {
		result.Write([]byte(str))
	}
	return result.String()
}