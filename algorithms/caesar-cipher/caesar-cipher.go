package caesar_cipher

import "fmt"

const NUMBER_OF_LETTERS = 26

func main() {
	var l int
	var k uint8
	var s string
	fmt.Scan(&l)
	fmt.Scan(&s)
	fmt.Scan(&k)

	fmt.Println(cipher(s, k))
}

func cipher(s string, k uint8) (r string) {
	r = ""
	for _, v := range s {
		switch {
		default:
			r += string(v)
		case isUppercase(v):
			a := uint8(v) + k%NUMBER_OF_LETTERS
			if a > uint8('Z') {
				b := a - uint8('Z')
				r += string(b + uint8('A') - 1)
			} else {
				r += string(a)
			}
		case isLowercase(v):
			a := uint8(v) + k%NUMBER_OF_LETTERS
			if a > uint8('z') {
				b := a - uint8('z')
				r += string(b + uint8('a') - 1)
			} else {
				r += string(a)
			}
		}
	}

	return r
}

func isLowercase(char rune) bool {
	return char >= 'a' && char <= 'z'
}

func isUppercase(char rune) bool {
	return char >= 'A' && char <= 'Z'
}
