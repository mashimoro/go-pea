package fizzbuzz

import "strconv"

func Fizzbuzz(number int) string {

	if number%3 == 0 && number%5 == 0 {
		return "Fizzbuzz"
	}
	if number%3 == 0 {
		return "fizz"
	}
	if number%5 == 0 {
		return "buzz"
	}
	return strconv.Itoa(number)
}
