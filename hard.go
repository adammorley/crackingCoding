package main

import (
	"fmt"
	"math/rand"
    "strconv"
	"strings"
	"time"
)

const oneBits = uint8(8)

// 17.1 in cracking coding book; took about 40 mins
func one() {
	var a, b uint8 = 4, 4
	var c, d uint8 = 12, 12
	var e, f uint8 = 10, 3
	if add(a, b) != uint8(8) {
		fmt.Println("no dice for ab")
	} else if add(c, d) != uint8(24) {
		fmt.Println("no dice for cd")
	} else if add(e, f) != uint8(13) {
		fmt.Println("no dice for ef")
	} else {
		fmt.Println("got", add(a, b), add(c, d), add(e, f))
	}
}
func add(a, b uint8) (r uint8) {
	var carry uint8 = uint8(1)
	var i uint8
	for i = 0; i < oneBits; i++ { // could just use a bit vector here too
		var mask uint8 = ^uint8(0) >> 7 << i
		at, bt := a&mask, b&mask
		if at^bt == mask {
			r = r | at | bt
		} else if at&mask == mask && bt&mask == mask {
			// both one
			carry = carry << 1
			r = r & ^mask
		} else if carry > 1 {
			// when carried, how does that manifest?
			// what happens to carried bits when you have an xor == 1 or both are 1?
			if at&bt == 0 {
				// set the bits
				t := mask
				for carry != uint8(1) {
					carry = carry >> 1
					r = r | t
					t = t >> 1
				}
				carry = uint8(1)
			}
		}
	}
	return
}

func two() {
	fullDeck := newDeck()
    fullDeck.shuffle()
	for _, c := range fullDeck {
		fmt.Println(c)
	}
}

type deck []card
type card struct {
	suit int
	val  int
}

const (
	_     = iota
	Clubs = iota
	Diamonds
	Hearts
	Spades
)
const numberOfSuits = 4
const lowValue = 2
const faceValue = 11
const highValue = 14

func newDeck() deck {
	suits := [numberOfSuits]int{Clubs, Diamonds, Hearts, Spades}
	full := []card{}
	for _, suit := range suits {
		for i := lowValue; i <= highValue; i++ {
			full = append(full, card{suit: suit, val: i})
		}
	}
	return full
}
func (c card) lookupSuit() string {
	suits := map[int]string{1: "clubs", 2: "diamonds", 3: "hearts", 4: "spades"}
	return suits[c.suit]
}
func (c card) lookupFaceValue() string {
	faceValues := map[int]string{faceValue: "jack", 12: "queen", 13: "king", highValue: "ace"}
	if c.val < lowValue {
		panic("cards start at 2")
	} else if c.val < faceValue {
		return strconv.Itoa(c.val)
	} else if c.val <= highValue {
		return faceValues[c.val]
	} else {
		panic("cards stop at 14")
	}
}
func (c card) String() string {
	var sb strings.Builder
	sb.WriteString(c.lookupFaceValue())
	sb.WriteString(" of ")
	sb.WriteString(c.lookupSuit())
	return sb.String()
}
func (d deck) swap(i, j int) {
    d[i], d[j] = d[j],d[i]
}
func (d deck) randomCardIndex() int {
    return rand.Intn(len(d))
}
func (d deck) shuffle() {
    for i:=0;i<len(d);i++{
        d.swap(i, d.randomCardIndex())
    }
}

// 15 mins; 17.3 in cracking coding
func three() {
	var a []int = []int{5, 16, 2, 9, 22}
	var b []int = []int{3, 88, 1, 4, 55, 10}
	fmt.Println(randSet(a, 2), randSet(a, 3), randSet(a, 4))
	fmt.Println(randSet(b, 4), randSet(b, 2), randSet(b, 3))
}
func randSet(v []int, n int) (r []int) {
	rand.Seed(time.Now().UnixNano())
	var m map[int]bool = make(map[int]bool)
	for i := 0; i <= n; i++ {
		d := rand.Intn(len(v))
		_, ok := m[d]
		for ok {
			d = rand.Intn(len(v))
			_, ok = m[d]
		}
		m[d] = true
		r = append(r, v[d])
	}
	return
}

func six() {
	if countTwos(25) != 9 {
		fmt.Println("boo")
	} else if countTwos(250) != 106 {
		fmt.Println("boo")
	} else {
		fmt.Println("got", countTwos(25), countTwos(250))
	}
}
func countTwos(n int) (r int) {
	for i := 0; i <= n; i++ {
		c, j := 0, i
		for j > 0 {
			if j%10 == 2 {
				c++
			}
			j = j / 10
		}
		r += c
	}
	return
}

func main() {
	fmt.Println("hello")
	one()
	two()
	three()
	six()
}
