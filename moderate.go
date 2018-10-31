package main

import (
	"fmt"
	"log"
	"math"
    "math/rand"
	"sort"
	"strings"
)

func one() {
	var i, j *int = new(int), new(int)
	*i, *j = 6, 7
	swap(i, j)
	fmt.Println(*i, *j)
}
func swap(i, j *int) {
	*i = *j ^ *i
	*j = *j ^ *i
	*i = *j ^ *i
}
func three() {
	/*
	   an algorithm to figure out if someone won tic tac toe

	   first there would need to be a representation of a board:
	   [3][3]uint8

	   row, column

	   we'll use uint8 because we need 0 for not played, 1 for X and 2 for O
	   can't use bool for that.

	   "winning" at tic tac toe is defined as three in a row in a horizontal line
	   or three in a row diagonally

	   so if [0][0-2] || [1][0-2] || [2][0-2] (horizontal)
	   or if [0-2][0] || [0-2][1] || [0-2][2] (veritcal)
	   or if [0][0], [1][1], [2][2] || [2][0], [1][1], [0][2]

	   we'd wanrt to check after each round in case somebody is real smart

	   so something like the following loop:

	   for {
	       if board.Winner() {
	           break
	       }
	       if board.Full() {
	           break
	       }
	       board.PlayerOne()
	       board.PlayerTwo()
	   }

	   since we've already defined the board, the function for board.Winner is just

	   if board.Twon(X) || board.Twon(O) {
	   return true
	   }
	   return false

	   then function is like this:

	   if [0][0]board == T && [0][1]board == T && [0][2]board == T {
	   }
	   this repeats for each of the above sets.  is there a way to further simplify?

	   could build a range function to take the variadic list of cells to check for each pair

	   type cell struct {
	       r, c uint8
	       }
	   board.CheckWin(T, cell{0, 0}, cell{0,1},cell{0,2})

	   but that's about the same amount of boilerplate

	   switch doesn't really help here

	   so how did I do this before (looks on internet archive of personal code)

	   did it same way

	   is there a nice way to improve this?

	   }
	*/

}

func five() {
	if countZeros(10) != 2 {
		fmt.Println("boo")
	} else if countZeros(15) != 3 {
		fmt.Println("boo")
	} else if countZeros(23) != 4 {
		fmt.Println("boo")
	}
	fmt.Println(factorial(10))
	fmt.Println(factorial(15))
	fmt.Println(factorial(20))
}
func countZeros(n int) int {
	return n / 5
}
func factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

const intMax = 1<<32 - 1

func six() {
	one, two := []int{1, 3, 15, 11, 2}, []int{23, 127, 235, 19, 8}
	three, four := []int{8, 4, 20, 18}, []int{45, 23, 1, 22}
	fmt.Println(lowestDiff(one, two))
	fmt.Println(lowestDiff(three, four))
	fmt.Println(optLowestDiff(one, two))
	fmt.Println(optLowestDiff(three, four))
}
func lowestDiff(a, b []int) (v int) {
	v = intMax
	for _, av := range a {
		for _, bv := range b {
			if av > intMax || bv > intMax {
				log.Fatal("number too big")
			}
			if isNeg(av - bv) {
				if (bv - av) < v {
					v = bv - av
				}
			} else if (av - bv) < v {
				v = av - bv
			}
		}
	}
	return
}
func isNeg(a int) bool {
	if a < 0 {
		return true
	}
	return false
}
func optLowestDiff(a, b []int) (v int) {
	v = 1<<32 - 1
	sort.Ints(a)
	sort.Ints(b)
	var t int
	for ai, bi := 0, 0; ai < len(a) && bi < len(b); {
		t = abs(a[ai] - b[bi])
		if t < v {
			v = t
		} else if t > v && (ai == len(a)-1 || bi == len(b)-1) {
			break
		}
		if a[ai] < b[bi] {
			ai++
		} else {
			bi++
		}
	}
	return v
}
func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func eight() {
	fmt.Println(englishNumber(132))
	fmt.Println(englishNumber(124315))
}

func englishNumber(n int) string {
	/*
	   divide by each "number" to get the english representation
	*/
	output := []string{}
	/*
	   need counter to track when done
	   need to divide by a thousand
	*/
	var countThousands int
	for n/10 > 0 {
		ones := baseCase(n)
		n = n / 10
		tens := baseCase(n)
		n = n / 10
		hundreds := baseCase(n)
		n = n / 10
		var sb strings.Builder
		sb.WriteString(englishHundreds(hundreds))
		teens, isTeen := englishTens(tens)
		sb.WriteString(teens)
		sb.WriteString(englishOnes(ones, isTeen))
		if countThousands > 4 {
			log.Fatal("not impl yet")
		} else if countThousands == 4 {
			sb.WriteString(" trillion ")
		} else if countThousands == 3 {
			sb.WriteString(" billion ")
		} else if countThousands == 2 {
			sb.WriteString(" million ")
		} else if countThousands == 1 {
			sb.WriteString(" thousand ")
		}
		countThousands++
		output = append(output, sb.String())
	}
	var sb strings.Builder
	for i := len(output) - 1; i >= 0; i-- {
		sb.WriteString(output[i])
	}
	return sb.String()
}
func baseCase(n int) int {
	return n % 10
}
func englishHundreds(n int) string {
	switch n {
	case 0:
		return ""
	case 1:
		return " one hundred "
	case 2:
		return " two hundred "
	case 3:
		return " three hundred "
	}
	log.Fatal("hundreds no worky")
	return ""
}

