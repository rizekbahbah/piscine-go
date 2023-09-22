package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func countCards(deck []int) int {
	count := 0
	for range deck {
		count++
	}
	return count
}

func DealAPackOfCards(deck []int) {
	numPlayers := 4
	deckSize := countCards(deck)
	cardsPerPlayer := deckSize / numPlayers

	for i := 0; i < numPlayers; i++ {
		fmt.Printf("Player %d:", i+1)
		for j := 0; j < cardsPerPlayer; j++ {
			cardIndex := i*cardsPerPlayer + j
			fmt.Printf(" %d", deck[cardIndex])
			if j != cardsPerPlayer-1 {
				z01.PrintRune(',')
			}
		}
		z01.PrintRune('\n')
	}
}
