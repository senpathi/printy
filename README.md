# printy (in development)
This library prints strings in pre-defined shapes

### Rectangle
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

### Triangle
```go
package main

import "github.com/senpathi/printy"

func main() {
	str := `This is a sample string to print in senpathi/printy`
	printy.PrintInTriangle(str)
}
```
```text
Output:
            *
           * *
          *   *
         *  T  *
        *  his  *
       *   is a  *
      *   sample  *
     *   string t  *
    *  o print in   *
   *  senpathi/prin  *
  *  ty               *
 * * * * * * * * * * * *
```