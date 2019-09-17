package solution

import (
	"strconv"
)

/*
Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

Example:

n = 15,

Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
*/

func fizzBuzz(n int) []string {
	res := make([]string, n)

	for i := 0; i < n; i++ {
		res[i] = strconv.Itoa(i + 1)
	}

	for i := 1; 3*i <= n; i++ {
		res[3*i-1] = "Fizz"
	}

	for i := 1; 5*i <= n; i++ {
		res[5*i-1] = "Buzz"
	}

	for i := 1; 15*i <= n; i++ {
		res[15*i-1] = "FizzBuzz"
	}
	return res
}
