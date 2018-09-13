package main 

import (
	"fmt"
)


type Calendar struct {

	dates []Date
}

type Date struct {
	start int
	end int
}

func (date *Date) Overlap(number int) bool {
	return number > date.start && number < date.end
}


func (calendar *Calendar) Book(start int, end int)  bool {

	for _, date := range calendar.dates {

		if date.Overlap(start) || date.Overlap(end) {
			return false
		}
		

	}

	calendar.dates = append(calendar.dates, constructDate(start, end))
	return true
}


func constructDate(start int, end int) Date {
	var d Date
	d.start = start
	d.end = end
	return d
}



func main() {
	var calendar Calendar
	fmt.Printf("%v\n", calendar.Book(10,20))
	fmt.Printf("%v\n", calendar.Book(15,25))
	fmt.Printf("%v\n", calendar.Book(20,30))

}