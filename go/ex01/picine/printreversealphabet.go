package picine

import "fmt"

func printreversealphabet() {
	for text:= 'z'; text <= 'a'; text--{
		fmt.Printf("%c", text)
	}
	fmt.Println()
}