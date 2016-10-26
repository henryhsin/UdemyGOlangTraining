package main

import "fmt"

func main() {

	for j := 2 ; j<= 9; j++{
		for k := 1 ; k<= 9; k++{
			fmt.Print(j," * ",k," = ",j*k," ")
		}
		fmt.Println()
	}



}