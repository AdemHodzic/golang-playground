package line

import (
	"strings"

	"github.com/fogleman/gg"
)

type Line struct {
	words   []Word
	spacing float64
}

type Word struct {
	text  string
	color string
}

func (line *Line) Print(context *gg.Context, content string, x, y float64) {
	line.populate(content)
	line.spacing = 10
	used := x
	for _, word := range line.words {
		context.SetHexColor(word.color)
		w, _ := context.MeasureString(word.text)
		context.DrawString(word.text, used, y)
		used += w + line.spacing
	}
}

func (line *Line) populate(text string) {
	list := strings.Split(text, " ")
	for _, elem := range list {
		var newWord Word
		newWord.text = elem
		newWord.findColor()
		line.words = append(line.words, newWord)
	}
}

func (word *Word) findColor() {
	// Temp map
	syntax := map[string]string{
		"for":      "#f442b9",
		"if":       "#f442b9",
		"in":       "#2e31f4",
		"function": "#f7e222",
		"default":  "#eeeeee",
		"variable": "#14fc90",
	}

	if word.isFunction() {
		word.color = syntax["function"]
		return
	}

	if color, ok := syntax[word.text]; ok {
		word.color = color
	} else if word.isVariable() {
		word.color = syntax["variable"]
	} else {
		word.color = syntax["default"]
	}
}

func (word *Word) isFunction() bool {
	return strings.Contains(word.text, "(")
}

func (word *Word) isVariable() bool {
	for _, char := range word.text {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
			return false
		}
	}
	return true
}
