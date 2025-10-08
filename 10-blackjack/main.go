package main

import (
	"fmt"

	"example.com/deck"
)

func calculateScore(hand []deck.Card) int {
	score := 0
	numAces := 0
	for _, card := range hand {
		if card.Rank == 1 {
			numAces++
		} else if card.Rank > 10 {
			score += 10
		} else {
			score += card.Rank
		}
	}

	for i := 0; i < numAces; i++ {
		if score+11 > 21 {
			score += 1
		} else {
			score += 11
		}
	}
	return score
}

func main() {
	d := deck.New(
		deck.WithShuffle(true),
	)
	yourHand := make([]deck.Card, 0)
	dealerHand := make([]deck.Card, 0)
	for range 2 {
		yourHand = append(yourHand, d[0])
		dealerHand = append(dealerHand, d[1])
		d = d[2:]
	}

	stand := false
	playAgain := true
	for playAgain {
		fmt.Println("Your hand:", yourHand)
		fmt.Println("Your score:", calculateScore(yourHand))
		fmt.Println("Dealer's hand:", dealerHand[0], "and a hidden card")
		fmt.Println("Dealer's score:", calculateScore(dealerHand[:1]))
		stand = false
		for !stand {
			fmt.Println("Do you want to (h)it or (s)tand?")
			var input string
			fmt.Scanln(&input)
			switch input {
			case "h":
				card := d[0]
				d = d[1:]
				yourHand = append(yourHand, card)
				fmt.Println("Your hand:", yourHand)
				fmt.Println("Your score:", calculateScore(yourHand))
				if calculateScore(yourHand) > 21 {
					fmt.Println("You bust!")
				}
			case "s":
				stand = true
			default:
				fmt.Println("Invalid input")
				continue
			}
		}
		if stand {
			fmt.Println("Dealer's hand:", dealerHand)
			fmt.Println("Dealer's score:", calculateScore(dealerHand))
			for calculateScore(dealerHand) < 17 {
				card := d[0]
				d = d[1:]
				dealerHand = append(dealerHand, card)
				fmt.Println("Dealer's hand:", dealerHand)
				fmt.Println("Dealer's score:", calculateScore(dealerHand))
			}
			fmt.Println("Your hand:", yourHand)
			fmt.Println("Your score:", calculateScore(yourHand))

			if calculateScore(dealerHand) > 21 {
				fmt.Println("Dealer busts!")
				fmt.Println("You win!")
			} else if calculateScore(yourHand) > calculateScore(dealerHand) {
				fmt.Println("You win!")
			} else if calculateScore(yourHand) < calculateScore(dealerHand) {
				fmt.Println("You lose!")
			} else {
				fmt.Println("It's a tie!")
			}
		}
		fmt.Println("Do you want to play again? (y/n)")
		var input string
		fmt.Scanln(&input)
		if input == "n" {
			playAgain = false
		}
	}
}