// returns true if teens
func englishTens(n int) (string, bool) {
	switch n {
	case 0:
		return "", false
	case 1:
		return "", true
	case 2:
		return " twenty ", false
	case 3:
		return " thirty ", false
	}
	log.Fatal("tens no worky")
	return "", false
}
func englishOnes(n int, teens bool) string {
	if !teens {
		switch n {
		case 0:
			return ""
		case 1:
			return " one "
		case 2:
			return " two "
		case 3:
			return " three "
		case 4:
			return " four "
		case 5:
			return " five "
		case 6:
			return " six "
		case 7:
			return " seven "
		case 8:
			return " eight "
		case 9:
			return " nine "
		}
	} else {
		switch n {
		case 0:
			return " ten "
		case 1:
			return " eleven "
		case 2:
			return " twelve "
		case 5:
			return " fifteen "
		}
	}
	log.Fatal("no worky")
	return ""
}

func nine() {
	fmt.Println((2 * 4), multiply(2, 4))
	fmt.Println((6 * 7), multiply(6, 7))
	fmt.Println((6 / 3), divide(6, 3))
	fmt.Println((12 / 4), divide(12, 4))
	fmt.Println((4 - 2), subtract(4, 2))
	fmt.Println((12 - 1), subtract(12, 1))
}
func multiply(a, b int) int {
	fmt.Println("multiply", a, b)
	if b == 1 {
		return a
	}
	t := a
	for i := 1; i < b; i++ {
		t += a
	}
	return t
}
func divide(a, b int) int {
	fmt.Println("divide", a, b)
	if a == b {
		return 1
	}
	i := 1
	t := b
	for {
		t += b
		i++
		if t == a {
			return i
		}
	}
}
func subtract(a, b int) int {
	fmt.Println("subtract", a, b)
	if a == b {
		return 0
	}
	i := 0
	for {
		b += 1
		i++
		if a == b {
			return i
		}
	}
}

func ten() {
	//https://play.golang.org/p/IrbvUFH_hkH
}

func thirteen() {
	sq0, sq1 := square{point{4, 4}, point{2, 2}}, square{point{8, 8}, point{6, 6}}
	b, l := bisect(sq0, sq1)
	log.Print(b, l.a, l.m, l.b)
}

type point struct {
	x, y float64
}
type square struct {
	ne, sw point
}

//ay=mx+b
type line struct {
	a, m, b float64
}

const (
	notBisectable = iota
	horizontal
	vertical
	diagonal
)

type biType int

func bisecType(sq0, sq1 square) biType {
	if sq0.ne.x == sq1.ne.x && sq0.sw.x == sq1.sw.x {
		return horizontal
	} else if sq0.ne.y == sq1.ne.y && sq0.sw.y == sq0.sw.y {
		return vertical
	} else if math.Abs(sq0.ne.x-sq1.ne.x) == math.Abs(sq0.sw.x-sq1.sw.x) && math.Abs(sq0.ne.y-sq1.ne.y) == math.Abs(sq0.sw.y-sq1.sw.y) {
		return diagonal
	} else {
		return notBisectable
	}
}
func bisect(sq0, sq1 square) (bool, *line) {
	b := bisecType(sq0, sq1)
	if b == notBisectable {
		return false, nil
	}
	line := new(line)
	if b == horizontal {
		line.a = 1
		line.b = (sq0.ne.y-sq0.sw.y)/2.0 + sq0.sw.y
		line.m = 0
	} else if b == vertical {
		line.a = 0
		line.m = 1
		line.b = -1.0*(sq0.ne.x-sq0.sw.x)/2.0 + sq0.sw.x
	} else if b == diagonal {
		// use the diagonal from a square to get the slope
		line.m = (sq0.ne.y - sq0.sw.y) / (sq0.ne.x - sq0.sw.x)
		// now figure out whether the slope is positive or negative
		if sq1.sw.x < sq0.sw.x && sq1.sw.y > sq0.sw.y || sq1.sw.x > sq0.sw.x && sq1.sw.y < sq0.sw.y {
			line.m = -1.0 * line.m
		}
		line.a = 1
		line.b = sq0.sw.y - line.m*sq0.sw.x
	} else {
		log.Fatal("no dice", sq0, sq1)
	}
	return true, line
}

func sixteen() {
	a := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	l, h := findLowHigh(a)
	fmt.Println(l, h)
}
func findLowHigh(a []int) (lowI, highI int) {
	lowI = len(a) + 1
	v := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < v {
			for j := 0; j < i; j++ {
				if a[j] > a[i] {
					if j < lowI {
						lowI = j - 1
					}
				}
			}
			highI = 0
		}
		if a[i] > v && highI == 0 {
			highI = i - 1
		}
		if a[i] > v {
			v = a[i]
		}
	}
	return
}

func twenty() {
	words := map[string]bool{"the": true, "vie": true, "tid": true}
	candidates := make(map[string]bool)
	buildCandidates("", [][]rune{[]rune{'t', 'u', 'v'}, []rune{'g', 'h', 'i'}, []rune{'d', 'e', 'f'}}, 0, candidates)
	for k, _ := range candidates {
		if _, ok := words[k]; ok {
			fmt.Println(k)
		}
	}
	//843
	//9224
}
func buildCandidates(prefix string, letters [][]rune, n int, candidates map[string]bool) {
	if n == len(letters) {
		return
	}
	var sb strings.Builder
	var s []string
	for _, suffix := range letters[n] {
		sb.WriteString(prefix)
		sb.WriteRune(suffix)
		candidates[sb.String()] = true
		s = append(s, sb.String())
		sb.Reset()
	}
	for _, w := range s {
		buildCandidates(w, letters, n+1, candidates)
	}
	return
}

func twentyThree() {
    fmt.Println(rand7())
}

func main() {
	one()
	five()
	six()
	eight()
	nine()
	ten()
	thirteen()
	sixteen()
	twenty()
    twentyThree()
}
