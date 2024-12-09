package main



import "testing"

func TestCombineValuesToString(t *testing.T) {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	result := combineValuesToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	expected := "4242423.14Golangtrue(1+2i)"
	if result != expected {
		t.Errorf("not equal expected:%s to result:%s", expected, result)
	}

	notExpected := "4242423.14Golangtrue(1+2)"

	if result == notExpected {
		t.Errorf("result:%s must not be equal to notExpected:%s", result, notExpected)
	}
}

func TestHashWithSalt(t *testing.T) {
	runes := []rune("4242423.14Golangtrue(1+2i)")
	salt := "go-2024"
	length := 64
	result := hashWithSalt(runes, salt)

	if len(result) != length {
		t.Errorf("Expected hash length %d, but got %d", length, len(result))
	}
}
