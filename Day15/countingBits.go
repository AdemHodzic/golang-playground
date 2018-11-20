package main

import(
	"strconv"
	"strings"
	"fmt"
)

func main() {
	fmt.Printf("%v", countBits(2))
}

func countBits(num int ) []int {

var list []int = make([]int, num + 1)

    
    for i := 0; i <= num; i++ {
        bin := strconv.FormatInt(int64(i) ,2)  
        list[i] = strings.Count(bin, "1")
    }
    
    return list


}
