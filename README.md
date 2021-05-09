# printy (in development)
This library prints strings in pre-defined shapes

```go
package main

import "github.com/senpathi/printy"

func main() {
	str := `This is a sample string to print in senpathi/printy`
	printy.PrintInRectangle(str, 7, 2)
}

```

```text
Output:
* * * * * * * * * * * * * * * * *
*  This is a sample string to   *
*  print in senpathi/printy     *
* * * * * * * * * * * * * * * * *
```
