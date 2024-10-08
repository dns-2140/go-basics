package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	fmt.Println("hello world")

	// arraySign([]int{2, 1})                  // 1
	// arraySign([]int{-2, 1})                   // -1
	// arraySign([]int{-1, -2, -3, -4, 3, 2, 1}) // 1

	// isAnagram("anak", "kana") // true
	// isAnagram("anak", "mana") // false
	// isAnagram("anagram", "managra") // true

	// findTheDifference("abcd", "abcde") // 'e'
	// findTheDifference("abcd", "abced") // 'e'
	// findTheDifference("", "y")         // 'y'

	// canMakeArithmeticProgression([]int{1, 5, 3})    // true; 1, 3, 5 adalah baris aritmatik +2
	// canMakeArithmeticProgression([]int{5, 1, 9})    // true; 9, 5, 1 adalah baris aritmatik -4
	// canMakeArithmeticProgression([]int{1, 2, 4, 8}) // false; 1, 2, 4, 8 bukan baris aritmatik, melainkan geometrik x2

	tesDeck()
}

// https://leetcode.com/problems/sign-of-the-product-of-an-array
func arraySign(nums []int) int {
	negativeMonitor := 1
	for _, val := range nums {
		if val == 0 {
			return 0
		}

		if val < 0 {
			negativeMonitor *= -1
		}

	}
	return negativeMonitor
}

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	// write code here

	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)
	// list first string
	for _, val := range s {
		m[val]++
	}
	// decrease with second string
	for _, val := range t {
		v, exist := m[val]

		if !exist || v == 0 {
			return false
		}

		m[val]--
	}

	return true
}

// https://leetcode.com/problems/find-the-difference
func findTheDifference(s string, t string) byte {
	// write code here
	m := make(map[rune]int)
	for _, val := range s {
		m[val]++
	}

	for _, val := range t {
		count, exist := m[val]
		if !exist || count == 0 {
			return byte(val)
		} else {
			m[val]--
		}
	}

	return 1
}

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence
func canMakeArithmeticProgression(arr []int) bool {
	// write code here

	sort.Ints(arr)

	diff := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != diff {
			return false
		}
	}
	return true
}

// Deck represent "standard" deck consist of 52 cards
type Deck struct {
	cards []Card
}

// Card represent a card in "standard" deck
type Card struct {
	symbol int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // Ace: 1, Jack: 11, Queen: 12, King: 13
}

// New insert 52 cards into deck d, sorted by symbol & then number.
// [A Spade, 2 Spade,  ..., A Heart, 2 Heart, ..., J Diamond, Q Diamond, K Diamond ]
// assume Ace-Spade on top of deck.
func (d *Deck) New() {
	// write code here
	symbols := []int{0, 1, 2, 3}
	for _, symbol := range symbols {
		for number := 1; number <= 13; number++ {
			d.cards = append(d.cards, Card{symbol: symbol, number: number})
		}
	}
}

// PeekTop return n cards from the top
func (d Deck) PeekTop(n int) []Card {
	// write code here
	return d.cards[len(d.cards)-n:]
}

// PeekTop return n cards from the bottom
func (d Deck) PeekBottom(n int) []Card {
	// write code here
	return d.cards[:n]
}

// PeekCardAtIndex return a card at specified index
func (d Deck) PeekCardAtIndex(idx int) Card {
	return d.cards[idx]
}

// Shuffle randomly shuffle the deck
func (d *Deck) Shuffle() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Fisher-Yates shuffle algorithm
	n := len(d.cards)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)                           // Generate a random index from 0 to i
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i] // Swap cards
	}
}

// Cut perform single "Cut" technique. Move n top cards to bottom
// e.g. Deck: [1, 2, 3, 4, 5]. Cut(3) resulting Deck: [4, 5, 1, 2, 3]
func (d *Deck) Cut(n int) {
	// write code here
	// Split the deck into two parts
	top := d.cards[:n]    // Top n cards
	bottom := d.cards[n:] // Remaining cards

	// Reassemble the deck: bottom part first, then top part
	d.cards = append(bottom, top...)
}

func (c Card) ToString() string {
	textNum := ""
	switch c.number {
	case 1:
		textNum = "Ace"
	case 11:
		textNum = "Jack"
	case 12:
		textNum = "Queen"
	case 13:
		textNum = "King"
	default:
		textNum = fmt.Sprintf("%d", c.number)
	}
	texts := []string{"Spade", "Heart", "Club", "Diamond"}
	return fmt.Sprintf("%s %s", textNum, texts[c.symbol])
}

func tesDeck() {
	deck := Deck{}
	deck.New()

	top5Cards := deck.PeekTop(5)
	fmt.Println("PeekTop 5")
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}
	fmt.Println("---\n")

	fmt.Println("PeekCardAtIndex index array 10 - 15")
	fmt.Println(deck.PeekCardAtIndex(10).ToString()) // Jack Spade
	fmt.Println(deck.PeekCardAtIndex(11).ToString()) // Queen Spade
	fmt.Println(deck.PeekCardAtIndex(12).ToString()) // King Spade
	fmt.Println(deck.PeekCardAtIndex(13).ToString()) // Ace Heart
	fmt.Println(deck.PeekCardAtIndex(14).ToString()) // 2 Heart
	fmt.Println(deck.PeekCardAtIndex(15).ToString()) // 3 Heart
	fmt.Println("---\n")

	deck.Shuffle()
	top5Cards = deck.PeekTop(5)
	fmt.Println("Deck Shuffle")
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}

	fmt.Println("---\n")
	deck.New()
	deck.Cut(5)
	bottomCards := deck.PeekBottom(10)
	fmt.Println("Deck Cut")
	for _, c := range bottomCards {
		fmt.Println(c.ToString())
	}
}
