package main

import (
	"fmt"
	"strconv"
	"strings"
)

var morseCode = map[string]string{
	"1": ".----", "2": "..---", "3": "...--", "4": "....-", "5": ".....",
	"6": "-....", "7": "--...", "8": "---..", "9": "----.", "0": "-----",
}

var ninePalaceGrid = map[string]string{
	"A": "21", "B": "22", "C": "23",
	"D": "31", "E": "32", "F": "33",
	"G": "41", "H": "42", "I": "43",
	"J": "51", "K": "52", "L": "53",
	"M": "61", "N": "62", "O": "63",
	"P": "71", "Q": "72", "R": "73", "S": "74",
	"T": "81", "U": "82", "V": "83",
	"W": "91", "X": "92", "Y": "93", "Z": "94",
}

var keyboard = "QWERTYUIOPASDFGHJKLZXCVBNM"
var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func reverseMap(m map[string]string) map[string]string {
	result := make(map[string]string)
	for key, value := range m {
		result[value] = key
	}
	return result
}

func encryptMorse(ciphertext string) []int {
	list := strings.Split(ciphertext, "/")
	m := reverseMap(morseCode)
	var nums []int
	for _, ele := range list {
		if len(ele) != 0 {
			num, err := strconv.Atoi(m[ele])
			if err != nil {
				panic("ascii to int error")
			}
			nums = append(nums, num)
		}
	}
	return nums
}

func decryptMorse(nums []int) string {
	var result []string
	for _, list := range nums {
		value := fmt.Sprintf("%d", list)
		result = append(result, morseCode[string(value[0])], morseCode[string(value[1])])
	}
	return strings.Join(result, "/")
}

func encryptNinePalaceGrid(nums []int) []string {
	m := reverseMap(ninePalaceGrid)
	var letterList []string
	for i := 0; i < len(nums)-1; i += 2 {
		value := fmt.Sprintf("%d%d", nums[i], nums[i+1])
		letterList = append(letterList, m[value])
	}
	return letterList
}

func decryptNinePalaceGrid(letters []string) []int {
	var result []int
	for _, letter := range letters {
		num, err := strconv.Atoi(ninePalaceGrid[letter])
		if err != nil {
			panic("ascii to int error")
		}
		result = append(result, num)
	}
	return result
}

func mapKeyboardLetter(letterList []string, from, to string) []string {
	var newLetterList []string
	for _, letter := range letterList {
		newLetterList = append(newLetterList, string(to[strings.Index(from, letter)]))
	}
	return newLetterList
}

func decryptFence(letterList []string) []string {
	var length int
	if len(letterList)%2 == 0 {
		length = len(letterList) / 2
	} else {
		length = len(letterList)/2 + 1
		letterList = append(letterList, "")
	}

	fence1 := letterList[:length]
	fence2 := letterList[length:]

	var result []string
	for idx := 0; idx < length; idx++ {
		result = append(result, fence1[idx], fence2[idx])
	}
	return result
}

func encryptFence(letterList []string) []string {
	var fence1 []string
	var fence2 []string
	for idx, letter := range letterList {
		if idx%2 == 0 {
			fence1 = append(fence1, letter)
		} else {
			fence2 = append(fence2, letter)
		}
	}
	return append(fence1, fence2...)
}

func reverse(input []string) []string {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

func decrypt(ciphertext string) string {
	ciphertext = strings.TrimSpace(ciphertext)
	morseNumList := encryptMorse(ciphertext)
	letterList := encryptNinePalaceGrid(morseNumList)
	newLetterList := mapKeyboardLetter(letterList, keyboard, alphabet)
	fenceResult := decryptFence(newLetterList)
	result := reverse(fenceResult)
	return strings.Join(result, "")
}

func encrypt(rawtext string) string {
	rawtext = strings.ToUpper(strings.TrimSpace(rawtext))
	ciphertext := reverse(strings.Split(rawtext, ""))
	ciphertext = encryptFence(ciphertext)
	letterList := mapKeyboardLetter(ciphertext, alphabet, keyboard)
	nums := decryptNinePalaceGrid(letterList)
	result := decryptMorse(nums)
	return result
}

func main() {
	rawtext := "ILOVEYOUTOO"
	result := encrypt(rawtext)
	fmt.Println(result) // ....-/.----/----./....-/....-/.----/---../.----/....-/.----/-..../...--/....-/.----/----./..---/-..../..---/..---/...--/--.../....-

	ciphertext := "....-/.----/----./....-/....-/.----/---../.----/....-/.----/-..../...--/....-/.----/----./..---/-..../..---/..---/...--/--.../....-/"
	result_ := decrypt(ciphertext)
	fmt.Println(result_) // ILOVEYOUTOO
}
