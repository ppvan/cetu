package main

var BASE_CHARS = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var BASE int64 = int64(len(BASE_CHARS))
var BASE_COUNTER int64 = 100_000_000_000
var MAX_LENGTH = 7

// base62Encode encodes a number to a base62 string.
// Param number: the number to encode.
// Return: the base62 string.
func base62Encode(number int64) (str string) {
	result := make([]byte, MAX_LENGTH)
	i := 0
	number += BASE_COUNTER
	for number > 0 {
		result[i] = BASE_CHARS[number%BASE]
		number = number / BASE
		i++
	}
	return string(reverseBytes(result[:i]))
}

// ReverseBytes reverses a byte array.
// It also change the original array.
// Examples: []byte("Hello, World!") -> []byte("!dlroW ,olleH")
func reverseBytes(bytes []byte) []byte {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	return bytes
}
