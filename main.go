package main

import "fmt"

func main() {

	var x int64
	var y int64

	//get x and y vals
	for {
		fmt.Print("Enter height value: ")
		_, _ = fmt.Scanf("%d", &y)
		fmt.Println()
		if y > 0 {
			break
		}
	}

	for {
		fmt.Print("Enter width value: ")
		_, _ = fmt.Scanf("%d", &x)
		fmt.Println()
		if x > 0 {
			break
		}
	}

	playingField := make([][]bool, y+2)
	for i := range playingField {
		playingField[i] = make([]bool, x+2)
	}

}
