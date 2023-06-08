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