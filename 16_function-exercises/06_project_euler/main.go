package main

import "fmt"

/*
Working from left-to-right if no digit is exceeded by the digit to its left it is called an increasing number; for example, 134468.

Similarly if no digit is exceeded by the digit to its right it is called a decreasing number; for example, 66420.

We shall call a positive integer that is neither increasing nor decreasing a "bouncy" number; for example, 155349.

Clearly there cannot be any bouncy numbers below one-hundred, but just over half of the numbers below one-thousand (525) are bouncy. In fact, the least number for which the proportion of bouncy numbers first reaches 50% is 538.

Surprisingly, bouncy numbers become more and more common and by the time we reach 21780 the proportion of bouncy numbers is equal to 90%.

Find the least number for which the proportion of bouncy numbers is exactly 99%.
 */

 var acc float64 = 0.99

func main() {
	var inc int
	var dec int
	var bouncy int
	//No bouncy less than 100 starting from 100
	for i := 100 ;; i++{
		//Check for increasing sequence
 		for j := i ; j >= 10; {
 			 prev := j % 10
 			 j = j / 10
 			 current := j % 10
			if prev >= current {
				if j < 10 {
					inc++
				}
			} else {
				break
			}
		}
		//Check for decreasing sequence
		for j := i ; j >= 10; {
			prev := j % 10
			j = j / 10
			current := j % 10
			if current >= prev {
				if j < 10 {
					dec++
				}
			} else {
				break
			}
		}
		//Check for bouncy number
		if !(inc == 1 || dec == 1) {
			bouncy++
		}
		inc = 0
		dec = 0
		if bouncy != 0 {
			if  float64(bouncy)/ float64(i)  >= acc {
				//Maximum bouncy for which the ratio is 99 percent
				fmt.Printf("%d is the least number for which the proportion is %f percent.",i,acc*100)
				break
			}
		}
	}
}
