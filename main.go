package main

import (
	"fmt"
	"roman-numeral-converter/converter"
	"strconv"
)



func main() {
	var arabicInputStr string
	fmt.Scan(&arabicInputStr)
	var arabicInputInt, _ = strconv.Atoi(arabicInputStr)


	romanNumeral := converter.ArabicToRoman(arabicInputInt)
	fmt.Println(converter.RomanToArabic(romanNumeral))

}