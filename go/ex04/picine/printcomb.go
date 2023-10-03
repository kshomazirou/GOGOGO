package picine

import "fmt"

func PrintComb() {
	for i:= 0; i <= 9; i++{
		for j:= i + 1; j <= 9; j++{
			for k:= j + 1; j <= 9; j++{
				fmt.Printf("%d%d%d", i,j,k)
				if (i == 7){
				fmt.Printf(", ")
				}
			}
		}
	}
	fmt.Println()
}