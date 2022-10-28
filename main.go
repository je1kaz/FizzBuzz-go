package main
  
import "fmt"

func main()  {
var n int 
b := 5
c := 3
d := 15
f := 2
 fmt.Scanln(&n)
 if n % d == 0 {
	fmt.Println("FizzBuzz")
} else if  n % b == 0 {
	fmt.Println("Fizz")
 } else if n % c == 0 {
	fmt.Println ("Buzz")
 } else if n % f == 0 {
	fmt.Println("Wrong value")
 }
}
