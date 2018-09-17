package main

import "fmt"

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
	var carry uint8
	var i uint8
	for i = 0; i < oneBits; i++ {
		var mask uint8 = ^uint8(0) >> 7 << i
		at, bt := a&mask, b&mask
		if at^bt == mask {
			r = r | at | bt
		} else if at&mask == mask && bt&mask == mask {
			// both one
			carry++
		} else if carry > 0 {
			// when carried, how does that manifest?
			// what happens to carried bits when you have an xor == 1 or both are 1?
			if at&bt == 0 {
				// set the bits
				r = r | mask
				for j := carry-1; j > 0; j-- {
					r = r | (mask >> j)
				}
			    carry = 0
			}
		}
	}
	return
}

func main() {
	fmt.Println("hello")
    one()
}
