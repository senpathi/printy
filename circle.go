package printy

import (
	"fmt"
	"math"
)

func PrintInCircle(text string) {
	radius := dd(len(text)) + 2
	diameter := (radius * 2) + 1
	tolerance := 0.5
	count := 0
	dCount := 0

	for y := 1; y <= diameter; y++ {
		for x := 1; x <= diameter*3; x++ {
			if x > 198 {
				continue
			}
			a := math.Abs(float64(radius) - float64(x)/2.5 + 1)
			b := math.Abs(float64(radius - y + 1))
			c := math.Sqrt((a*a)+(b*b)) + tolerance
			//c_ := inDistance(x, y, float64(diameter), fudge)
			//c = c_
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
				dCount += 1
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
	fmt.Println(`count =`, dCount)
	dd(dCount)
}

// area printing circle = (5r + 2) / 2  * r * PI
func dd(a int) int {
	//A=5, B=2, C = (-2 * Area / PI)
	r := math.Ceil(math.Abs(-2 + (math.Sqrt(4+(4*5*(2*float64(a)/math.Pi))) / 10)))
	fmt.Println(`radius := `, r)
	fmt.Println(`area  := `, a)
	//r := math.Round(-2 + math.Sqrt(discriminant)/(2*a))
	return int(3)
}
