package main

import (
	"container/list"
	"errors"
	"fmt"
	"strings"
)

func one() {
	f := countMemo()
	for i := 1; i <= 9; i++ {
		if f(i) != countRec(i) {
			panic("this shit is busted")
		}
	}
}

func countMemo() func(int) int {
	var c map[int]int = make(map[int]int)
	c[0] = 0
	c[1] = 1
	c[2] = 2
	c[3] = 4
	return func(n int) int {
		if v, ok := c[n]; ok {
			return v
		} else {
			for i := len(c); i <= n; i++ {
				c[i] = c[i-1] + c[i-2] + c[i-3]
			}
			return c[n]
		}
	}
}

func countRec(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n == 3 {
		return 4
	} else {
		return countRec(n-1) + countRec(n-2) + countRec(n-3)
	}
}

type point struct{ r, c int }
type path map[point]bool
type grid map[point]bool

func two() {
	var g grid = make(map[point]bool)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			g[point{i, j}] = true
		}
	}
	g[point{3, 1}] = false
	g[point{1, 2}] = false
	var goal point = point{3, 3}
	var robot point = point{0, 0}
	var p path = make(path)
	p[goal] = true
	p = findPath(g, goal, robot, p)
	fmt.Println(p)
}

func findPath(g grid, current point, destination point, p path) path {
	if destination == current {
		p[current] = true
		return p
	}
	var mu point = point{current.r - 1, current.c}
	var ml point = point{current.r, current.c - 1}
	muv, muok := g[mu]
	mlv, mlok := g[ml]
	if mlv && mlok {
		p[ml] = true
		return findPath(g, ml, destination, p)
	} else if muv && muok {
		p[mu] = true
		return findPath(g, mu, destination, p)
	} else {
		fmt.Println(g, mu, ml, current, destination, p)
		panic("oops")
	}
}

func three() {
	var a0, a1, a2 []int = []int{1, 5, 8, 9}, []int{-12, -3, 44, 27}, []int{-4, 0, 2, 8, 9}
	v, e := findMagic(a0)
	fmt.Println(v, e)
	v, e = findMagic(a1)
	fmt.Println(v, e)
	v, e = findMagic(a2)
	fmt.Println(v, e)
}

func findMagic(A []int) (int, error) {
	if len(A) == 0 {
		return 0, errors.New("did not find it")
	}

	i := midpoint(A)
	if A[i] == i {
		return A[i], nil
	} else if A[i] > i {
		A = A[:i]
		return findMagic(A)
	} else if A[i] < i {
		A = A[i+1:]
		return findMagic(A)
	} else {
		panic("this should not happen")
	}

	/*for i := 0; i < len(A); i++ {
		if i == A[i] {
			return i, nil
		}
	}
	return 0, errors.New("didn't find it") //XXX update to have typed error*/
}

func midpoint(A []int) int {
	return len(A) / 2
}

func four() {
	a0 := []int{0}
	a1 := []int{0, 1}
	a2 := []int{0, 1, 2}
	a3 := []int{0, 1, 2, 3}
	fmt.Println(subset(a0))
	fmt.Println(subset(a1))
	fmt.Println(subset(a2))
	fmt.Println(subset(a3))
}

func subset(a []int) (r [][]int) {
	r = make([][]int, 0)
	for i := 0; i < len(a); i++ {
		for j := i + 1; j <= len(a); j++ {
			r = append(r, a[i:j])
		}
	}
	return
}

func seven() {
	var s string = "abcdefghi"
	fmt.Println(stringSubset(s))
}

func stringSubset(a string) (r []string) {
	r = make([]string, 0)
	for i := 0; i < len(a); i++ {
		for j := i + 1; j <= len(a); j++ {
			r = append(r, a[i:j])
		}
	}
	return
}

func eight() {
	var s string = "aabcdeffghi"
	var d string = removeDup(s)
	fmt.Println(stringSubset(d))
}

func removeDup(s string) (r string) {
	var m map[byte]bool = make(map[byte]bool)
	var sb strings.Builder
	var b byte
	for i := range s {
		b = byte(s[i])
		if _, ok := m[b]; !ok {
			m[b] = true
			sb.WriteByte(b)
		}
	}
	return sb.String()
}

func eleven() {
	fmt.Println(change(21, [4]int{25, 10, 5, 1}, 0))
}
func change(v int, cs [4]int, i int) (n int) {
	if i == len(cs) {
		return 1
	}
	c := cs[i]
	for j := 0; j*c <= v; j++ {
		n += change(v-j*c, cs, i+1)
	}
	return
}

func twelve() {
	//gobuild()
}

const rowCount = int(8)

func gobuild() {
	var s [rowCount]int
	l := list.New()
	placeQueen(0, s, l)
	e := l.Front()
	for e != nil {
		fmt.Println(e.Value.([rowCount]int))
		e = e.Next()
	}
}
func placeQueen(r int, s [rowCount]int, l *list.List) {
	if r == rowCount {
		l.PushBack(s)
	} else {
		for c := 0; c < rowCount; c++ {
			if validForQueen(r, c, s) {
				s[r] = c
				placeQueen(r+1, s, l)
			}
		}
	}
}
func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
func validForQueen(r, c int, s [rowCount]int) bool {
	for i := 0; i < r; i++ {
		if s[i] == c {
			return false
		}
		if abs(r-i) == abs(c-s[i]) {
			return false
		}
	}
	return true
}

func main() {
	//one()   // success
	//two()   // success
	//three() // success
	//four()  // did without recursion by noticing iterative pattern
	//five() // see github towerOfHanoi
	//seven()
	//eight()
	eleven()
	//twelve() // see github nQueens
	twelve()
}
