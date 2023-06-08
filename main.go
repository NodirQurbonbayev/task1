package main

import "fmt"

// perimetrni hisoblash
func main() {
	var a int
	fmt.Scan(&a)
	var p int = 4 * a
	fmt.Println("Perimeter (p) =", p)
}

package main

import "fmt"
//Yuzini topish
func main()  {
	var a,S int 
	fmt.Scan(&a)
	S=a*a
	fmt.Println("Yuzi=",S)
}
package main

import "fmt"

//Yuzi va Perimeterni hisoblash
func main() {
	var a, b, S, P int
	fmt.Scan(&a, &b)
	S = a * b
	P = 2 * (a + b)
	fmt.Println("Yuzi=", S)
	fmt.Println("Perimeter=", P)
}
package main 

import  "fmt"
//L=p*d uzunligini topish 	
func main()  {
	var d int
	var p float32=3.14
	fmt.Scan(&d)
	L:=p*float32(d)
	fmt.Println(L)
}
package main

import "fmt"

// hajmi va s ni topish
func main() {
	var a, S, V int
	fmt.Scan(&a)
	V = a * a * a
	S = 6 * (a * a)
	fmt.Println("Hajm=", V)
	fmt.Println("Yuzi=", S)
}
package main

import "fmt"

func main()  {
	var a,b,c int 
	fmt.Scan(&a,&b,&c)
	V:=a*b*c
	S:=2*(a*b+b*c+a*c)
	fmt.Println("Hajm=",V)
	fmt.Println("Yuzi=",S)
}
package main 

import "fmt"
// yuzi va uzunligini topish
func main()  {
	var r int=2 
	var p=3.14 float64
	L:=2*p*r 
	S:=p*r*r
	fmt.Println("Uzunligi=",L)
	fmt.Println("Yuzi=",S)
}

package main

import "fmt"
//O'rta arifmetigi
func main() {
    var a,b,sum int 
	fmt.Scan(&a,&b)
	sum=(a+b)/2
	fmt.Println("Yig'indi",sum)

}
package main

import (
	"fmt"
	"math"
)

// o'rta geometrigi
func main() {
	var num1, num2 float64
	fmt.Scan(&num1, &num2)

	sqrt1 := math.Sqrt(num1)
	sqrt2 := math.Sqrt(num2)
	sqrt := sqrt1 * sqrt2

	fmt.Printf(" o'rta geometr= %.2f", sqrt)
}