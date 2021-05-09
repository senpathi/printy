package printy

import (
	"fmt"
	"math"
)

func PrintInTriangle(text string) {
	height := calculateTriangleHeight(len(text))
	fmt.Printf(getTriangleString(text, height))
}

func getTriangleString(text string, height int) string {
	str := ""
	count := 0
	for i := 0; i < height; i++ {
		for j := 0; j < height-i; j++ {
			str += " "
		}
		for k := 0; k <= 2*i; k++ {
			if i == height-1 && k%2 != 0 {
				str += " "
				continue
			}
			if i == height-1 || k == 0 || k == 2*i {
				str += "*"
				continue
			}
			if i > 2 && k > 2 && k < 2*i-2 && count < len(text) {
				str += text[count : count+1]
				count += 1
				continue
			}
			str += " "
		}
		str += "\n"
	}

	return str
}

func calculateTriangleHeight(area int) int {
	return int(math.Ceil(math.Sqrt(float64(area))))
}
