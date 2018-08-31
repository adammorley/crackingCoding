// cracking the coding interview sort problems
package main

import "container/list"
import "fmt"
import "math/rand"
import "sort"

func one() {
	var a []int = make([]int, 12)
	var b []int = []int{2, 6, 8, 33, 45, 99}
	a[0], a[1], a[2], a[3], a[4], a[5] = 1, 3, 5, 7, 9, 12
	for i := 6; i < len(a); i++ {
		a[i] = 1000000000
	}
	compare([]int{1, 2, 3, 5, 6, 7, 8, 9, 12, 33, 45, 99}, mergeOne(a, b))
}
func mergeOne(a, b []int) []int {
	var i, j int
	for {
		if a[i] < b[j] || a[i] == b[j] {
			i++
		} else if b[j] < a[i] {
			a = shiftRight(a, i)
			a[i] = b[j]
			j++
			i++
		}
		if j == len(b) {
			break
		}
	}
	return a
}
func shiftRight(a []int, i int) []int {
	for j := len(a) - 2; j >= i; j-- {
		a[j+1] = a[j]
	}
	return a
}

type word []byte

func (w word) Less(i, j int) bool {
	return w[i] < w[j]
}
func (w word) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}
func (w word) Len() int {
	return len(w)
}
func two() {
	var strs []string = []string{"ram", "binary", "cat", "mar", "brainy"}
	// map sorted characters to slice of strings
	var m map[string]*list.List = make(map[string]*list.List)
	for i := range strs {
		var w word = word(strs[i])
		sort.Sort(w)
		var key string = string(w)
		if _, ok := m[key]; !ok {
			m[key] = list.New()
		}
		m[key].PushBack(strs[i])
	}
	var r sort.StringSlice
	for _, v := range m {
		n := v.Front()
		for n != nil {
			r = append(r, n.Value.(string))
			n = n.Next()
		}
	}
	fmt.Println(strs, r)
}

func three() {
	var a []int = []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}
	var v int = 5
	fmt.Println(findValue(a, v, 0, len(a)-1))
}

// a is the slice of ints to search; v is the value to find; s is the start location to peek; e is the end location to peek, i is the index in a of value v (panics if can't find; extend to return an error if v might not be in a and caller can check for the error)
func findValue(a []int, v, s, e int) int {
	var m int = (e-s)/2 + s
	if a[m] == v {
		return m
	} else if v > a[m] && a[s] > v { // v is on the "right"
		return findValue(a, v, m, e)
	} else {
		return findValue(a, v, s, m)
	}
}

func four() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 9, -1, -1, -1}
	for i := 0; i < 100; i++ {
		a = append(a, -1)
	}
	v := 6
	fmt.Println(fourSearch(a, v), v)
}
func fourSearch(a []int, v int) int {
	i := 10
	for {
		e := a[i]
		if e == -1 || e > v {
			i = i/2 + 1
		} else if e < v {
			i = i * 2
		} else if e == v {
			return i
		}
	}
}

func five() {
	var str string = "cat"
	var strs []string = []string{"", "cat", "", "", "dog", ""}
	fmt.Println("find string, should be 1")
	fmt.Println(findString(str, strs, 0, len(strs)-1, false))
	fmt.Println("find string, should be 4")
	str = "dog"
	fmt.Println(findString(str, strs, 0, len(strs)-1, false))
}
func findString(str string, strs []string, s, e int, up bool) int {
	if str == "" {
		panic("cannot search for empty string as it may be duplicated")
	}
	m := (e-s)/2 + s
	if strs[m] == str {
		return m
	} else if strs[m] == "" {
		if up {
			return findString(str, strs, s+1, e, up)
		} else {
			return findString(str, strs, s, e-1, up)
		}
	} else if str > strs[m] {
		return findString(str, strs, m, e, true)
	} else if str < strs[m] {
		return findString(str, strs, s, m, false)
	} else {
		panic("unhandled")
	}
}

type bitVec struct {
	length uint
	bits   []uint64
}

