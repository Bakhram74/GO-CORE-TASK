package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
)

var numDecimal int = 42
var numOctal int = 052
var numHexadecimal int = 0x2A
var pi float64 = 3.14
var name string = "Golang"
var isActive bool = true
var complexNum complex64 = 1 + 2i

func main() {
	fmt.Printf("type %T\n", numDecimal)
	fmt.Println(reflect.TypeOf(numOctal))
	fmt.Println(reflect.TypeOf(numHexadecimal))
	fmt.Println(reflect.TypeOf(pi))
	fmt.Printf("type %T\n", name)
	fmt.Println(reflect.TypeOf(isActive))
	fmt.Println(reflect.TypeOf(complexNum))
	result := combineValuesToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)

	fmt.Println(hashWithSalt([]rune(result), "go-2024"))
}

func combineValuesToString(n1, n2, n3 int, f float64, str string, b bool, compl complex64) string {
	return fmt.Sprintf("%d%d%d%1.2f%s%t%v", n1, n2, n3, f, str, b, compl)
}

func hashWithSalt(password []rune, salt string) string {

	mid := len(password) / 2
	start := password[:mid]
	end := password[mid:]

	passwordWithSalt := string(start) + salt + string(end)

	hasher := sha256.New()
	hasher.Write([]byte(passwordWithSalt))
	return hex.EncodeToString(hasher.Sum(nil))
}
