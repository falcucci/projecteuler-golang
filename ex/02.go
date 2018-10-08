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
//  exercise 02  //
///////////////////

Each new term in the Fibonacci sequence is generated by adding
the previous two terms. By starting with 1 and 2,
the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values
do not exceed four million, find the sum of the even-valued
terms.
*/

package main

import (
	"fmt"
)

/* generate a fibonacci sequence based in the max of terms given */
func Fibonacci(limit int) int {
	i := 0
	r := 1
	sum := 0
	for i < limit {
		i, r = r, i+r
		if r%2 == 0 {
			sum += r
		}
	}
	return sum
}

func main() {
	sum := Fibonacci(4000000)
	fmt.Println(sum)
}
