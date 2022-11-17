package main

import "fmt"

/*
TO run -

	go run main.go deck.go card.go
*/
func main() {
	var deck = newDeck()
	fmt.Println("ALL CARDS = ", deck)

	var split1, split2 = deck.splitInTwo(5)
	fmt.Println("After splitting at index = 5, SPLIT1 = ", split1)
	fmt.Println("After splitting at index = 5, SPLIT2 = ", split2)

	convertStringToByteArray("hemant rocks")
	deck.saveToFile("hem.txt")
}

// Needed this function to read/write to file
func convertStringToByteArray(str string) {
	var byteArr = []byte(str)
	fmt.Println("String = ", str, " AND its byteArr =", byteArr)
}
