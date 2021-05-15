package printy

import (
	"fmt"
	"math"
)

// PrintInCircle prints sent text in a circle.
// Circle's size is determined according to the text's length
func PrintInCircle(text string) {
	radius := getRadius(len(text)) + 2
	diameter := (radius * 2) + 1
	tolerance := 0.5
	count := 0
	//dCount := 0

	for y := 1; y <= diameter; y++ {
		for x := 1; x <= diameter*3; x++ {
			if x > 198 {
				continue
			}
			a := math.Abs(float64(radius) - float64(x)/2.5 + 1)
			b := math.Abs(float64(radius - y + 1))
			c := math.Sqrt((a*a)+(b*b)) + tolerance

			if float64(c)-1 > float64(radius) {
				//fmt.Println(float64(c) - 1)
				fmt.Print(" ")
				continue
			}
			if c-tolerance > float64(radius) {
				fmt.Print("*")
			} else if c > float64(radius) {
				fmt.Print(" ")
			} else {

				if int(c) < radius-1 && count < len(text) {
					fmt.Print(text[count : count+1])
					count += 1
					continue
				}
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// getRadius calculated a circle's minimum radius to print text's length in it.
func getRadius(area int) int {
	a := 7.96
	b := 6.72
	c := 2.04 - float64(area)

	r := math.Ceil(-b+math.Sqrt((b*b)-(4*a*c))) / (2 * a)

	return int(r)
}

//
