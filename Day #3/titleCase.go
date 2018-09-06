package main

import ("strings"
)

func TitleCase(str string) string {
	
	arr := strings.Split(str, " ")
	var result []string
	
	for _, word := range arr {
		var builder strings.Builder
		upper := strings.ToUpper(string(word[0]))
		lower := strings.ToLower(string(word[1:]))
		builder.Write([]byte(upper))
		builder.Write([]byte(lower))
		result = append(result, builder.String())
	}
	return strings.Join(result, " ")
}