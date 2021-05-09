package printy

import (
	"fmt"
	"math"
)

func PrintInRectangle(text string, lengthFactor, widthFactor int) {
	length, width := calculateLengthAndWidth(len(text), lengthFactor, widthFactor)
	fmt.Printf(getRectangleString(text, length+2, width+2))
}

func getRectangleString(text string, length, width int) string {
	str := ""
	count := 0
	//lBorder := 2
	for w := 1; w <= width; w++ {
		for l := 1; l <= length; l++ {
			if w == 1 || w == width {
				if l == 1 {
					str = str + "* *"
					continue
				}
				if l == length {
					str = str + "* *"
					continue
				}
				if l%2 == 1 {
					str += "*"
					continue
				}
				str += " "
				continue
			}
			if l == 1 {
				str = str + fmt.Sprintf("%-3s", "*")
				continue
			}
			if l == length {
				str = str + fmt.Sprintf("%3s", "*")
				continue
			}

			if count > len(text)-1 {
				str += " "
				continue
			}
			str = str + text[count:count+1]
			count = count + 1
		}
		str += "\n"
	}

	return str
}

func calculateLengthAndWidth(area int, lFac int, wFac int) (int, int) {

	width := math.Sqrt(float64(area*wFac) / float64(lFac))
	length := width * float64(lFac) / float64(wFac)

	if int(math.Ceil(length)) > 60 {
		return 60, int(math.Ceil(float64(area) / 60))
	}

	return int(math.Ceil(length)), int(math.Ceil(width))
}
