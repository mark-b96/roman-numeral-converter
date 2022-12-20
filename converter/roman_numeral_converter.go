package converter

import (
	"math"
	"strconv"
)

var ROMANTOARABIC = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var ARABICTOROMAN = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func RomanToArabic(inputRomanSeq string) int {
	var total int = 0
	var currentBase float64 = 10.0
	var currentTotals []int
	var inputRomanSeqLength int = len(inputRomanSeq)

	for i := inputRomanSeqLength - 1; i > -1; i-- {

		var romanChar string = inputRomanSeq[i : i+1]
		var arabicInt int = ROMANTOARABIC[romanChar]
		var arabicFloat float64 = float64(arabicInt)

		if arabicFloat > currentBase {
			currentBase = math.Pow(currentBase, 2)
			for _, v := range currentTotals {
				total += v
			}
			currentTotals = currentTotals[:0]
		}

		if len(currentTotals) > 0 {
			if currentTotals[len(currentTotals)-1] > arabicInt {
				arabicInt = -arabicInt
			}
		}

		currentTotals = append(currentTotals, arabicInt)

	}

	for _, v := range currentTotals {
		total += v
	}
	return total
}

func ArabicToRoman(inputArabicSeq int) string {
	var inputArabicSeqStr string = strconv.Itoa(inputArabicSeq)
	var inputArabicSeqLength = len(inputArabicSeqStr)
	var currentBase int = 1
	var romanNumeral string = ""

	for i := inputArabicSeqLength - 1; i > -1; i-- {

		char := inputArabicSeqStr[i : i+1]
		targetNumber, _ := strconv.Atoi(char)
		targetNumber *= currentBase

		romanNumber, isInMap := ARABICTOROMAN[targetNumber]

		if isInMap {
			romanNumeral = romanNumber + romanNumeral
			currentBase *= 10
			continue
		}

		currentRoman := ""
		switch {
		case (5*currentBase < targetNumber) && (targetNumber < 9*currentBase):
			currentRoman += ARABICTOROMAN[5*currentBase]
			multiplierFactor := (targetNumber / currentBase) - 5
			for i := 0; i < multiplierFactor; i++ {
				currentRoman += ARABICTOROMAN[currentBase]
			}
		case targetNumber == 9*currentBase:
			currentRoman = ARABICTOROMAN[currentBase] + ARABICTOROMAN[10*currentBase]
		case (currentBase < targetNumber) && (targetNumber < 4*currentBase):
			multiplierFactor := targetNumber / currentBase
			for i := 0; i < multiplierFactor; i++ {
				currentRoman += ARABICTOROMAN[currentBase]
			}
		case targetNumber == 4*currentBase:
			currentRoman = ARABICTOROMAN[currentBase] + ARABICTOROMAN[5*currentBase]
		}

		romanNumeral = currentRoman + romanNumeral

		currentBase *= 10
	}

	return romanNumeral

}
