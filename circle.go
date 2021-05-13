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

// area printing circle = (5r + 2) / 2  * r * PI
func dd(area int) int {
	//A=5, B=2, C = (-2 * Area / PI)
	a := 7.96
	b := 6.72
	c := 2.04 - float64(area)

	r := (-b + math.Sqrt((b*b)-(4*a*c))) / (2 * a)
	//fmt.Print(`r = `, r)
	r = math.Ceil(r) //gofloat.ToFloat(r, 1).Float64()

	//fmt.Print(`   radius := `, r)
	//fmt.Print(`   area  := `, area)
	//r := math.Round(-2 + math.Sqrt(discriminant)/(2*a))

	//fmt.Println(`   true area := `, (a * r * r) + (b * r) + 1.2)
	return int(r)
}

//
