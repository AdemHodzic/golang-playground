package main

import (
	"golang-playground/Day14/models"
	"strings"

	"github.com/fogleman/gg"
)

var (
	width  float64 = 1024
	height float64 = 768
	text   []string
)

func main() {

	context := gg.NewContext(int(width), int(height))
	drawWindow(context)
	drawButtons(context)
	text = reader.ReadFile("buble.py")
	writeToWindow(text, context)
	context.SavePNG("out.png")
}

func drawWindow(context *gg.Context) {
	context.SetRGB(100, 100, 128)
	context.Clear()
	context.DrawRectangle(width*0.1, height*0.05, width*0.8, height*0.05)
	context.SetRGB255(0, 0, 0)
	context.Fill()
	context.DrawRectangle(width*0.1, height*0.1, width*0.8, height*0.8)
	context.SetRGB255(30, 30, 30)
	context.Fill()
}

func drawButtons(context *gg.Context) {

	context.SetRGB255(242, 31, 31)
	context.DrawCircle(width*0.87, height*0.07, 7)
	context.Fill()

	context.SetRGB255(162, 255, 109)
	context.DrawCircle(width*0.82, height*0.07, 7)
	context.Fill()

	context.SetRGB255(249, 255, 84)
	context.DrawCircle(width*0.77, height*0.07, 7)
	context.Fill()
}

func writeToWindow(text []string, context *gg.Context) {
	context.SetRGB255(240, 255, 43)
	context.LoadFontFace("Roboto-Black.ttf", 18)
	for index, line := range text {
		x := width*0.2 + countTabs(line)*0.02*width
		y := height*0.2 + (height*0.05)*float64(index)
		line = strings.Replace(line, "\t", "", -1)
		context.DrawString(line, x, y)
	}
}

func countTabs(line string) float64 {

	return float64(strings.Count(line, "\t"))
}
