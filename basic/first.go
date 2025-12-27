// // package main

// // import "fmt"

// // func swap(x, y string) (string, string) {
// // 	return y, x
// // }

// // func main() {

// // 	// var s string = "    hello husnain"

// // 	// name := "Hasnain"
// // 	// age := 21

// // 	// s := fmt.Sprintf("My name is %s and I am %d", name, age)
// // 	// fmt.Println( s)

// // 	// var numbers [5]int
// // 	//  fmt.Println(numbers)
// // 	//  fmt.Println(len(numbers))
// // 	// for index, char := range numbers {
// // 	// 	fmt.Println(index,char)
// // 	// }

// // 	var s []int
// // 	si:= []int{}
// // 	fmt.Println(si)
// // 	fmt.Println(s==nil)
// //   si = append(si, 3)
// // 	fmt.Println(si)

// // }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

//     // var a  string = "initial"
//     // fmt.Println(a)

//     // var b, c int = 1, 2
//     // fmt.Println(b, c)

//     // var d = true
//     // fmt.Println(d)

//     // var e int
//     // fmt.Println(e)

//     // f := "apple"
//     // f="Mang"

//     // fmt.Println(f)
//     //  for i := range 3 {
//     //     fmt.Println("range", i)
//     // }

// 	// 	if (8%2==0 || 7%2==0){
// 	// 		fmt.Println("either 8 or 7 are even")
// 	// 	}

// 	// 	 t:= time.Now()
//     // switch {
//     // case t.Hour() < 12:
//     //     fmt.Println("It's before noon")
//     // default:
//     //     fmt.Println("It's after noon")
// 	// 	}
// }

package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    fmt.Println(*iptr)
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)
}