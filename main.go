package main
import "fmt"

func main() {

	// implicit setting a var
	temp := 20

	// semi-implicit setting a var
	var t = 0


	// explicit setting a var
	var b int = 30

	fmt.Println(b)
	fmt.Println("hello world")
	fmt.Print(temp)
	for t = 0; t <= 9; t++ {
		fmt.Println(t)
	}
}