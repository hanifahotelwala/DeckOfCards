package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

}

/*

Slices:

	cards := []string{newCard(), newCard()} //slice of strings
	cards = append(cards, "Six of spades") //add to slice & returns a NEW slice.



Type conversion:

greeting := "hi there"
fmt.Println([]byte(greeting))

output: [104 105 32 116 104 101 114 101]

*/
