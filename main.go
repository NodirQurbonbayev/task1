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
package main

import (
	"fmt"
)
// Yig'indi,ko'paytmasini va kvadrat hisoblash 
func main ()  {
	var a int
	var b int 
	fmt.Scan(&a,&b)
	Yigindi:=a+b
	kopaytmasi:=a*b
	kvadrat:=a*a*b*b
	fmt.Println("Yigindi=",Yigindi)
	fmt.Println("kopaytmasi=",kopaytmasi)
	fmt.Println("kvadrat=",kvadrat)
}
package main

import (
	"fmt"
	"math"
)

// Yig'indi,ko'paytmasi va moduli hisoblash
func main() {
	var a, b int
	fmt.Scan(&a, &b)
	Yigindi := a + b
	kopaytmasi := a * b
	moduli := math.Abs(float64(a) + float64(b))
	fmt.Println("yig'indi=", Yigindi)
	fmt.Println("kopaytmasi=", kopaytmasi)
	fmt.Println("moduli=", moduli)
}
package main

import (
	"fmt"
	"math"
)
// gipotenuza va perimetr hisoblash
func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	c := math.Sqrt(a*a + b*b)
	P := a + b + c

	fmt.Printf("gipotenuza= %.2f\n", c)
	fmt.Println("perimetr=", P)

}
package main 

import "fmt"

func main ()  {
	var r1,r2 int 
	var p float32=3.14
	fmt.Scan(&r1,&r2)
		s1:=float32(r1)*p
		s2:=float32(r2)*p
		s3:=p*float32(r1-r2)
		fmt.Println("s1=",s1)
		fmt.Println("s2=",s2)
		fmt.Println("s3=",s3)
}
package main

import (
    "fmt"
    "math"
)
// Radiusi va yuzasini hisoblash 
func main() {
    var L float64
    fmt.Scan(&L)
    
    R := L / (2 * math.Pi)
    S := math.Pi * math.Pow(R, 2)
    
    fmt.Printf("Radiusi: %.2f\n", R)
    fmt.Printf("Yuzasi: %.2f\n", S)
}
package main 

import (
	"fmt"
	"math"
)	
//Diametr  va radiusini hisoblash 
func main ()  {
	var S float64
	fmt.Scan(&S)
	R:=math.Sqrt( S/(2 * math.Pi))
	d:=2*R
	fmt.Println(R)
	fmt.Println(d)
}
package main

import (
	"fmt"
	"math"
)

// ikkita nuqta orasidagi masofani topish
func main() {
	var x, y float64
	fmt.Scan(&x, &y)
	distance := math.Abs(x - y)
	fmt.Println(distance)

}
