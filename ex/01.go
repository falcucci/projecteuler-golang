/*
Copyright © 2018 falcucci

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the "Software"),
to deal in the Software without restriction, including without limitation
the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the
Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

///////////////////
//  exercide 01  //
///////////////////

If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/
package main

import (
	"fmt"
)

/* method to verify if an item exists in a Slice */
func Has(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

/* sum of the multiple based on any numbers */
func SumMultiples(nums []int) int {
	var multiples []int
	var sum int
	for _, v := range nums {
		for _, n := range nums {
			if (3*n == v || 5*n == v) && !Has(multiples, v) {
				sum += v
				multiples = append(multiples, v)
			}
		}
	}
	return sum
}

func BuildSlice(max int) []int {
	slice := make([]int, max-1)
	for k := range slice {
		slice[k] = k + 1
	}
	return slice
}

func main() {
	numbers := BuildSlice(1000)
	fmt.Println(numbers)
	SumMultiples(numbers)
}
