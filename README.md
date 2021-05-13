# printy (in development)
This library prints strings in pre-defined shapes

### Rectangle
```go
package main

import "github.com/senpathi/printy"

func main() {
	str := `This is a sample string to print in senpathi/printy Rectangle`
	printy.PrintInRectangle(str, 16, 1)
}

```
```text
Output:
* * * * * * * * * * * * * * * * * ** *
*  This is a sample string to print  *
*   in senpathi/printy Rectangle     *
* * * * * * * * * * * * * * * * * ** *
```

### Triangle
```go
package main

import "github.com/senpathi/printy"

func main() {
	str := `This is a sample string to print in senpathi/printy Triangle`
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
  *  ty Triangle      *
 * * * * * * * * * * * *
```

### Circle
```bigquery
package main

import (
	"github.com/senpathi/printy"
)

func main() {
	str := `This is a sample string to print in senpathi/printy Circle`
	printy.PrintInCircle(str)
}
```
```text
Output: 
         ***** *****             
     **               **         
   *      This is a      *       
  *     sample string     *      
 *    to print in senpa    *     
 *    thi/printy Circle    *     
 *                         *     
  *                       *      
   *                     *       
     **               **         
         ***** *****   
```