// create a new bit vec of n "bytes" of bits (packed into the slice)
func bitVecNew(n uint) *bitVec {
	var b *bitVec = new(bitVec)
	b.bits = make([]uint64, n)
	b.length = n
	return b
}
func (b *bitVec) calculateBitLocation(n uint) (pos, slot uint) {
	pos = n / 64
	slot = n % 64
	return pos, slot
}

// returns true if already set, false if not (but now set)
func (b *bitVec) set(n uint) bool {
	pos, slot := b.calculateBitLocation(n)
	if b.get(n) { // bit was set
		return true
	} else { // set the bit
		b.bits[pos] = b.bits[pos] | 1<<slot
		return false
	}
}

// returns true if bit is set
func (b *bitVec) get(n uint) bool {
	pos, slot := b.calculateBitLocation(n)
	if b.bits[pos]>>slot&1 == 1 {
		return true
	}
	return false
}

// need 32k of bits!  that's a lot of bits.
func eight() {
	var b *bitVec = bitVecNew(512)
	var a []uint = make([]uint, 32000)
	for i := 0; i < len(a); i++ {
		a[i] = uint(i)
	}
	a[10] = 11
	for i := 0; i < len(a); i++ {
		if b.set(a[i]) {
			fmt.Println("found duplicate at ", i, "was", a[i])
		}
	}
}

func nine() {
	var matrix [][]int = createMatrix()
    var s []int = []int{35, 63, 83, 126}
    for i := range s {
	    r, c := searchMatrix(s[i], matrix)
        fmt.Printf("%v is located at %v, %v: %v\n", s[i], r, c, matrix[r][c])
    }
}
func searchMatrix(v int, matrix [][]int) (int, int) {
	m := len(matrix) / 2
	r := searchRows(v, matrix, 0, m, len(matrix)-1)
	m = len(matrix[r]) / 2
	c := searchColumns(v, matrix[r], 0, m, len(matrix[r])-1)
	return r, c
}
func searchRows(v int, matrix [][]int, s, m, e int) int {
	if v > matrix[m][0] && v < matrix[m][len(matrix[m])-1] {
		return m
	} else if v < matrix[m][0] {
		return searchRows(v, matrix, s, (m-s)/2+s, m)
	} else if v > matrix[m][len(matrix[m])-1] {
		return searchRows(v, matrix, m, (e-m)/2+m, e)
	} else {
		panic("foobar")
	}
}
func searchColumns(v int, smatrix []int, s, m, e int) int {
	if v == smatrix[m] {
		return m
	} else if v < smatrix[m] {
		return searchColumns(v, smatrix, s, (m-s)/2+s, m)
	} else if v > smatrix[m] {
		return searchColumns(v, smatrix, m, (e-m)/2+m, e)
	} else {
		panic("foobar")
	}
}

/*
rs: 5, re: 10, rm: 7, cs: 4, ce: 8, cm: 6, m: 118
rs: 5, re: 7, rm: 6, cs: 4, ce: 6, cm: 5, m: 108
rs: 5, re: 6, rm: 5, cs: 4, ce: 5, cm: 4, m: 81*/
// searching for v
func binarySearchMatrix(v int, matrix [][]int, rs, cs, re, ce int) (int, int) {
	var rm, cm int = (re-rs)/2 + rs, (ce-cs)/2 + cs
	fmt.Printf("rs: %v, re: %v, rm: %v, cs: %v, ce: %v, cm: %v, m: %v\n", rs, re, rm, cs, ce, cm, matrix[rm][cm])
	if matrix[rm][cm] == v {
		return rm, cm
	} else if matrix[rm][cm] < v {
		return binarySearchMatrix(v, matrix, rm, cm, re, ce)
	} else if matrix[rm][cm] > v {
		return binarySearchMatrix(v, matrix, rs, cs, rm, cm)
	} else {
		panic("huh")
	}
}
func createMatrix() [][]int {
	var matrix [][]int = make([][]int, 10)
	var n int
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, 8)
		for j := 0; j < len(matrix[i]); j++ {
			n = n + rand.Intn(5)
			matrix[i][j] = n
		}
	}
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	return matrix
}

func main() {
	fmt.Println("hello")

	one()
	two()
	three()
	four()
	five()
	eight()
	nine()
}
