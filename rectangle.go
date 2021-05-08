package printy

import (
	"fmt"
	"math"
)

func PrintInRectangle(text string, lengthFactor, widthFactor int) {
	length, width := calculateLengthAndWidth(len(text), lengthFactor, widthFactor)
	fmt.Printf(getRectangleString(text, length+1, width))
}

func getRectangleString(text string, length, width int) string {
	str := ""
	count := 0
	for h := 1; h <= width; h++ {
		for w := 1; w <= length; w++ {
			if h == 1 || h == width {
				str += "*  "
				continue
			} else if w == length-1 {
				str += "   "
				continue
			} else if w == 1 || w == length {
				str += "*  "
				continue
			} else {
				if count == len(text)-1 {
					str = str + text[count:count+1] + "  "
					count = count + 1
					continue
				}
				if count == len(text)-2 {
					str = str + text[count:count+2] + " "
					count = count + 2
					continue
				}
				if count > len(text)-3 {
					str += "   "
					continue
				}
				str += text[count : count+3]
				count = count + 3
			}
		}
		str += "\n"
	}

	return str
}

func calculateLengthAndWidth(area int, lFac int, wFac int) (int, int) {
	width := math.Sqrt(float64(area*wFac) / float64(lFac))
	width = math.Ceil(width)
	length := width * float64(lFac) / float64(wFac)

	return int(length), int(width)
}
