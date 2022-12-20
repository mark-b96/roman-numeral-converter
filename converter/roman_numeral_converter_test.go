package converter_test

import (
	"roman-numeral-converter/converter"
	"testing"
)


func TestRomanToArabic(t *testing.T){
	if converter.RomanToArabic("V") != 5 {
		t.Fatal("Conversion failure") 
	}
}

func TestArabicToRoman(t *testing.T){
	if converter.ArabicToRoman(5) != "V" {
		t.Fatal("Conversion failure") 
	}
}

func TestConverter(t *testing.T){
	for i := 0; i < 3999; i++{
		if i != converter.RomanToArabic(converter.ArabicToRoman(i)) {
			t.Fatal("Conversion failure") 
		}
	}
}
